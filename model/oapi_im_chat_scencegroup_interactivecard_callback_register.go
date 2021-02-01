package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatScencegroupInteractivecardCallbackRegisterRequest() *OapiImChatScencegroupInteractivecardCallbackRegisterRequest {
	return &OapiImChatScencegroupInteractivecardCallbackRegisterRequest{}
}

type OapiImChatScencegroupInteractivecardCallbackRegisterRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImChatScencegroupInteractivecardCallbackRegisterResponse
	ApiSecret       string
	CallbackUrl     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) SetApiSecret(apiSecret2 string) {
	this.ApiSecret = apiSecret2
}
func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) GetApiSecret() string {
	return this.ApiSecret
}
func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) SetCallbackUrl(callbackUrl2 string) {
	this.CallbackUrl = callbackUrl2
}
func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) GetCallbackUrl() string {
	return this.CallbackUrl
}
func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.scencegroup.interactivecard.callback.register"
}
func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["api_secret"] = this.ApiSecret
	txtParams["callback_url"] = this.CallbackUrl
	return txtParams
}
func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CallbackUrl, "callbackUrl"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatScencegroupInteractivecardCallbackRegisterRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatScencegroupInteractivecardCallbackRegisterResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
