package iEmail

type Sense string

const (
	RegisterSense Sense = "register"
	ResetSense    Sense = "reset-pwd"
)

type CaptchaCache struct {
	SendTimes    uint   // 发送次数
	TryTimes     uint   // 尝试次数
	MaxSendTimes uint   // 最大发送次数
	MaxTryTimes  uint   // 最大尝试次数
	Code         string // 验证码
}

type CaptchaSendIn struct {
	Email        string
	Sense        Sense
	Subject      string
	Template     string
	MaxSendTimes uint // 最大发送次数
	MaxTryTimes  uint // 最大尝试次数
}
