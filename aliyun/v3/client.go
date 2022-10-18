/**
 * @Author:      leafney
 * @Date:        2022-10-17 19:56
 * @Project:     rose-sms
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package v3

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/leafney/rose-sms/utils"
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
		domain:       "dysmsapi.aliyuncs.com",
	}
}

func (s *SmsClient) Send(phone string, payload map[string]interface{}) error {
	config := &openapi.Config{
		AccessKeyId:     tea.String(s.keyId),
		AccessKeySecret: tea.String(s.keySecret),
	}

	if utils.IsNotEmpty(s.domain) {
		config.SetEndpoint(s.domain)
	}

	if utils.IsNotEmpty(s.regionId) {
		config.SetRegionId(s.regionId)
	}

	client, err := dysmsapi20170525.NewClient(config)
	if err != nil {
		return err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:     tea.String(s.signName),
		TemplateCode: tea.String(s.templateCode),
	}
	sendSmsRequest.SetPhoneNumbers(phone)

	data, _ := json.Marshal(payload)
	sendSmsRequest.SetTemplateParam(string(data))

	runtime := &util.RuntimeOptions{}
	resp, err := client.SendSmsWithOptions(sendSmsRequest, runtime)
	if err != nil {
		if s.debug {
			fmt.Println(err)
		}
		return err
	}

	if s.debug {
		fmt.Println(resp.Body)
	}

	return nil
}

func (s *SmsClient) SetRegionId(regionId string) *SmsClient {
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
