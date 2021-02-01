package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpChatbotInstallRequest() *CorpChatbotInstallRequest {
	return &CorpChatbotInstallRequest{}
}

type CorpChatbotInstallRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpChatbotInstallResponse
	ChatbotVo       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpChatbotInstallRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpChatbotInstallRequest) SetChatbotVo(chatbotVo2 string) {
	this.ChatbotVo = chatbotVo2
}
func (this *CorpChatbotInstallRequest) GetChatbotVo() string {
	return this.ChatbotVo
}
func (this *CorpChatbotInstallRequest) GetApiMethodName() string {
	return "dingtalk.corp.chatbot.install"
}
func (this *CorpChatbotInstallRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpChatbotInstallRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpChatbotInstallRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpChatbotInstallRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpChatbotInstallRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpChatbotInstallRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatbot_vo"] = this.ChatbotVo
	return txtParams
}
func (this *CorpChatbotInstallRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpChatbotInstallRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpChatbotInstallRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ChatbotVo struct {
	Authority      int64  `json:"authority,omitempty"`
	BotType        int64  `json:"bot_type,omitempty"`
	Breif          string `json:"breif,omitempty"`
	ChatbotId      string `json:"chatbot_id,omitempty"`
	Description    string `json:"description,omitempty"`
	Function       int64  `json:"function,omitempty"`
	Icon           string `json:"icon,omitempty"`
	IconMdify      int64  `json:"icon_mdify,omitempty"`
	MobileSwitch   int64  `json:"mobile_switch,omitempty"`
	Name           string `json:"name,omitempty"`
	NameModify     int64  `json:"name_modify,omitempty"`
	OtoSupport     int64  `json:"oto_support,omitempty"`
	OutgoingToken  string `json:"outgoing_token,omitempty"`
	OutgoingUrl    string `json:"outgoing_url,omitempty"`
	PreviewMediaId string `json:"preview_media_id,omitempty"`
}
type CorpChatbotInstallResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
