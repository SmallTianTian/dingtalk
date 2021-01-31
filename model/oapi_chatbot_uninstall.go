package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiChatbotUninstallRequest() *OapiChatbotUninstallRequest {
	return &OapiChatbotUninstallRequest{}
}

type OapiChatbotUninstallRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiChatbotUninstallResponse
	ChatbotId       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiChatbotUninstallRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatbotUninstallRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatbotUninstallRequest) SetChatbotId(chatbotId2 string) {
	this.ChatbotId = chatbotId2
}
func (this *OapiChatbotUninstallRequest) GetChatbotId() string {
	return this.ChatbotId
}
func (this *OapiChatbotUninstallRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chatbot.uninstall"
}
func (this *OapiChatbotUninstallRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatbotUninstallRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatbotUninstallRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatbotUninstallRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatbotUninstallRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatbot_id"] = this.ChatbotId
	return txtParams
}
func (this *OapiChatbotUninstallRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatbotId, "chatbotId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiChatbotUninstallRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatbotUninstallRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatbotUninstallResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
