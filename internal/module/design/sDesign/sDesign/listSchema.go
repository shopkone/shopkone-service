package sDesign

import (
	"encoding/json"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/gogf/gf/v2/frame/g"
	"shopkone-service/internal/api/vo"
)

func (s *sDesign) ListSchema(ctx g.Ctx, req *vo.DesignSchemaListReq) (res []vo.DesignSchemaListRes, err error) {
	URL := "http://localhost:3100/design/schemas?type=" + slice.Join(req.Type, ",")
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
