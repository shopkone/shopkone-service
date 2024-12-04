package sDesign

import (
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
)

func (s *sDesign) BlockUpdate(ctx g.Ctx, req *vo.DesignUpdateBlockReq) (res vo.DesignUpdateBlockRes, err error) {
	// 获取这个数据
	URL := s.host + "/design/template"
	_, err = g.Client().Post(ctx, URL, req)
	return res, err
}
