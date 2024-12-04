package sDesign

import (
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
)

func (s *sDesign) ListData(ctx g.Ctx) (res *vo.DesignDataListRes, err error) {
	URL := s.host + "/design/templates?shop_id=5"
	g.Dump(URL)
	r, err := g.Client().Get(ctx, URL)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	if err = json.Unmarshal(r.ReadAll(), &res); err != nil {
		return res, err
	}
	return res, err
}
