package balance

import (
	"errors"
	"math/rand"
)

// 随机负载均衡
type RandomBalance struct {
	// 当前目标索引值
	curIndex int
	// 目标服务器数组
	addrs []string
}

func (r *RandomBalance) Add(addr ...string) error {
	if len(addr) == 0 {
		return errors.New("addr len 1 at least")
	}
	r.addrs = append(r.addrs, addr...)
	return nil
}

func (r *RandomBalance) Next() string {
	if len(r.addrs) == 0 {
		return ""
	}
	r.curIndex = rand.Intn(len(r.addrs))
	return r.addrs[r.curIndex]
}

func (r *RandomBalance) Get() (string, error) {
	return r.Next(), nil
}
