package balance

import (
	"fmt"
	"testing"
)

func TestPollingBalance(t *testing.T) {
	pb := &PollingBalance{}
	pb.Add("127.0.0.1:7001")
	pb.Add("127.0.0.1:7002")
	pb.Add("127.0.0.1:7003")
	pb.Add("127.0.0.1:7004")

	fmt.Println(pb.Next())
	fmt.Println(pb.Next())
	fmt.Println(pb.Next())
	fmt.Println(pb.Next())
}
