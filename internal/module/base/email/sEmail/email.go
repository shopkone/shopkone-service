package sEmail

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/module/base/email/iEmail"
)

type IEmail interface {
	Send() (err error) // 发送邮件
}

type sEmail struct {
}

func (s *sEmail) Send(in iEmail.EmailSendIn) (err error) {
	g.Dump(in)
	return err
}

func NewSystemEmail() *sEmail {
	return &sEmail{}
}

func NewUserEmail() *sEmail {
	return &sEmail{}
}
