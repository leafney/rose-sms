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
	keyId        string
	keySecret    string
	templateCode string
	debug        bool

	signName string

	regionId string
	domain   string
}

func NewSmsClient(keyId, keySecret, templateCode string) *SmsClient {

	return &SmsClient{
		keyId:        keyId,
		keySecret:    keySecret,
		templateCode: templateCode,
		regionId:     "cn-hangzhou",
	}
}

func (s *SmsClient) Send(phone string, payload map[string]interface{}) error {
	client, err := dysmsapi.NewClientWithAccessKey(s.regionId, s.keyId, s.keySecret)
	if err != nil {
		return err
	}

	request := dysmsapi.CreateSendSmsRequest()

	request.Scheme = "https"
	request.Domain = "dysmsapi.aliyuncs.com"

	request.SignName = s.signName
	request.TemplateCode = s.templateCode

	request.PhoneNumbers = phone
	data, _ := json.Marshal(payload)
	request.TemplateParam = string(data)

	response, err := client.SendSms(request)
	if err != nil {
		if s.debug {
			fmt.Println(err)
		}
		return err
	}

	if s.debug {
		fmt.Println(response)
	}

	return nil
}

func (s *SmsClient) SetRegionId(regionId string) *SmsClient {
	// - [dysmsapi 在某种情况下会请求错误地址 · Issue #492 · aliyun/alibaba-cloud-sdk-go](https://github.com/aliyun/alibaba-cloud-sdk-go/issues/492)
	if regionId == "cn-beijing" {
		s.domain = "dysmsapi.aliyuncs.com"
	}
	s.regionId = regionId
	return s
}

func (s *SmsClient) SetSignName(signName string) *SmsClient {
	s.signName = signName
	return s
}

func (s *SmsClient) SetTemplateCode(code string) *SmsClient {
	s.templateCode = code
	return s
}

func (s *SmsClient) SetDomain(domain string) *SmsClient {
	s.domain = domain
	return s
}

func (s *SmsClient) DebugMode() *SmsClient {
	s.debug = true
	return s
}
