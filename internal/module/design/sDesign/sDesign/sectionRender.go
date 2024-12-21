package sDesign

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
)

func (s *sDesign) SectionRender(ctx g.Ctx, req *vo.DesignSectionRenderReq) (html string, err error) {
	data := g.Map{
		"name":       req.PartName,
		"section_id": req.SectionID,
		"shop_id":    s.shopId,
	}
	r, err := s.GetClient().Get(ctx, "/section", data)
	defer r.Close()
	if err != nil {
		return "", err
	}
	return r.ReadAllString(), err
}
