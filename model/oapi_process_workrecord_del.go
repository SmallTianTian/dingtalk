package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessWorkrecordDeleteRequest() *OapiProcessWorkrecordDeleteRequest {
	return &OapiProcessWorkrecordDeleteRequest{}
}

type OapiProcessWorkrecordDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessWorkrecordDeleteResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessWorkrecordDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessWorkrecordDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessWorkrecordDeleteRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessWorkrecordDeleteRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessWorkrecordDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.workrecord.delete"
}
func (this *OapiProcessWorkrecordDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessWorkrecordDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessWorkrecordDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessWorkrecordDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessWorkrecordDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessWorkrecordDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessWorkrecordDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessWorkrecordDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessWorkrecordDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
