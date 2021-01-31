package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecPerformQueryRequest() *OapiRhinoMosExecPerformQueryRequest {
	return &OapiRhinoMosExecPerformQueryRequest{}
}

type OapiRhinoMosExecPerformQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiRhinoMosExecPerformQueryResponse
	ActiveCondition   string
	EntityIds         string
	EntityType        string
	OperationUids     string
	OrderId           int64
	PerformStatusList string
	TenantId          string
	TopHttpMethod     string
	TopResponseType   string
	Userid            string
	WorkstationCodes  string
}

func (this *OapiRhinoMosExecPerformQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecPerformQueryRequest) SetActiveCondition(activeCondition2 string) {
	this.ActiveCondition = activeCondition2
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetActiveCondition() string {
	return this.ActiveCondition
}
func (this *OapiRhinoMosExecPerformQueryRequest) SetEntityIds(entityIds2 string) {
	this.EntityIds = entityIds2
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetEntityIds() string {
	return this.EntityIds
}
func (this *OapiRhinoMosExecPerformQueryRequest) SetEntityType(entityType2 string) {
	this.EntityType = entityType2
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetEntityType() string {
	return this.EntityType
}
func (this *OapiRhinoMosExecPerformQueryRequest) SetOperationUids(operationUids2 string) {
	this.OperationUids = operationUids2
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetOperationUids() string {
	return this.OperationUids
}
func (this *OapiRhinoMosExecPerformQueryRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecPerformQueryRequest) SetPerformStatusList(performStatusList2 string) {
	this.PerformStatusList = performStatusList2
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetPerformStatusList() string {
	return this.PerformStatusList
}
func (this *OapiRhinoMosExecPerformQueryRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecPerformQueryRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecPerformQueryRequest) SetWorkstationCodes(workstationCodes2 string) {
	this.WorkstationCodes = workstationCodes2
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetWorkstationCodes() string {
	return this.WorkstationCodes
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.perform.query"
}
func (this *OapiRhinoMosExecPerformQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecPerformQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecPerformQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["active_condition"] = this.ActiveCondition
	txtParams["entity_ids"] = this.EntityIds
	txtParams["entity_type"] = this.EntityType
	txtParams["operation_uids"] = this.OperationUids
	txtParams["order_id"] = this.OrderId
	txtParams["perform_status_list"] = this.PerformStatusList
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["workstation_codes"] = this.WorkstationCodes
	return txtParams
}
func (this *OapiRhinoMosExecPerformQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.EntityIds, 500, "entityIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecPerformQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecPerformQueryResponse struct {
	taobao.TaobaoResponse
	Model   []OperationPerformDto `json:"model,omitempty"`
	Success bool                  `json:"success,omitempty"`
}
