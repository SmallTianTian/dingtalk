package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatServicegroupQueryRequest() *OapiImChatServicegroupQueryRequest {
	return &OapiImChatServicegroupQueryRequest{}
}

type OapiImChatServicegroupQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImChatServicegroupQueryResponse
	ChatId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImChatServicegroupQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatServicegroupQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatServicegroupQueryRequest) SetChatId(chatId2 string) {
	this.ChatId = chatId2
}
func (this *OapiImChatServicegroupQueryRequest) GetChatId() string {
	return this.ChatId
}
func (this *OapiImChatServicegroupQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.servicegroup.query"
}
func (this *OapiImChatServicegroupQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatServicegroupQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatServicegroupQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatServicegroupQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatServicegroupQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chat_id"] = this.ChatId
	return txtParams
}
func (this *OapiImChatServicegroupQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatId, "chatId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatServicegroupQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatServicegroupQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatServicegroupQueryResponse struct {
	taobao.TaobaoResponse
	Errcode int64                     `json:"errcode,omitempty"`
	Errmsg  string                    `json:"errmsg,omitempty"`
	Result  ServiceGroupQueryResponse `json:"result,omitempty"`
	Success bool                      `json:"success,omitempty"`
}
type ServiceGroupQueryResponse struct {
	AdminCount         int64  `json:"admin_count,omitempty"`
	MemberCount        int64  `json:"member_count,omitempty"`
	OpenConversationId string `json:"open_conversation_id,omitempty"`
	Title              string `json:"title,omitempty"`
	Url                string `json:"url,omitempty"`
}
