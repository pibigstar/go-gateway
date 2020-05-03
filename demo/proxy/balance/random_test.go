package balance

import (
	"fmt"
	"testing"
)

func TestRandomBalance(t *testing.T) {
	rb := &RandomBalance{}
	rb.Add("127.0.0.1:7001")
	rb.Add("127.0.0.1:7002")
	rb.Add("127.0.0.1:7003")
	rb.Add("127.0.0.1:7004")

	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
}
