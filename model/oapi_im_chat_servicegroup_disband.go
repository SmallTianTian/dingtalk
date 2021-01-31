package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatServicegroupDisbandRequest() *OapiImChatServicegroupDisbandRequest {
	return &OapiImChatServicegroupDisbandRequest{}
}

type OapiImChatServicegroupDisbandRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImChatServicegroupDisbandResponse
	ChatId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImChatServicegroupDisbandRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatServicegroupDisbandRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatServicegroupDisbandRequest) SetChatId(chatId2 string) {
	this.ChatId = chatId2
}
func (this *OapiImChatServicegroupDisbandRequest) GetChatId() string {
	return this.ChatId
}
func (this *OapiImChatServicegroupDisbandRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.servicegroup.disband"
}
func (this *OapiImChatServicegroupDisbandRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatServicegroupDisbandRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatServicegroupDisbandRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatServicegroupDisbandRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatServicegroupDisbandRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chat_id"] = this.ChatId
	return txtParams
}
func (this *OapiImChatServicegroupDisbandRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatId, "chatId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatServicegroupDisbandRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatServicegroupDisbandRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatServicegroupDisbandResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
