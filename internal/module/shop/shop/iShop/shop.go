package iShop

import "github.com/gogf/gf/v2/frame/g"

type CreateTrialIn struct {
	Email   string
	UserId  uint
	Country string
	Zone    string
	Ctx     g.Ctx
}
