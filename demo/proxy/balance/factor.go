package balance

type BalanceType int

const (
	RandomType BalanceType = iota
	PollingType
	WeightPollingType
	HashType
)

// 工厂方法，返回指定负载均衡策略
func LoadBalanceFactory(balanceType BalanceType) Balance {
	switch balanceType {
	case RandomType:
		return &RandomBalance{}
	case HashType:
		return NewHashBalance(10, nil)
	case PollingType:
		return &PollingBalance{}
	case WeightPollingType:
		return &WeightPollingBalance{}
	default:
		return &RandomBalance{}
	}
}
