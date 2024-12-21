package sDesign

import (
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
)

func (s *sDesign) ConfigGet(ctx g.Ctx, req *vo.DesignGetConfigReq) (res vo.DesignGetConfigRes, err error) {
	r, err := s.GetClient().Get(ctx, "/design/config")
	defer r.Close()
	if err != nil {
		return vo.DesignGetConfigRes{}, err
	}
	if err = json.Unmarshal(r.ReadAll(), &res); err != nil {
		return res, err
	}
	return res, err
}
