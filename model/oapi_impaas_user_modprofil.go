package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasUserModprofileRequest() *OapiImpaasUserModprofileRequest {
	return &OapiImpaasUserModprofileRequest{}
}

type OapiImpaasUserModprofileRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasUserModprofileResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasUserModprofileRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasUserModprofileRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasUserModprofileRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasUserModprofileRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasUserModprofileRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.user.modprofile"
}
func (this *OapiImpaasUserModprofileRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasUserModprofileRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasUserModprofileRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasUserModprofileRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasUserModprofileRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasUserModprofileRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasUserModprofileRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasUserModprofileRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ModProfileReq struct {
	AccountInfo   AccountInfo `json:"account_info,omitempty"`
	AvatarMediaid string      `json:"avatar_mediaid,omitempty"`
	Extension     string      `json:"extension,omitempty"`
	Nick          string      `json:"nick,omitempty"`
}
type OapiImpaasUserModprofileResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
