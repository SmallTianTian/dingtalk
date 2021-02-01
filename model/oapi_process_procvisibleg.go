package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessProcvisibleGetRequest() *OapiProcessProcvisibleGetRequest {
	return &OapiProcessProcvisibleGetRequest{}
}

type OapiProcessProcvisibleGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessProcvisibleGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessProcvisibleGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessProcvisibleGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessProcvisibleGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessProcvisibleGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessProcvisibleGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.procvisible.get"
}
func (this *OapiProcessProcvisibleGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessProcvisibleGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessProcvisibleGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessProcvisibleGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessProcvisibleGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessProcvisibleGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessProcvisibleGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessProcvisibleGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessProcvisibleGetResponse struct {
	taobao.TaobaoResponse
	Result  []Result `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
