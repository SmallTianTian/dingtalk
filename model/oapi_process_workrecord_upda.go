package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessWorkrecordUpdateRequest() *OapiProcessWorkrecordUpdateRequest {
	return &OapiProcessWorkrecordUpdateRequest{}
}

type OapiProcessWorkrecordUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessWorkrecordUpdateResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessWorkrecordUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessWorkrecordUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessWorkrecordUpdateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessWorkrecordUpdateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessWorkrecordUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.workrecord.update"
}
func (this *OapiProcessWorkrecordUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessWorkrecordUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessWorkrecordUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessWorkrecordUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessWorkrecordUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessWorkrecordUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessWorkrecordUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessWorkrecordUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessWorkrecordUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
