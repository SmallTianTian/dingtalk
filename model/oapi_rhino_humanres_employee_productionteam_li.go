package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoHumanresEmployeeProductionteamListRequest() *OapiRhinoHumanresEmployeeProductionteamListRequest {
	return &OapiRhinoHumanresEmployeeProductionteamListRequest{}
}

type OapiRhinoHumanresEmployeeProductionteamListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                             OapiRhinoHumanresEmployeeProductionteamListResponse
	QueryEmployeeProductionTeamParam string
	TopHttpMethod                    string
	TopResponseType                  string
}

func (this *OapiRhinoHumanresEmployeeProductionteamListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoHumanresEmployeeProductionteamListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoHumanresEmployeeProductionteamListRequest) SetQueryEmployeeProductionTeamParam(queryEmployeeProductionTeamParam2 string) {
	this.QueryEmployeeProductionTeamParam = queryEmployeeProductionTeamParam2
}
func (this *OapiRhinoHumanresEmployeeProductionteamListRequest) GetQueryEmployeeProductionTeamParam() string {
	return this.QueryEmployeeProductionTeamParam
}
func (this *OapiRhinoHumanresEmployeeProductionteamListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.humanres.employee.productionteam.list"
}
func (this *OapiRhinoHumanresEmployeeProductionteamListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoHumanresEmployeeProductionteamListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoHumanresEmployeeProductionteamListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoHumanresEmployeeProductionteamListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoHumanresEmployeeProductionteamListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["query_employee_production_team_param"] = this.QueryEmployeeProductionTeamParam
	return txtParams
}
func (this *OapiRhinoHumanresEmployeeProductionteamListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoHumanresEmployeeProductionteamListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoHumanresEmployeeProductionteamListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type QueryCorpEmployeeProductionTeamDto struct {
	IncludeInactive bool    `json:"include_inactive,omitempty"`
	TenantId        string  `json:"tenant_id,omitempty"`
	UicIds          []int64 `json:"uic_ids,omitempty"`
	Userid          string  `json:"userid,omitempty"`
}
type OapiRhinoHumanresEmployeeProductionteamListResponse struct {
	taobao.TaobaoResponse
	CorpEmployeeInfo []Model `json:"corp_employee_info,omitempty"`
	Errcode          int64   `json:"errcode,omitempty"`
	Errmsg           string  `json:"errmsg,omitempty"`
	ExternalMsgInfo  string  `json:"external_msg_info,omitempty"`
	StatusCode       int64   `json:"status_code,omitempty"`
	Success          bool    `json:"success,omitempty"`
}
type ProductionTeamList struct {
	BizId              string `json:"biz_id,omitempty"`
	CapacityType       string `json:"capacity_type,omitempty"`
	Deleted            bool   `json:"deleted,omitempty"`
	EmpNum             int64  `json:"emp_num,omitempty"`
	GroupCode          string `json:"group_code,omitempty"`
	Modifier           string `json:"modifier,omitempty"`
	ProductionTeamCode string `json:"production_team_code,omitempty"`
	ProductionTeamName string `json:"production_team_name,omitempty"`
	TenantId           string `json:"tenant_id,omitempty"`
}
