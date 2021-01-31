package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosLayoutOperationdefGetflowRequest() *OapiRhinoMosLayoutOperationdefGetflowRequest {
	return &OapiRhinoMosLayoutOperationdefGetflowRequest{}
}

type OapiRhinoMosLayoutOperationdefGetflowRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosLayoutOperationdefGetflowResponse
	FlowVersion     int64
	NeedAssignInfo  bool
	OrderId         int64
	TenantId        string
	TmpSave         bool
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) SetFlowVersion(flowVersion2 int64) {
	this.FlowVersion = flowVersion2
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) GetFlowVersion() int64 {
	return this.FlowVersion
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) SetNeedAssignInfo(needAssignInfo2 bool) {
	this.NeedAssignInfo = needAssignInfo2
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) GetNeedAssignInfo() bool {
	return this.NeedAssignInfo
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) SetTmpSave(tmpSave2 bool) {
	this.TmpSave = tmpSave2
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) GetTmpSave() bool {
	return this.TmpSave
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.layout.operationdef.getflow"
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["flow_version"] = this.FlowVersion
	txtParams["need_assign_info"] = this.NeedAssignInfo
	txtParams["order_id"] = this.OrderId
	txtParams["tenant_id"] = this.TenantId
	txtParams["tmp_save"] = this.TmpSave
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.NeedAssignInfo, "needAssignInfo"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosLayoutOperationdefGetflowRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosLayoutOperationdefGetflowResponse struct {
	taobao.TaobaoResponse
	Errcode int64               `json:"errcode,omitempty"`
	Errmsg  string              `json:"errmsg,omitempty"`
	Result  OperationDefFlowDto `json:"result,omitempty"`
}
type OperationDefFlowDto struct {
	Active        bool              `json:"active,omitempty"`
	FlowVersion   int64             `json:"flow_version,omitempty"`
	OperationDefs []OperationDefDto `json:"operation_defs,omitempty"`
	OrderId       int64             `json:"order_id,omitempty"`
	Source        string            `json:"source,omitempty"`
	TenantId      string            `json:"tenant_id,omitempty"`
}
