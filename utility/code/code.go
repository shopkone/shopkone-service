package code

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func err(code int, message string) error {
	return gerror.NewCode(gcode.New(code, message, message))
}

func baseError(code int, message string) error {
	return err(1000+code, message)
}
func userError(code int, message string) error {
	return err(2000+code, message)
}
func fileError(code int, message string) error {
	return err(3000+code, message)
}
func locationError(code int, message string) error {
	return err(4000+code, message)
}
func productError(code int, message string) error {
	return err(5000+code, message)
}
func shippingError(code int, message string) error {
	return err(6000+code, message)
}
func taxError(code int, message string) error {
	return err(7000+code, message)
}
func SettingError(code int, message string) error {
	return err(8000+code, message)
}
func CustomerError(code int, message string) error {
	return err(9000+code, message)
}
