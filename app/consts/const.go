package consts

const (
	UserTokenSessionKey = "user"

	//load_type
	LoadTypeHTTP = 0
	LoadTypeTCP  = 1
	LoadTypeGRPC = 2

	// rule type
	RulePrefixUrl = 0 // url 前缀匹配
	RuleDomain    = 1 // 域名匹配

	//default check setting
	DefaultCheckMethod    = 0
	DefaultCheckTimeout   = 2
	DefaultCheckMaxErrNum = 2
	DefaultCheckInterval  = 5
)
