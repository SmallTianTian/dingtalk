package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiOrgOpenencryptAuthappcloseRequest() *OapiOrgOpenencryptAuthappcloseRequest {
	return &OapiOrgOpenencryptAuthappcloseRequest{}
}

type OapiOrgOpenencryptAuthappcloseRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                 OapiOrgOpenencryptAuthappcloseResponse
	TopAuthMicroAppClose string
	TopHttpMethod        string
	TopResponseType      string
}

func (this *OapiOrgOpenencryptAuthappcloseRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOrgOpenencryptAuthappcloseRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOrgOpenencryptAuthappcloseRequest) SetTopAuthMicroAppClose(topAuthMicroAppClose2 string) {
	this.TopAuthMicroAppClose = topAuthMicroAppClose2
}
func (this *OapiOrgOpenencryptAuthappcloseRequest) GetTopAuthMicroAppClose() string {
	return this.TopAuthMicroAppClose
}
func (this *OapiOrgOpenencryptAuthappcloseRequest) GetApiMethodName() string {
	return "dingtalk.oapi.org.openencrypt.authappclose"
}
func (this *OapiOrgOpenencryptAuthappcloseRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOrgOpenencryptAuthappcloseRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOrgOpenencryptAuthappcloseRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOrgOpenencryptAuthappcloseRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOrgOpenencryptAuthappcloseRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["top_auth_micro_app_close"] = this.TopAuthMicroAppClose
	return txtParams
}
func (this *OapiOrgOpenencryptAuthappcloseRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiOrgOpenencryptAuthappcloseRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOrgOpenencryptAuthappcloseRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiOrgOpenencryptAuthappcloseResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
