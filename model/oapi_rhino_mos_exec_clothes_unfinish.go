package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRhinoMosExecClothesUnfinishRequest() *OapiRhinoMosExecClothesUnfinishRequest {
	return &OapiRhinoMosExecClothesUnfinishRequest{}
}

type OapiRhinoMosExecClothesUnfinishRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRhinoMosExecClothesUnfinishResponse
	Req             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRhinoMosExecClothesUnfinishRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRhinoMosExecClothesUnfinishRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRhinoMosExecClothesUnfinishRequest) SetReq(req2 string) {
	this.Req = req2
}
func (this *OapiRhinoMosExecClothesUnfinishRequest) GetReq() string {
	return this.Req
}
func (this *OapiRhinoMosExecClothesUnfinishRequest) GetApiMethodName() string {
	return "dingtalk.oapi.rhino.mos.exec.clothes.unfinish"
}
func (this *OapiRhinoMosExecClothesUnfinishRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRhinoMosExecClothesUnfinishRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRhinoMosExecClothesUnfinishRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRhinoMosExecClothesUnfinishRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRhinoMosExecClothesUnfinishRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["req"] = this.Req
	return txtParams
}
func (this *OapiRhinoMosExecClothesUnfinishRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRhinoMosExecClothesUnfinishRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRhinoMosExecClothesUnfinishRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRhinoMosExecClothesUnfinishResponse struct {
	taobao.TaobaoResponse
	Model   bool `json:"model,omitempty"`
	Success bool `json:"success,omitempty"`
}
