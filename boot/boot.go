package boot

import (
	"github.com/gogf/gf/frame/g"
	"github/pibigstar/go-gateway/utils/config"
)

func init() {
	// 初始化配置
	err := g.Cfg().GetStruct("cluster", &config.Cluster)
	if err != nil {
		panic(err)
	}
}
