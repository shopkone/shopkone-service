package sCollection

import (
	"github.com/duke-git/lancet/v2/slice"
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/product/collection/iCollection"
	"shopkone-service/internal/module/product/collection/mCollection"
	"shopkone-service/internal/module/product/inventory/mInventory"
	"shopkone-service/internal/module/product/product/mProduct"
)

func MatchOne(action string, value string, key string) (out iCollection.MatchOneOut) {
	a := MatchAction(action, value)
	if a.Valid == false {
		out.Valid = false
		return out
	}
	out.Type = iCollection.MatchOneOutTypeProduct
	out.Valid = true
	out.Value = a.Value
	out.NeedValue = a.NeedValue
	switch key {
	case "title":
		out.Query = "title" + a.Query
		out.Type = iCollection.MatchOneOutTypeProduct
	case "type":
		out.Query = "type" + a.Query
		out.Type = iCollection.MatchOneOutTypeProduct
	case "category":
		out.Query = "category" + a.Query
		out.Type = iCollection.MatchOneOutTypeProduct
	case "vendor":
		out.Query = "vendor" + a.Query
		out.Type = iCollection.MatchOneOutTypeProduct
	case "tag":
		out.Query = "tags Like ?"
		out.Type = iCollection.MatchOneOutTypeProduct
	case "price":
		out.Query = "price" + a.Query
		out.Type = iCollection.MatchOneOutTypeVariant
	case "compare_at_price":
		out.Query = "compare_at_price" + a.Query
		out.Type = iCollection.MatchOneOutTypeVariant
	case "weight":
		out.Query = "weight" + a.Query
		out.Type = iCollection.MatchOneOutTypeVariant
	case "inventory_stock":
		out.Query = "quantity" + a.Query
		out.Type = iCollection.MatchOneOutTypeInventory
	case "variant_title":
		out.Query = "value" + a.Query
		out.Type = iCollection.MatchOneOutTypeVariantName
	}
	return out
}

func MatchAction(action string, value string) (out iCollection.MatchActionOut) {
	out.Valid = true
	out.Value = value
	out.NeedValue = true
	switch action {
	case "eq":
		out.Query = " = ?"
		return out
	case "neq":
		out.Query = " != ?"
		return out
	case "sw":
		out.Value = value + "%"
		out.Query = " LIKE ?"
		return out
	case "ew":
		out.Value = "%" + value
		out.Query = " LIKE ?"
		return out
	case "ct":
		out.Value = "%" + value + "%"
		out.Query = " LIKE ?"
		return out
	case "nct":
		out.Query = " NOT LIKE ?"
		out.Value = "%" + value + "%"
		return out
	case "gt":
		out.Query = " > ?"
		return out
	case "lt":
		out.Query = " < ?"
		return out
	case "empty":
		out.Query = " IS NULL"
		out.NeedValue = false
		return out
	case "nempty":
		out.Query = " IS NOT NULL"
		out.NeedValue = false
		return out
	}
	out.Valid = false
	return out
}

func FilterKey(key string) bool {
	switch key {
	case "title", "type", "category", "vendor", "tag", "price", "compare_at_price", "weight", "inventory_stock", "variant_title":
		return true
	}
	return false
}

func FilterAction(action string) bool {
	switch action {
	case "eq", "neq", "sw", "ew", "ct", "nct", "gt", "lt", "empty", "nempty":
		return true
	}
	return false
}

func GenSql(conditions []vo.BaseCondition, mode mCollection.CollectionMatchMode, Type iCollection.MatchOneOutType) (query string, values []interface{}) {
	slice.ForEach(conditions, func(index int, item vo.BaseCondition) {
		ret := MatchOne(item.Action, item.Value, item.Key)
		if ret.Type == Type {
			if query == "" {
				query = query + ret.Query
			} else if mode == mCollection.CollectionMatchModeAny {
				query = query + " OR " + ret.Query
			} else if mode == mCollection.CollectionMatchModeAll {
				query = query + " AND " + ret.Query
			}
			if ret.NeedValue {
				values = append(values, ret.Value)
			}
		}
	})
	return query, values
}

func (s *sCondition) MatchProducts(in iCollection.MatchProductsIn) (err error) {

	productQuery, productValues := GenSql(in.Conditions, in.MatchMode, iCollection.MatchOneOutTypeProduct)
	variantQuery, variantValues := GenSql(in.Conditions, in.MatchMode, iCollection.MatchOneOutTypeVariant)
	inventoryQuery, inventoryValues := GenSql(in.Conditions, in.MatchMode, iCollection.MatchOneOutTypeInventory)
	variantNameQuery, variantNameValues := GenSql(in.Conditions, in.MatchMode, iCollection.MatchOneOutTypeVariantName)

	var ids []uint

	if variantQuery != "" {
		var variants []mProduct.Variant
		vq := s.orm.Model(&variants).Where("shop_id = ?", s.shopId)
		vq = vq.Where(variantQuery, variantValues...).Select("product_id")
		if err = vq.Find(&variants).Error; err != nil {
			return err
		}
		ids = slice.Map(variants, func(_ int, item mProduct.Variant) uint {
			return item.ProductId
		})
	}

	if productQuery != "" {
		var products []mProduct.Product
		pq := s.orm.Model(&products).Where("shop_id = ?", s.shopId)
		pq = pq.Where(productQuery, productValues...).Select("id")
		if err = pq.Find(&products).Error; err != nil {
			return err
		}
		ids = append(ids, slice.Map(products, func(_ int, item mProduct.Product) uint {
			return item.ID
		})...)
	}

	if inventoryQuery != "" {
		var inventories []mInventory.Inventory
		iq := s.orm.Model(&inventories).Where("shop_id = ?", s.shopId)
		iq = iq.Where(inventoryQuery, inventoryValues...).Select("variant_id")
		if err = iq.Find(&inventories).Error; err != nil {
			return err
		}
		variantIds := slice.Map(inventories, func(_ int, item mInventory.Inventory) uint {
			return item.VariantId
		})
		variantIds = slice.Unique(variantIds)
		if len(variantIds) != 0 {
			var variants []mProduct.Variant
			err = s.orm.Model(&variants).Select("product_id").Where("shop_id = ? AND id IN ?", s.shopId, variantIds).Find(&variants).Error
			if err != nil {
				return err
			}
			pids := slice.Map(variants, func(_ int, item mProduct.Variant) uint {
				return item.ProductId
			})
			ids = append(ids, slice.Unique(pids)...)
		}
	}

	if variantNameQuery != "" {
		var variantNames []mProduct.VariantNameHandler
		vq := s.orm.Model(&variantNames).Where("shop_id = ?", s.shopId)
		vq = vq.Where(variantNameQuery, variantNameValues...).Select("product_id")
		if err = vq.Find(&variantNames).Error; err != nil {
			return err
		}
		ids = append(ids, slice.Map(variantNames, func(_ int, item mProduct.VariantNameHandler) uint {
			return item.ProductId
		})...)
	}

	// 关联商品
	if len(ids) != 0 {
		//var cps []mCollection.CollectionProduct
		ids = slice.Unique(ids)
		cps := slice.Map(ids, func(_ int, item uint) mCollection.CollectionProduct {
			i := mCollection.CollectionProduct{}
			i.ShopId = s.shopId
			i.ProductId = item
			i.CollectionId = in.CollectionId
			return i
		})
		return s.orm.Create(&cps).Error
	}

	return err
}
