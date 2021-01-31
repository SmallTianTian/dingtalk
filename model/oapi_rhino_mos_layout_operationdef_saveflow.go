package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosLayoutOperationdefSaveflowRequest() *OapiRhinoMosLayoutOperationdefSaveflowRequest {
	return &OapiRhinoMosLayoutOperationdefSaveflowRequest{}
}

type OapiRhinoMosLayoutOperationdefSaveflowRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosLayoutOperationdefSaveflowResponse
	Active          bool
	FlowVersion     int64
	OperationDefs   string
	OrderId         int64
	Source          string
	TenantId        string
	TmpSave         bool
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) SetActive(active2 bool) {
	this.Active = active2
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetActive() bool {
	return this.Active
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) SetFlowVersion(flowVersion2 int64) {
	this.FlowVersion = flowVersion2
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetFlowVersion() int64 {
	return this.FlowVersion
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) SetOperationDefs(operationDefs2 string) {
	this.OperationDefs = operationDefs2
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetOperationDefs() string {
	return this.OperationDefs
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) SetSource(source2 string) {
	this.Source = source2
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetSource() string {
	return this.Source
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) SetTmpSave(tmpSave2 bool) {
	this.TmpSave = tmpSave2
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetTmpSave() bool {
	return this.TmpSave
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.layout.operationdef.saveflow"
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["active"] = this.Active
	txtParams["flow_version"] = this.FlowVersion
	txtParams["operation_defs"] = this.OperationDefs
	txtParams["order_id"] = this.OrderId
	txtParams["source"] = this.Source
	txtParams["tenant_id"] = this.TenantId
	txtParams["tmp_save"] = this.TmpSave
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Active, "active"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosLayoutOperationdefSaveflowRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosLayoutOperationdefSaveflowResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  int64  `json:"result,omitempty"`
}
