package balance

import (
	"fmt"
	"strconv"
)

// 加权轮询负载均衡
type WeightPollingBalance struct {
	curIndex int
	// 目标服务器数组
	nodes []*WeightNode
	rsw   []int
}

type WeightNode struct {
	addr            string
	weight          int //权重值
	currentWeight   int //节点当前权重
	effectiveWeight int //有效权重
}

func (w *WeightPollingBalance) Add(params ...string) error {
	if len(params) != 2 {
		return fmt.Errorf("params must at lease 2")
	}
	weight, err := strconv.Atoi(params[1])
	if err != nil {
		return err
	}
	node := &WeightNode{addr: params[0], weight: weight}
	node.effectiveWeight = node.weight
	w.nodes = append(w.nodes, node)
	return nil
}

func (w *WeightPollingBalance) Next() string {
	total := 0
	var best *WeightNode
	for i := 0; i < len(w.nodes); i++ {
		w := w.nodes[i]
		//step 1 统计所有有效权重之和
		total += w.effectiveWeight

		//step 2 变更节点临时权重为的节点临时权重+节点有效权重
		w.currentWeight += w.effectiveWeight

		//step 3 有效权重默认与权重相同，通讯异常时-1, 通讯成功+1，直到恢复到weight大小
		if w.effectiveWeight < w.weight {
			w.effectiveWeight++
		}
		//step 4 选择最大临时权重点节点
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
	}
	if best == nil {
		return ""
	}
	//step 5 变更临时权重为 临时权重-有效权重之和
	best.currentWeight -= total
	return best.addr
}

func (w *WeightPollingBalance) Get(key ...string) (string, error) {
	return w.Next(), nil
}
