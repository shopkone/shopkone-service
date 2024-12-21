package sPolicy

import "shopkone-service/internal/module/shop/policy/mPolicy"

func (s *sPolicy) PolicyInit() (err error) {
	privacyData := mPolicy.Policy{
		Body:  "",
		Title: "Privacy Policy",
		Type:  mPolicy.PolicyTypePrivacy,
		Url:   "/privacy-policy",
	}
	if err = s.orm.Create(&privacyData).Error; err != nil {
		return err
	}

	refundData := mPolicy.Policy{
		Body:  "",
		Title: "Refund Policy",
		Type:  mPolicy.PolicyTypeRefund,
		Url:   "/refund-policy",
	}
	if err = s.orm.Create(&refundData).Error; err != nil {
		return err
	}

	serviceData := mPolicy.Policy{
		Body:  "",
		Title: "Service Policy",
		Type:  mPolicy.PolicyTypeService,
		Url:   "/service-policy",
	}
	if err = s.orm.Create(&serviceData).Error; err != nil {
		return err
	}

	shippingData := mPolicy.Policy{
		Body:  "",
		Title: "Shipping Policy",
		Type:  mPolicy.PolicyTypeShipping,
		Url:   "/shipping-policy",
	}
	return s.orm.Create(&shippingData).Error
}
