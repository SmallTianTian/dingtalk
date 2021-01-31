package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpChatbotListbychatbotidsRequest() *CorpChatbotListbychatbotidsRequest {
	return &CorpChatbotListbychatbotidsRequest{}
}

type CorpChatbotListbychatbotidsRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpChatbotListbychatbotidsResponse
	ChatbotIds      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpChatbotListbychatbotidsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpChatbotListbychatbotidsRequest) SetChatbotIds(chatbotIds2 string) {
	this.ChatbotIds = chatbotIds2
}
func (this *CorpChatbotListbychatbotidsRequest) GetChatbotIds() string {
	return this.ChatbotIds
}
func (this *CorpChatbotListbychatbotidsRequest) GetApiMethodName() string {
	return "dingtalk.corp.chatbot.listbychatbotids"
}
func (this *CorpChatbotListbychatbotidsRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpChatbotListbychatbotidsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpChatbotListbychatbotidsRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpChatbotListbychatbotidsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpChatbotListbychatbotidsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpChatbotListbychatbotidsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatbot_ids"] = this.ChatbotIds
	return txtParams
}
func (this *CorpChatbotListbychatbotidsRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatbotIds, "chatbotIds"); err != nil {
		return err
	}
	return nil
}
func (this *CorpChatbotListbychatbotidsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpChatbotListbychatbotidsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpChatbotListbychatbotidsResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type ChatbotResultVo struct {
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
