/**
 * @Author cc
 * @Date 2021/4/1 10:41
 * @Description $
 **/

package monitor_message

import "testing"

func TestSend(t *testing.T) {
	NewDingDing("9aec70a47db2a6bbea12682f113e15d86b66fae93bd9f5d70391557369a55798")
	if err := Send("Error", " PublicIp", true); err != nil {
		t.Log("Cc360428 2021/1/25 09:57", err.Error())
	}
}
