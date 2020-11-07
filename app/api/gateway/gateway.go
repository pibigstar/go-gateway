package gateway

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github/pibigstar/go-gateway/app/consts"
	"github/pibigstar/go-gateway/app/model"
	"github/pibigstar/go-gateway/app/request"
	"github/pibigstar/go-gateway/app/response"
	"github/pibigstar/go-gateway/utils/config"
	"strings"
)

// List godoc
// @Summary 服务列表接口
// @Description 服务列表接口
// @Tags 网关服务接口
// @ID /gateway/list
// @Accept  json
// @Produce  json
// @Param content query string false "模糊查询"
// @Param page query int false "页数"
// @Param size query int false "每页多少个"
// @Success 200 {object} response.Response{data=response.Response} "success"
// @Router /gateway/list [get]
func List(r *ghttp.Request) {
	var req *request.ServiceInfoListReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err)
	}
	if req.Page == nil {
		req.Page = &request.Paginate{
			Page: 1,
			Size: 20,
		}
	}

	infos, total, err := model.MServiceInfoModel.PageList(req)
	if err != nil {
		response.Error(r, err)
	}
	var list []*response.ServiceInfo

	serviceHttpBaseURL := fmt.Sprintf("%s:%d", config.Cluster.Ip, config.Cluster.Port)
	for _, info := range infos {
		var nodeNum int
		loadBalance, err := model.MGatewayServiceLoadBalanceModel.GetByServiceId(info.Id)
		if err == nil {
			ipList := strings.Split(loadBalance.IpList, ",")
			nodeNum = len(ipList)
		}

		var serviceAddr string
		if info.LoadType == consts.LoadTypeHTTP {
			httpRule, err := model.MGatewayServiceHttpRuleModel.GetByServiceId(info.Id)
			if err != nil {
				glog.Error("获取http规则失败, req: %d", info.Id)
				continue
			}
			if httpRule.RuleType == consts.RulePrefixUrl {
				serviceAddr = fmt.Sprintf("%s%s", serviceHttpBaseURL, httpRule.Rule)
			}
			if httpRule.RuleType == consts.RuleDomain {
				serviceAddr = httpRule.Rule
			}
		}

		if info.LoadType == consts.LoadTypeTCP {
			tcpRule, err := model.MGatewayServiceTcpRuleModel.GetByServiceId(info.Id)
			if err != nil {
				glog.Error("获取tcp规则失败, req: %d", info.Id)
				continue
			}
			serviceAddr = fmt.Sprintf("%s:%d", config.Cluster.Ip, tcpRule.Port)
		}

		if info.LoadType == consts.LoadTypeGRPC {
			grpcRule, err := model.MGatewayServiceGrpcRuleModel.GetByServiceId(info.Id)
			if err != nil {
				glog.Error("获取grpc规则失败, req: %d", info.Id)
				continue
			}
			serviceAddr = fmt.Sprintf("%s:%d", config.Cluster.Ip, grpcRule.Port)
		}

		list = append(list, &response.ServiceInfo{
			Id:          info.Id,
			LoadType:    info.LoadType,
			ServiceName: info.ServiceName,
			ServiceDesc: info.ServiceDesc,
			UpdateAt:    info.UpdateAt,
			CreateAt:    info.CreateAt,
			ServiceAddr: serviceAddr,
			TotalNode:   nodeNum,
		})
	}
	resp := &response.ServiceInfoListResp{
		Total: total,
		List:  list,
	}
	response.Success(r, resp)
}

// List godoc
// @Summary 服务详情接口
// @Description 服务详情接口
// @Tags 网关服务接口
// @ID /gateway/detail
// @Accept  json
// @Produce  json
// @Param id query string false "服务ID"
// @Success 200 {object} response.Response{data=response.Response} "success"
// @Router /gateway/detail [get]
func Detail(r *ghttp.Request) {
	var req *request.ServiceDetailReq
	if err := r.Parse(&req); err != nil {
		response.Error(r, err)
	}

	info, err := model.MServiceInfoModel.Get(req.Id)
	if err != nil {
		response.Error(r, err)
	}

	resp := &response.ServiceDetail{
		Id:          info.Id,
		LoadType:    info.LoadType,
		ServiceName: info.ServiceName,
		ServiceDesc: info.ServiceDesc,
	}

	// 获取代理规则
	if httpRule, err := model.MGatewayServiceHttpRuleModel.GetByServiceId(info.Id); err == nil {
		resp.HTTP = httpRule.Rule
	}

	if grpcRule, err := model.MGatewayServiceGrpcRuleModel.GetByServiceId(info.Id); err == nil {
		resp.GRPC = grpcRule.HeaderTransfor
	}

	if tcpRule, err := model.MGatewayServiceTcpRuleModel.GetByServiceId(info.Id); err == nil {
		resp.GRPC = fmt.Sprintf("%s:%d", config.Cluster.Ip, tcpRule.Port)
	}

	response.Success(r, resp)
}

// Stat godoc
// @Summary 服务统计接口
// @Description 服务统计接口
// @Tags 网关服务接口
// @ID /gateway/stat
// @Accept  json
// @Produce  json
// @Param id query string false "服务ID"
// @Success 200 {object} response.Response{data=response.Response} "success"
// @Router /gateway/stat [get]
func Stat(r *ghttp.Request) {
	resp := &response.ServiceStat{
		Today:     []int{1, 2, 3},
		Yesterday: []int{4, 5, 6},
	}
	response.Success(r, resp)
}
