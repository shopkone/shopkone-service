package sLocation

import (
	"github.com/duke-git/lancet/v2/slice"
	"gorm.io/gorm"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/base/address/mAddress"
	"shopkone-service/internal/module/base/address/sAddress"
	"shopkone-service/internal/module/delivery/in-store-pick-up/sInStorePickup"
	"shopkone-service/internal/module/delivery/local-delivery/sLocalDelivery/sLocalDelivery"
	"shopkone-service/internal/module/product/inventory/sInventory/sInventory"
	"shopkone-service/internal/module/setting/location/mLocation"
	"shopkone-service/utility/code"
)

type sLocation struct {
	orm    *gorm.DB
	shopId uint
}

func NewLocation(orm *gorm.DB, shopId uint) *sLocation {
	return &sLocation{orm: orm, shopId: shopId}
}

type ILocation interface {
	List(active *bool) (list []vo.LocationListRes, err error)
	Create(in vo.LocationAddReq) (id uint, err error)
	Info(id uint) (info vo.LocationInfoRes, err error)
	Update(in vo.LocationUpdateReq) (err error)
	Delete(id uint) (err error)
	IsAllActive(ids []uint) (err error)
}

func (s *sLocation) List(active *bool) (list []vo.LocationListRes, err error) {
	// 查询列表
	var locationList []mLocation.Location
	query := s.orm.Model(&mLocation.Location{}).Order("order_num ASC")
	query = query.Where("shop_id = ?", s.shopId)
	query = query.Omit("deleted_at", "created_at", "updated_at")
	if active != nil {
		query = query.Where("active = ?", active)
	}
	if err = query.Find(&locationList).Error; err != nil {
		return
	}
	// 查询地址
	addressIds := slice.Map(locationList, func(index int, item mLocation.Location) uint {
		return item.AddressId
	})
	addressList, err := sAddress.NewAddress(s.orm, s.shopId).ListByIds(addressIds)
	if err != nil {
		return
	}
	// 组装数据
	list = slice.Map(locationList, func(index int, item mLocation.Location) vo.LocationListRes {
		address, ok := slice.FindBy(addressList, func(index int, a mAddress.Address) bool {
			return a.ID == item.AddressId
		})
		if !ok {
			return vo.LocationListRes{}
		}
		return vo.LocationListRes{
			Id:      item.ID,
			Active:  item.Active,
			Address: address,
			Name:    item.Name,
			Default: item.IsDefault,
			Order:   item.OrderNum,
		}
	})
	return list, err
}

func (s *sLocation) Create(in vo.LocationAddReq, timezone string) (id uint, err error) {
	// 查询名称是否已存在
	var count int64
	if err = s.orm.Model(&mLocation.Location{}).
		Where("shop_id = ? AND name = ?", s.shopId, in.Name).
		Count(&count).Error; err != nil {
		return id, err
	}
	if count > 0 {
		return 0, code.ErrLocationNameExist
	}
	// 创建地址
	addressId, err := sAddress.NewAddress(s.orm, s.shopId).CreateAddress(in.Address)
	if err != nil {
		return
	}
	// 查询是否有默认地址
	var defaultLocation mLocation.Location
	err = s.orm.Model(&defaultLocation).Where("shop_id = ? AND is_default = ?", s.shopId, true).
		Select("id").Find(&defaultLocation).Error
	if err != nil {
		return id, err
	}
	order, err := s.LocationGetNextOrder()
	if err != nil {
		return id, err
	}
	// 创建位置
	location := mLocation.Location{}
	location.OrderNum = uint(order)
	location.Name = in.Name
	location.AddressId = addressId
	location.Active = true
	location.ShopId = s.shopId
	location.FulfillmentDetails = true
	// 如果没有默认地址，则设置该位置为默认
	if defaultLocation.ID == 0 {
		location.IsDefault = true
	}
	if err = s.orm.Create(&location).Error; err != nil {
		return id, err
	}
	// 创建本地配送
	if err = sLocalDelivery.NewLocalDelivery(s.orm, s.shopId).LocalDeliveryCreate(location.ID); err != nil {
		return 0, err
	}
	// 创建到店自提
	if err = sInStorePickup.NewInStorePickup(s.orm, s.shopId).Create(location.ID, timezone); err != nil {
		return 0, err
	}
	return location.ID, nil
}

