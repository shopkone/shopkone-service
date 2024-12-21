package sDesign

import (
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sDesign) DesignInit(ctx g.Ctx) (err error) {
	URL := "/design/template/init"
	r, err := s.GetClient().Post(ctx, URL)
	defer r.Close()
	return err
}
