package sShippingZone

import (
	"shopkone-service/internal/module/delivery/shipping/mShipping"
)

type CodeIsChangeOut struct {
	Remove []mShipping.ShippingZoneCode
	Insert []mShipping.ShippingZoneCode
	Update []mShipping.ShippingZoneCode
}

func (s *sShippingZone) CodesIsChange(newCodes, oldCodes []mShipping.ShippingZoneCode) (out CodeIsChangeOut) {
	// 标记在 newCodes 中出现但不在 oldCodes 中的项为 Insert
	for _, newCode := range newCodes {
		found := false
		for _, oldCode := range oldCodes {
			if newCode.CountryCode == oldCode.CountryCode {
				found = true
				// 比较 ZoneCodes 的内容是否相同，判断是否需要 Update
				if !equalStringSlices(newCode.ZoneCodes, oldCode.ZoneCodes) {
					out.Update = append(out.Update, newCode)
				}
				break
			}
		}
		if !found {
			out.Insert = append(out.Insert, newCode)
		}
	}

	// 标记在 oldCodes 中存在但不在 newCodes 中的项为 Remove
	for _, oldCode := range oldCodes {
		found := false
		for _, newCode := range newCodes {
			if oldCode.CountryCode == newCode.CountryCode {
				found = true
				break
			}
		}
		if !found {
			out.Remove = append(out.Remove, oldCode)
		}
	}

	return out
}

// 比较两个字符串切片是否相等的辅助函数
func equalStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
