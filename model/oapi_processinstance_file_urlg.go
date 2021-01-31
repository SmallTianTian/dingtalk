package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessinstanceFileUrlGetRequest() *OapiProcessinstanceFileUrlGetRequest {
	return &OapiProcessinstanceFileUrlGetRequest{}
}

type OapiProcessinstanceFileUrlGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessinstanceFileUrlGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessinstanceFileUrlGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessinstanceFileUrlGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessinstanceFileUrlGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessinstanceFileUrlGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessinstanceFileUrlGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.processinstance.file.url.get"
}
func (this *OapiProcessinstanceFileUrlGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessinstanceFileUrlGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessinstanceFileUrlGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessinstanceFileUrlGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessinstanceFileUrlGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessinstanceFileUrlGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessinstanceFileUrlGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessinstanceFileUrlGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessinstanceFileUrlGetResponse struct {
	taobao.TaobaoResponse
	Result  AppSpaceResponse `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
