package balance

import (
	"errors"
)

// 轮询负载均衡
type PollingBalance struct {
	curIndex int
	// 目标服务器数组
	addrs []string
}

func (p *PollingBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least")
	}
	p.addrs = append(p.addrs, params...)
	return nil
}

func (p *PollingBalance) Next() string {
	if len(p.addrs) == 0 {
		return ""
	}
	lens := len(p.addrs)
	if p.curIndex >= lens {
		p.curIndex = 0
	}
	curAddr := p.addrs[p.curIndex]
	// 当前索引加一
	p.curIndex = (p.curIndex + 1) % lens
	return curAddr
}

func (p *PollingBalance) Get(key ...string) (string, error) {
	return p.Next(), nil
}
