/**
 * @Author:      leafney
 * @Date:        2022-10-17 19:54
 * @Project:     rose-sms
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package v1

import "testing"

func TestClient_Send(t *testing.T) {

	client, err := New("", "", SignName(""), Template(""))
	if err != nil {
		t.Error(err)
	}

	if err := client.Send(Mobile(""), Parameter(map[string]string{"code": "1234"})); err != nil {
		t.Log(err)
	}
	t.Log("success")
}
