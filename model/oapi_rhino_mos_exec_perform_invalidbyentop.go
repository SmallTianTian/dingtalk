package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecPerformInvalidbyentopRequest() *OapiRhinoMosExecPerformInvalidbyentopRequest {
	return &OapiRhinoMosExecPerformInvalidbyentopRequest{}
}

type OapiRhinoMosExecPerformInvalidbyentopRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecPerformInvalidbyentopResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosExecPerformInvalidbyentopRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecPerformInvalidbyentopRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecPerformInvalidbyentopRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiRhinoMosExecPerformInvalidbyentopRequest) GetReq() string {
	return this.Req
}
func (this *OapiRhinoMosExecPerformInvalidbyentopRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.perform.invalidbyentop"
}
func (this *OapiRhinoMosExecPerformInvalidbyentopRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecPerformInvalidbyentopRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecPerformInvalidbyentopRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecPerformInvalidbyentopRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecPerformInvalidbyentopRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiRhinoMosExecPerformInvalidbyentopRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecPerformInvalidbyentopRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecPerformInvalidbyentopRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type EntityOperation struct {
	EntityId     int64 `json:"entity_id,omitempty"`
	OperationUid int64 `json:"operation_uid,omitempty"`
}
type SpecificEntitiesAndSpecificOperationsReq struct {
	EntityOperations []EntityOperation `json:"entity_operations,omitempty"`
	EntityType       string            `json:"entity_type,omitempty"`
	OrderId          int64             `json:"order_id,omitempty"`
	TenantId         string            `json:"tenant_id,omitempty"`
	Userid           string            `json:"userid,omitempty"`
}
type OapiRhinoMosExecPerformInvalidbyentopResponse struct {
	taobao.TaobaoResponse
	Model   bool `json:"model,omitempty"`
	Success bool `json:"success,omitempty"`
}
