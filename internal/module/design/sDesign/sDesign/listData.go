package sDesign

import (
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
)

func (s *sDesign) ListData(ctx g.Ctx, page string) (res *vo.DesignDataListRes, err error) {
	URL := "/design/template/list"
	r, err := s.GetClient().Get(ctx, URL, g.Map{"page": page})
	defer r.Close()
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(r.ReadAll(), &res); err != nil {
		return res, err
	}
	return res, err
}
