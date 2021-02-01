package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiChatbotInstallRequest() *OapiChatbotInstallRequest {
	return &OapiChatbotInstallRequest{}
}

type OapiChatbotInstallRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatbotInstallResponse
	ChatbotVo       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiChatbotInstallRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatbotInstallRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatbotInstallRequest) SetChatbotVo(chatbotVo2 string) {
	this.ChatbotVo = chatbotVo2
}
func (this *OapiChatbotInstallRequest) GetChatbotVo() string {
	return this.ChatbotVo
}
func (this *OapiChatbotInstallRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chatbot.install"
}
func (this *OapiChatbotInstallRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatbotInstallRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatbotInstallRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatbotInstallRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatbotInstallRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatbot_vo"] = this.ChatbotVo
	return txtParams
}
func (this *OapiChatbotInstallRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiChatbotInstallRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatbotInstallRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatbotInstallResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
