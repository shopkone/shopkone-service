package sAli

import (
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/gogf/gf/v2/encoding/gbase64"
)

type sOss struct {
}

func NewOss() *sOss {
	return &sOss{}
}

/*GetUploadToken 获取上传token*/
func (s *sOss) GetUploadToken(shopUUID string) (token string, err error) {
	// 权限
	policy := map[string]interface{}{
		"Version": "1",
		"Statement": []map[string]interface{}{
			{
				"Effect": "Allow",
				"Action": []string{
					"oss:PutObject",
					"oss:DoMetaQuery",
				},
				"Resource": "acs:oss:*:*:*/" + shopUUID + "/*",
			},
		},
	}
	// 鉴权信息
	assumeRoleRequest := &sts20150401.AssumeRoleRequest{
		RoleArn:         tea.String("acs:ram::1168054989853217:role/shopkone"),
		RoleSessionName: tea.String("shopkone"),
		Policy:          tea.String(convertor.ToString(policy)),
	}
	// 获取临时凭证
	ret, err := client.AssumeRole(assumeRoleRequest)
	access := ret.Body.Credentials
	info := map[string]interface{}{
		"id":     access.AccessKeyId,
		"secret": access.AccessKeySecret,
		"token":  *access.SecurityToken,
	}
	// 返回
	token = gbase64.EncodeToString([]byte(convertor.ToString(info)))
	return token, err
}
