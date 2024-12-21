package sDesign

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
)

func (s *sDesign) ConfigUpdate(ctx g.Ctx, in vo.DesignConfigUpdateReq) (err error) {
	URL := "/design/config"
	r, err := s.GetClient().Put(ctx, URL, in)
	defer r.Close()
	if err != nil {
		return err
	}
	return err
}
