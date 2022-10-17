/**
 * @Author:      leafney
 * @Date:        2022-10-17 19:50
 * @Project:     rose-sms
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package v1

import "fmt"

type comResponse struct {
	Code      string
	Message   string
	RequestId string
}

func (cr *comResponse) GetError() error {
	if cr.Code == "OK" {
		return nil
	}
	return fmt.Errorf("alisms: code = %s , message = %s", cr.Code, cr.Message)
}

type sendResponse struct {
	comResponse
	BizId string
}
