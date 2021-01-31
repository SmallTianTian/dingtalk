package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoHumanresEmployeeProcessCapacityQueryRequest() *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest {
	return &OapiRhinoHumanresEmployeeProcessCapacityQueryRequest{}
}

type OapiRhinoHumanresEmployeeProcessCapacityQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                           OapiRhinoHumanresEmployeeProcessCapacityQueryResponse
	Category                       string
	ProcessStructuralClusterIdList string
	TenantId                       string
	TopHttpMethod                  string
	TopResponseType                string
	Userid                         string
	WorkNos                        string
}

func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) SetCategory(category2 string) {
	this.Category = category2
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) GetCategory() string {
	return this.Category
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) SetProcessStructuralClusterIdList(processStructuralClusterIdList2 string) {
	this.ProcessStructuralClusterIdList = processStructuralClusterIdList2
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) GetProcessStructuralClusterIdList() string {
	return this.ProcessStructuralClusterIdList
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) SetWorkNos(workNos2 string) {
	this.WorkNos = workNos2
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) GetWorkNos() string {
	return this.WorkNos
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.humanres.employee.process.capacity.query"
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["category"] = this.Category
	txtParams["process_structural_cluster_id_list"] = this.ProcessStructuralClusterIdList
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["work_nos"] = this.WorkNos
	return txtParams
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.ProcessStructuralClusterIdList, 20, "processStructuralClusterIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoHumanresEmployeeProcessCapacityQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoHumanresEmployeeProcessCapacityQueryResponse struct {
	taobao.TaobaoResponse
	Errcode         int64                                 `json:"errcode,omitempty"`
	Errmsg          string                                `json:"errmsg,omitempty"`
	ExternalMsgInfo string                                `json:"externalMsgInfo,omitempty"`
	Hsfcode         int64                                 `json:"hsfcode,omitempty"`
	Model           []EmployeeProcessCapacityStatisticDto `json:"model,omitempty"`
	Success         bool                                  `json:"success,omitempty"`
}
type EmployeeProcessCapacityStatisticDto struct {
	Category                   string `json:"category,omitempty"`
	ProcessCapacityId          int64  `json:"process_capacity_id,omitempty"`
	ProcessStructuralClusterId int64  `json:"process_structural_cluster_id,omitempty"`
	ProductionQuantity         int64  `json:"production_quantity,omitempty"`
	StatisticCategory          int64  `json:"statistic_category,omitempty"`
	StatisticsValue            string `json:"statistics_value,omitempty"`
	TenantId                   string `json:"tenant_id,omitempty"`
	WorkNo                     string `json:"work_no,omitempty"`
}
