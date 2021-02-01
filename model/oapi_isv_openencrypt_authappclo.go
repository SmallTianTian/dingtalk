package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiIsvOpenencryptAuthappcloseRequest() *OapiIsvOpenencryptAuthappcloseRequest {
	return &OapiIsvOpenencryptAuthappcloseRequest{}
}

type OapiIsvOpenencryptAuthappcloseRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                 OapiIsvOpenencryptAuthappcloseResponse
	TopAuthMicroAppClose string
	TopHttpMethod        string
	TopResponseType      string
}

func (this *OapiIsvOpenencryptAuthappcloseRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiIsvOpenencryptAuthappcloseRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiIsvOpenencryptAuthappcloseRequest) SetTopAuthMicroAppClose(topAuthMicroAppClose2 string) {
	this.TopAuthMicroAppClose = topAuthMicroAppClose2
}
func (this *OapiIsvOpenencryptAuthappcloseRequest) GetTopAuthMicroAppClose() string {
	return this.TopAuthMicroAppClose
}
func (this *OapiIsvOpenencryptAuthappcloseRequest) GetApiMethodName() string {
	return "dingtalk.oapi.isv.openencrypt.authappclose"
}
func (this *OapiIsvOpenencryptAuthappcloseRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiIsvOpenencryptAuthappcloseRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiIsvOpenencryptAuthappcloseRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiIsvOpenencryptAuthappcloseRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiIsvOpenencryptAuthappcloseRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["top_auth_micro_app_close"] = this.TopAuthMicroAppClose
	return txtParams
}
func (this *OapiIsvOpenencryptAuthappcloseRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiIsvOpenencryptAuthappcloseRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiIsvOpenencryptAuthappcloseRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopAuthMicroAppClose struct {
	Appid     int64  `json:"appid,omitempty"`
	Authcode  string `json:"authcode,omitempty"`
	Requestid string `json:"requestid,omitempty"`
	Resource  string `json:"resource,omitempty"`
}
type OapiIsvOpenencryptAuthappcloseResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
