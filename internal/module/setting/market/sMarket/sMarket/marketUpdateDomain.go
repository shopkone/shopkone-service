package sMarket

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/setting/market/mMarket"
	"shopkone-service/utility/code"
)

func (s *sMarket) MarketUpdateDomain(in vo.MarketUpDomainReq) (err error) {
	data := mMarket.Market{
		DomainSuffix: in.DomainSuffix,
		DomainType:   in.DomainType,
		SubDomainID:  in.SubDomainID,
	}
	if in.DomainType == mMarket.DomainTypeSuffix && in.DomainSuffix == "" {
		return code.MarketMustPrefixDomain
	}
	if in.DomainType == mMarket.DomainTypeSub && in.SubDomainID == 0 {
		return code.MarketMustSubDomain
	}
	if err = s.orm.Model(&mMarket.Market{}).
		Where("id = ?", in.ID).
		Select("domain_suffix", "domain_type", "sub_domain_id").
		Updates(&data).Error; err != nil {
		return err
	}

	_, err = s.MarketCheck(false, in.ID)
	return err
}
