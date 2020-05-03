package balance

import (
	"fmt"
	"testing"
)

func TestWeightPollingBalance(t *testing.T) {
	wpb := &WeightPollingBalance{}
	wpb.Add("127.0.0.1:7001", 2) // 出现2次
	wpb.Add("127.0.0.1:7002", 1) // 出现1次
	wpb.Add("127.0.0.1:7003", 2) // 出现1次
	wpb.Add("127.0.0.1:7004", 1) // 出现2次

	fmt.Println(wpb.Next())
	fmt.Println(wpb.Next())
	fmt.Println(wpb.Next())
	fmt.Println(wpb.Next())
	fmt.Println(wpb.Next())
	fmt.Println(wpb.Next())
}
