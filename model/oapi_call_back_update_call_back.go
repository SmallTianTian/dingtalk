package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCallBackUpdateCallBackRequest() *OapiCallBackUpdateCallBackRequest {
	return &OapiCallBackUpdateCallBackRequest{}
}

type OapiCallBackUpdateCallBackRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCallBackUpdateCallBackResponse
	AesKey          string
	CallBackTag     []string
	Token           string
	TopHttpMethod   string
	TopResponseType string
	Url             string
}

func (this *OapiCallBackUpdateCallBackRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCallBackUpdateCallBackRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCallBackUpdateCallBackRequest) SetAesKey(aesKey2 string) {
	this.AesKey = aesKey2
}
func (this *OapiCallBackUpdateCallBackRequest) GetAesKey() string {
	return this.AesKey
}
func (this *OapiCallBackUpdateCallBackRequest) SetCallBackTag(callBackTag2 []string) {
	this.CallBackTag = callBackTag2
}
func (this *OapiCallBackUpdateCallBackRequest) GetCallBackTag() []string {
	return this.CallBackTag
}
func (this *OapiCallBackUpdateCallBackRequest) SetToken(token2 string) {
	this.Token = token2
}
func (this *OapiCallBackUpdateCallBackRequest) GetToken() string {
	return this.Token
}
func (this *OapiCallBackUpdateCallBackRequest) SetUrl(url2 string) {
	this.Url = url2
}
func (this *OapiCallBackUpdateCallBackRequest) GetUrl() string {
	return this.Url
}
func (this *OapiCallBackUpdateCallBackRequest) GetApiMethodName() string {
	return "dingtalk.oapi.call_back.update_call_back"
}
func (this *OapiCallBackUpdateCallBackRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCallBackUpdateCallBackRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCallBackUpdateCallBackRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCallBackUpdateCallBackRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCallBackUpdateCallBackRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["aes_key"] = this.AesKey
	txtParams["call_back_tag"] = this.CallBackTag
	txtParams["token"] = this.Token
	txtParams["url"] = this.Url
	return txtParams
}
func (this *OapiCallBackUpdateCallBackRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCallBackUpdateCallBackRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCallBackUpdateCallBackRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCallBackUpdateCallBackResponse struct {
	taobao.TaobaoResponse
}
