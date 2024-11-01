package sAli

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"shopkone-service/hack"
)

var client *sts20150401.Client

func SetClient() error {
	// 获取配置
	conf, err := hack.GetConfig()
	if err != nil {
		return err
	}
	// 初始化客户端
	client, err = sts20150401.NewClient(&openapi.Config{
		AccessKeyId:     &conf.AliYun.AccessKeyId,
		AccessKeySecret: &conf.AliYun.AccessKeySecret,
		Endpoint:        &conf.AliYun.EndPoint,
	})
	return nil
}
