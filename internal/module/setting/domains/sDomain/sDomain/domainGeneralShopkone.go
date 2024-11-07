package sDomain

import (
	"math/rand"
	"shopkone-service/internal/module/setting/domains/mDomains"
	"shopkone-service/utility/code"
	"strings"
	"time"
)

func (s *sDomain) GeneralShopkoneDomain() (domain string, err error) {
	//生成子域名前缀，预计12位
	prefix := generateSubdomain(12)
	domain = prefix + "." + "shopkone.com"
	for i := 0; i < 1000; i++ {
		// 判断是否重复了，重复则重新生成
		var count int64
		if err = s.orm.Model(mDomains.Domain{}).Where("domain = ?", domain).
			Count(&count).Error; err != nil {
			return domain, err
		}
		if count == 0 {
			return domain, nil
		}
	}
	return domain, code.SystemError
}

func generateSubdomain(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	const hyphen = '-'
	if length < 1 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder

	for i := 0; i < length; i++ {
		if i == 0 || i == length-1 {
			// First and last character should not be a hyphen
			sb.WriteByte(charset[rand.Intn(len(charset))])
		} else {
			// Allow both charset and hyphen for middle characters
			if rand.Float32() < 0.1 { // 10% chance of adding a hyphen
				sb.WriteByte(hyphen)
			} else {
				sb.WriteByte(charset[rand.Intn(len(charset))])
			}
		}
	}

	return sb.String()
}
