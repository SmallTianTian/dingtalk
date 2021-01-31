package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiImChatServicegroupUpgradeRequest() *OapiImChatServicegroupUpgradeRequest {
	return &OapiImChatServicegroupUpgradeRequest{}
}

type OapiImChatServicegroupUpgradeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImChatServicegroupUpgradeResponse
	ChatId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImChatServicegroupUpgradeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImChatServicegroupUpgradeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImChatServicegroupUpgradeRequest) SetChatId(chatId2 string) {
	this.ChatId = chatId2
}
func (this *OapiImChatServicegroupUpgradeRequest) GetChatId() string {
	return this.ChatId
}
func (this *OapiImChatServicegroupUpgradeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.im.chat.servicegroup.upgrade"
}
func (this *OapiImChatServicegroupUpgradeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImChatServicegroupUpgradeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImChatServicegroupUpgradeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImChatServicegroupUpgradeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImChatServicegroupUpgradeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chat_id"] = this.ChatId
	return txtParams
}
func (this *OapiImChatServicegroupUpgradeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatId, "chatId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiImChatServicegroupUpgradeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImChatServicegroupUpgradeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImChatServicegroupUpgradeResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
