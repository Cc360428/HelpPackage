/**
 * @Author cc
 * @Date 2021/4/1 10:41
 * @Description $
 **/

package monitor_message

import "testing"

func TestSend(t *testing.T) {
	if err := Send("Error", " testaee", true); err != nil {
		t.Log("Cc360428 2021/1/25 09:57", err.Error())
	}
}
