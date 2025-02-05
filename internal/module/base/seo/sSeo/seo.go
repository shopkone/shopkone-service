package sSeo

import (
	"gorm.io/gorm"
	"shopkone-service/internal/module/base/seo/mSeo"
	"shopkone-service/utility/code"
)

type ISeo interface {
	Create(seo mSeo.Seo) (seoId uint, err error)
	Info(id uint) (seo mSeo.Seo, err error)
	Update(seo mSeo.Seo) (err error)
}

type sSeo struct {
	orm    *gorm.DB
	shopId uint
}

func NewSeo(orm *gorm.DB, shopId uint) *sSeo {
	return &sSeo{orm: orm, shopId: shopId}
}

func (s *sSeo) Create(seo mSeo.Seo) (seoId uint, err error) {
	seo.ShopId = s.shopId
	seo.ID = 0
	err = s.orm.Model(&seo).Create(&seo).Error
	if err != nil {
		return seoId, err
	}
	return seo.ID, nil
}

func (s *sSeo) Info(id uint) (seo mSeo.Seo, err error) {
	err = s.orm.Model(&seo).
		Where("id = ? and shop_id = ?", id, s.shopId).
		Select("id", "meta_description", "page_title", "url_handle", "follow").First(&seo).Error
	return seo, err
}

func (s *sSeo) Update(seo mSeo.Seo) (err error) {
	if seo.ID == 0 {
		return code.SeoIdMissing
	}
	err = s.orm.Model(&seo).
		Where("id = ? and shop_id = ?", seo.ID, s.shopId).
		Select("meta_description", "page_title", "follow", "url_handle").
		Updates(seo).Error
	return err
}
