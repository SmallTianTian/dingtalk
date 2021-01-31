package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiMicroappListRequest() *OapiMicroappListRequest {
	return &OapiMicroappListRequest{}
}

type OapiMicroappListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMicroappListResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMicroappListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.list"
}
func (this *OapiMicroappListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiMicroappListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiMicroappListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappListResponse struct {
	taobao.TaobaoResponse
	AppList []Applist `json:"appList,omitempty"`
	Errcode int64     `json:"errcode,omitempty"`
	Errmsg  string    `json:"errmsg,omitempty"`
}
