package mProduct

type Vendor struct {
	Name string `gorm:"size:200"`
	Link string `gorm:"size:500"`
}
