package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatbotGetRequest() *OapiImChatbotGetRequest {
	return &OapiImChatbotGetRequest{}
}

type OapiImChatbotGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiImChatbotGetResponse
	OpenConversationId string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiImChatbotGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatbotGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatbotGetRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiImChatbotGetRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiImChatbotGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chatbot.get"
}
func (this *OapiImChatbotGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatbotGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatbotGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatbotGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatbotGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_conversation_id"] = this.OpenConversationId
	return txtParams
}
func (this *OapiImChatbotGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpenConversationId, "openConversationId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatbotGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatbotGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatbotGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64               `json:"errcode,omitempty"`
	Errmsg  string              `json:"errmsg,omitempty"`
	Result  []ChatbotInstanceVO `json:"result,omitempty"`
}
type ChatbotInstanceVO struct {
	ChatbotUserId string `json:"chatbot_user_id,omitempty"`
}
