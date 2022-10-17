/**
 * @Author:      leafney
 * @Date:        2022-10-17 19:55
 * @Project:     rose-sms
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package v2

import "testing"

func TestClient_Send(t *testing.T) {

	sms := NewSmsClient("", "", "")

	sms.SetSignName("").
		SetTemplateCode("").
		SetRegionId("").
		Send("", nil)
}
