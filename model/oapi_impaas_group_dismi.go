package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasGroupDismissRequest() *OapiImpaasGroupDismissRequest {
	return &OapiImpaasGroupDismissRequest{}
}

type OapiImpaasGroupDismissRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasGroupDismissResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasGroupDismissRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasGroupDismissRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasGroupDismissRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasGroupDismissRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasGroupDismissRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.group.dismiss"
}
func (this *OapiImpaasGroupDismissRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasGroupDismissRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasGroupDismissRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasGroupDismissRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasGroupDismissRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasGroupDismissRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasGroupDismissRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasGroupDismissRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DismissGroupRequest struct {
	Chatid string `json:"chatid,omitempty"`
}
type OapiImpaasGroupDismissResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
