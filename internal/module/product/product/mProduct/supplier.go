package mProduct

import "shopkone-service/internal/module/base/orm/mOrm"

type Supplier struct {
	mOrm.Model
	AddressId uint
}
