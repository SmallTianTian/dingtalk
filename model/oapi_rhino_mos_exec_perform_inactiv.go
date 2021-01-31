package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRhinoMosExecPerformInactiveRequest() *OapiRhinoMosExecPerformInactiveRequest {
	return &OapiRhinoMosExecPerformInactiveRequest{}
}

type OapiRhinoMosExecPerformInactiveRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecPerformInactiveResponse
	Ids             string
	OrderId         int64
	TenantId        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRhinoMosExecPerformInactiveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecPerformInactiveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecPerformInactiveRequest) SetIds(ids2 string) {
	this.Ids = ids2
}
func (this *OapiRhinoMosExecPerformInactiveRequest) GetIds() string {
	return this.Ids
}
func (this *OapiRhinoMosExecPerformInactiveRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiRhinoMosExecPerformInactiveRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiRhinoMosExecPerformInactiveRequest) SetTenantId(tenantId2 string) {
	this.TenantId = tenantId2
}
func (this *OapiRhinoMosExecPerformInactiveRequest) GetTenantId() string {
	return this.TenantId
}
func (this *OapiRhinoMosExecPerformInactiveRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRhinoMosExecPerformInactiveRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRhinoMosExecPerformInactiveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.perform.inactive"
}
func (this *OapiRhinoMosExecPerformInactiveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecPerformInactiveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecPerformInactiveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecPerformInactiveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecPerformInactiveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["ids"] = this.Ids
	txtParams["order_id"] = this.OrderId
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRhinoMosExecPerformInactiveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Ids, "ids"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRhinoMosExecPerformInactiveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecPerformInactiveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecPerformInactiveResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Model   bool   `json:"model,omitempty"`
	Success bool   `json:"success,omitempty"`
}
