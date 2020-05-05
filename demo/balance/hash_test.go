package balance

import (
	"fmt"
	"testing"
)

func TestHashBalance(t *testing.T) {
	hb := NewHashBalance(10, nil)
	hb.Add("127.0.0.1:7001") //0
	hb.Add("127.0.0.1:7002") //1
	hb.Add("127.0.0.1:7003") //2
	hb.Add("127.0.0.1:7004") //3

	//url hash
	fmt.Println(hb.Get("http://127.0.0.1:7002/base/getinfo"))
	fmt.Println(hb.Get("http://127.0.0.1:7002/base/error"))
	fmt.Println(hb.Get("http://127.0.0.1:7002/base/getinfo"))
	fmt.Println(hb.Get("http://127.0.0.1:7002/base/changepwd"))

	//ip hash
	fmt.Println(hb.Get("127.0.0.1"))
	fmt.Println(hb.Get("192.168.0.1"))
	fmt.Println(hb.Get("127.0.0.1"))
}
