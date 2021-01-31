package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiPbpEventSyncRequest() *OapiPbpEventSyncRequest {
	return &OapiPbpEventSyncRequest{}
}

type OapiPbpEventSyncRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiPbpEventSyncResponse
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiPbpEventSyncRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiPbpEventSyncRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiPbpEventSyncRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiPbpEventSyncRequest) GetParam() string {
	return this.Param
}
func (this *OapiPbpEventSyncRequest) GetApiMethodName() string {
	return "dingtalk.oapi.pbp.event.sync"
}
func (this *OapiPbpEventSyncRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiPbpEventSyncRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiPbpEventSyncRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiPbpEventSyncRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiPbpEventSyncRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiPbpEventSyncRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiPbpEventSyncRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiPbpEventSyncRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiPbpEventSyncResponse struct {
	taobao.TaobaoResponse
}
