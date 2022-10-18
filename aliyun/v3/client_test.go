/**
 * @Author:      leafney
 * @Date:        2022-10-17 19:56
 * @Project:     rose-sms
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package v3

import "testing"

func TestSmsClient_Send(t *testing.T) {
	sms := NewSmsClient("", "", "")

	payload := map[string]interface{}{"code": "1234"}
	if err := sms.
		SetSignName("").
		DebugMode().
		Send("", payload); err != nil {
		t.Log(err)
	}

}
