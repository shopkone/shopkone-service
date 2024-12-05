package sDesign

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
)

func (s *sDesign) SectionUpdate(ctx g.Ctx, req *vo.DesignSectionUpdateReq) (res vo.DesignSectionUpdateRes, err error) {
	data := g.Map{
		"key":        req.Key,
		"part_name":  req.PartName,
		"section_id": req.SectionID,
		"value":      req.Value,
	}
	_, err = s.GetClient().Put(ctx, "/design/template/section", data)
	return res, err
}
