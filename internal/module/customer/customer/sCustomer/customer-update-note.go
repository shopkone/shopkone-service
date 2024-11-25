package sCustomer

import (
	"shopkone-service/internal/api/vo"
	"shopkone-service/internal/module/customer/customer/mCustomer"
)

func (s *sCustomer) CustomerUpdateNote(in *vo.CustomerUpdateNoteReq) (err error) {
	return s.orm.Model(&mCustomer.Customer{}).Where("id = ?", in.ID).Update("note", in.Note).Error
}
