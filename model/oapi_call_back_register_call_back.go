package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCallBackRegisterCallBackRequest() *OapiCallBackRegisterCallBackRequest {
	return &OapiCallBackRegisterCallBackRequest{}
}

type OapiCallBackRegisterCallBackRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCallBackRegisterCallBackResponse
	AesKey          string
	CallBackTag     []string
	Token           string
	TopHttpMethod   string
	TopResponseType string
	Url             string
}

func (this *OapiCallBackRegisterCallBackRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCallBackRegisterCallBackRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCallBackRegisterCallBackRequest) SetAesKey(aesKey2 string) {
	this.AesKey = aesKey2
}
func (this *OapiCallBackRegisterCallBackRequest) GetAesKey() string {
	return this.AesKey
}
func (this *OapiCallBackRegisterCallBackRequest) SetCallBackTag(callBackTag2 []string) {
	this.CallBackTag = callBackTag2
}
func (this *OapiCallBackRegisterCallBackRequest) GetCallBackTag() []string {
	return this.CallBackTag
}
func (this *OapiCallBackRegisterCallBackRequest) SetToken(token2 string) {
	this.Token = token2
}
func (this *OapiCallBackRegisterCallBackRequest) GetToken() string {
	return this.Token
}
func (this *OapiCallBackRegisterCallBackRequest) SetUrl(url2 string) {
	this.Url = url2
}
func (this *OapiCallBackRegisterCallBackRequest) GetUrl() string {
	return this.Url
}
func (this *OapiCallBackRegisterCallBackRequest) GetApiMethodName() string {
	return "dingtalk.oapi.call_back.register_call_back"
}
func (this *OapiCallBackRegisterCallBackRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCallBackRegisterCallBackRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCallBackRegisterCallBackRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCallBackRegisterCallBackRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCallBackRegisterCallBackRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["aes_key"] = this.AesKey
	txtParams["call_back_tag"] = this.CallBackTag
	txtParams["token"] = this.Token
	txtParams["url"] = this.Url
	return txtParams
}
func (this *OapiCallBackRegisterCallBackRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCallBackRegisterCallBackRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCallBackRegisterCallBackRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCallBackRegisterCallBackResponse struct {
	taobao.TaobaoResponse
}
