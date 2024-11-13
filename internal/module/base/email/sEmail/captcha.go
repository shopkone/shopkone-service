package sEmail

import (
	"bytes"
	"github.com/gogf/gf/v2/util/grand"
	"html/template"
	"shopkone-service/internal/module/base/cache/sCache"
	"shopkone-service/internal/module/base/email/iEmail"
	"shopkone-service/utility/code"
)

type ICaptcha interface {
	Send(in iEmail.CaptchaSendIn) (err error)                         // 发送验证码
	Verify(email string, code string, sense iEmail.Sense) (err error) // 验证验证码
	Remove(email string, sense iEmail.Sense) (err error)              // 移除验证码
	getKey(email string, sense iEmail.Sense) string                   // 获取验证码的key
}

type sCaptcha struct{}

func NewCaptcha() *sCaptcha {
	return &sCaptcha{}
}

func (s *sCaptcha) Send(in iEmail.CaptchaSendIn) (err error) {
	key := s.getKey(in.Email, in.Sense)
	// 获取缓存信息
	var cache iEmail.CaptchaCache
	if err = sCache.NewEmailCache().Get(key, &cache); err != nil {
		return err
	}
	// 是否超过最大发送次数
	if cache.MaxSendTimes != 0 && cache.SendTimes >= cache.MaxSendTimes {
		return code.CaptchaMaxSendTimes
	}
	// 是否超过最大验证次数
	if cache.MaxTryTimes != 0 && cache.TryTimes >= cache.MaxTryTimes {
		return code.CaptchaMaxTryTimes
	}
	// 解析模板
	data := struct{ Code string }{Code: grand.Digits(6)}
	tmpl := template.Must(template.New(key + ".html").Parse(in.Template))
	var buf bytes.Buffer
	if err = tmpl.ExecuteTemplate(&buf, key+".html", data); err != nil {
		return err
	}
	// 发送邮件
	sendIn := iEmail.EmailSendIn{
		Email:   in.Email,
		Subject: in.Subject,
		Body:    buf.String(),
	}
	if err = NewSystemEmail().Send(sendIn); err != nil {
		return err
	}
	// 缓存验证码
	cache.SendTimes = cache.SendTimes + 1
	cache.MaxSendTimes = in.MaxSendTimes
	cache.MaxTryTimes = in.MaxTryTimes
	cache.Code = data.Code
	if err = sCache.NewEmailCache().Set(key, cache, 7); err != nil {
		return err
	}
	return nil
}

func (s *sCaptcha) Verify(email string, c string, sense iEmail.Sense) (err error) {
	key := s.getKey(email, sense)
	// 获取缓存信息
	var cache iEmail.CaptchaCache
	if err = sCache.NewEmailCache().Get(key, &cache); err != nil {
		return err
	}
	// 如果缓存信息不存在，则提示失效
	if cache.Code == "" {
		return code.CaptchaCodeExpired
	}
	// 是否超过禅师次数
	if cache.TryTimes >= cache.MaxTryTimes {
		return code.CaptchaMaxTryTimes
	}
	// 验证码是否正确
	if cache.Code != c {
		cache.TryTimes = cache.TryTimes + 1
		if err = sCache.NewEmailCache().Set(key, cache, 7); err != nil {
			return err
		}
		return code.CaptchaCodeError
	}
	return err
}

func (s *sCaptcha) Remove(email string, sense iEmail.Sense) (err error) {
	key := s.getKey(email, sense)
	return sCache.NewEmailCache().Remove(key)
}

func (s *sCaptcha) getKey(email string, sense iEmail.Sense) string {
	return sCache.CAPTCHA_PREFIX_KEY + string(sense) + "_" + email
}
