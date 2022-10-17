/**
 * @Author:      leafney
 * @Date:        2022-10-17 19:55
 * @Project:     rose-sms
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package v2

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type SmsClient struct {
	KeyId        string
	KeySecret    string
	TemplateCode string

	SignName string

	RegionId string
	Domain   string
}

func NewSmsClient(keyId, keySecret, templateCode string) *SmsClient {

	return &SmsClient{
		KeyId:        keyId,
		KeySecret:    keySecret,
		TemplateCode: templateCode,
		RegionId:     "cn-hangzhou",
	}
}

func (s *SmsClient) Send(phone string, payload map[string]interface{}) error {
	client, err := dysmsapi.NewClientWithAccessKey(s.RegionId, s.KeyId, s.KeySecret)
	if err != nil {
		return err
	}

	request := dysmsapi.CreateSendSmsRequest()

	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"

	request.SignName = s.SignName
	request.TemplateCode = s.TemplateCode

	request.PhoneNumbers = phone
	data, _ := json.Marshal(payload)
	request.TemplateParam = string(data)

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)

	return nil
}

func (s *SmsClient) SetRegionId(regionId string) *SmsClient {
	// - [dysmsapi 在某种情况下会请求错误地址 · Issue #492 · aliyun/alibaba-cloud-sdk-go](https://github.com/aliyun/alibaba-cloud-sdk-go/issues/492)
	if regionId == "cn-beijing" {
		s.Domain = "dysmsapi.aliyuncs.com"
	}
	s.RegionId = regionId
	return s
}

func (s *SmsClient) SetSignName(signName string) *SmsClient {
	s.SignName = signName
	return s
}

func (s *SmsClient) SetTemplateCode(code string) *SmsClient {
	s.TemplateCode = code
	return s
}

func (s *SmsClient) SetDomain(domain string) *SmsClient {
	s.Domain = domain
	return s
}
