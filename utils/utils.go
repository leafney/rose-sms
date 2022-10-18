/**
 * @Author:      leafney
 * @Date:        2022-10-17 20:25
 * @Project:     rose-sms
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package utils

import "strings"

func IsNotEmpty(s string) bool {
	return len(strings.TrimSpace(s)) > 0
}
