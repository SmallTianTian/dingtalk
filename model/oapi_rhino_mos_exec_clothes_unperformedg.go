package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecClothesUnperformedGetRequest() *OapiRhinoMosExecClothesUnperformedGetRequest {
	return &OapiRhinoMosExecClothesUnperformedGetRequest{}
}

type OapiRhinoMosExecClothesUnperformedGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecClothesUnperformedGetResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosExecClothesUnperformedGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecClothesUnperformedGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecClothesUnperformedGetRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiRhinoMosExecClothesUnperformedGetRequest) GetReq() string {
	return this.Req
}
func (this *OapiRhinoMosExecClothesUnperformedGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.clothes.unperformed.get"
}
func (this *OapiRhinoMosExecClothesUnperformedGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecClothesUnperformedGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecClothesUnperformedGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecClothesUnperformedGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecClothesUnperformedGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiRhinoMosExecClothesUnperformedGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecClothesUnperformedGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecClothesUnperformedGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GetUnPerformedClothesByOperationReq struct {
	BizTypes      []string `json:"biz_types,omitempty"`
	OperationUids []int64  `json:"operation_uids,omitempty"`
	OrderId       int64    `json:"order_id,omitempty"`
	Page          Page     `json:"page,omitempty"`
	SizeCode      string   `json:"size_code,omitempty"`
	Status        []string `json:"status,omitempty"`
	TenantId      string   `json:"tenant_id,omitempty"`
	Userid        string   `json:"userid,omitempty"`
}
type OapiRhinoMosExecClothesUnperformedGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64      `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Model   PageResult `json:"model,omitempty"`
	Success bool       `json:"success,omitempty"`
}