func (s *sLocation) Info(id uint) (info vo.LocationInfoRes, err error) {
	var location mLocation.Location
	err = s.orm.Model(&mLocation.Location{}).Where("id = ?", id).First(&location).Error
	if err != nil {
		return
	}
	addressList, err := sAddress.NewAddress(s.orm, s.shopId).ListByIds([]uint{location.AddressId})
	if addressList == nil || len(addressList) == 0 {
		return info, err
	}
	return vo.LocationInfoRes{
		Id:                 location.ID,
		Active:             location.Active,
		Address:            addressList[0],
		Name:               location.Name,
		Default:            location.IsDefault,
		FulfillmentDetails: location.FulfillmentDetails,
	}, err
}

func (s *sLocation) Update(in vo.LocationUpdateReq) (err error) {
	// 获取详情
	info := mLocation.Location{}
	err = s.orm.Select("is_default", "active", "address_id").Model(&info).
		Where("shop_id = ? AND id = ?", s.shopId, in.Id).First(&info).Error
	if err != nil {
		return err
	}
	// 判断是否是默认且为非启用
	if info.IsDefault && !in.Active {
		return code.ErrLocationDefaultDisable
	}
	// 如果是默认且设置为不允许作为发货地址，则报错
	if info.IsDefault && !in.FulfillmentDetails {
		return code.ErrLocationDefaultUnFulfillmentDetails
	}
	// 如果是禁用的，则判断是否有库存
	if in.Active != info.Active && !in.Active {
		has, err := sInventory.NewInventory(s.orm, s.shopId).ExistQuantityByLocationId(in.Id)
		if err != nil {
			return err
		}
		if has {
			return code.NoDeActiveByHasInventory
		}
	}
	// 更新地址
	in.Address.ID = info.AddressId
	if err = sAddress.NewAddress(s.orm, s.shopId).UpdateById(in.Address); err != nil {
		return err
	}
	// 更新位置
	data := mLocation.Location{
		Active:             in.Active,
		Name:               in.Name,
		FulfillmentDetails: in.FulfillmentDetails,
	}
	query := s.orm.Model(&data).Where("shop_id = ? AND id = ?", s.shopId, in.Id)
	if err = query.Select("active", "name", "fulfillment_details").Updates(data).Error; err != nil {
		return err
	}
	return err
}

func (s *sLocation) Delete(id uint) (err error) {
	// 获取详情
	info := mLocation.Location{}
	query := s.orm.Select("is_default").Model(&info).Where("shop_id = ? AND id = ?", s.shopId, id)
	if err = query.First(&info).Error; err != nil {
		return err
	}
	// 如果是默认的，则不允许删除
	if info.IsDefault {
		return code.ErrLocationDefaultDelete
	}
	// 如果状态是非停用，则不允许删除
	if info.Active {
		return code.ErrLocationDelete
	}
	// 删除本地配送
	if err = sLocalDelivery.NewLocalDelivery(s.orm, s.shopId).LocalDeliveryRemove(id); err != nil {
		return err
	}
	// 删除稻田自提
	if err = sInStorePickup.NewInStorePickup(s.orm, s.shopId).RemoveByLocationId(id); err != nil {
		return err
	}
	// 删除库存
	return s.orm.Model(&mLocation.Location{}).Where("shop_id = ? AND id = ?", s.shopId, id).Delete(&mLocation.Location{}).Error
}

func (s *sLocation) IsAllActive(ids []uint) (err error) {
	// 查询列表
	ids = slice.Unique(ids)
	var count int64
	query := s.orm.Model(&mLocation.Location{}).Where("shop_id = ? AND id IN ?", s.shopId, ids)
	query = query.Where("active = ?", true)
	if err = query.Count(&count).Error; err != nil {
		return err
	}
	if count != int64(len(ids)) {
		return code.ErrLocationNotAllActive
	}
	return err
}
