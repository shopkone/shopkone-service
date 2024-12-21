package sDesign

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
)

func (s *sDesign) BlockUpdate(ctx g.Ctx, req *vo.DesignUpdateBlockReq) (res vo.DesignUpdateBlockRes, err error) {
	// 获取这个数据
	URL := "/design/template/block"
	data := g.Map{
		"block_id":   req.BlockID,
		"key":        req.Key,
		"value":      req.Value,
		"part_name":  req.PartName,
		"section_id": req.SectionID,
	}
	r, err := s.GetClient().Put(ctx, URL, data)
	defer r.Close()
	return res, err
}
