package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecClothesUnperformedFilterRequest() *OapiRhinoMosExecClothesUnperformedFilterRequest {
	return &OapiRhinoMosExecClothesUnperformedFilterRequest{}
}

type OapiRhinoMosExecClothesUnperformedFilterRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecClothesUnperformedFilterResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosExecClothesUnperformedFilterRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecClothesUnperformedFilterRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecClothesUnperformedFilterRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiRhinoMosExecClothesUnperformedFilterRequest) GetReq() string {
	return this.Req
}
func (this *OapiRhinoMosExecClothesUnperformedFilterRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.clothes.unperformed.filter"
}
func (this *OapiRhinoMosExecClothesUnperformedFilterRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecClothesUnperformedFilterRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecClothesUnperformedFilterRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecClothesUnperformedFilterRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecClothesUnperformedFilterRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiRhinoMosExecClothesUnperformedFilterRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecClothesUnperformedFilterRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecClothesUnperformedFilterRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type FilterUnPerformedClothesReq struct {
	EntityIds     []int64 `json:"entity_ids,omitempty"`
	OperationUids []int64 `json:"operation_uids,omitempty"`
	OrderId       int64   `json:"order_id,omitempty"`
	TenantId      string  `json:"tenant_id,omitempty"`
	Userid        string  `json:"userid,omitempty"`
}
type OapiRhinoMosExecClothesUnperformedFilterResponse struct {
	taobao.TaobaoResponse
	Model   []int64 `json:"model,omitempty"`
	Success bool    `json:"success,omitempty"`
}
