package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasNewretailSendstaffgroupmessageRequest() *OapiImpaasNewretailSendstaffgroupmessageRequest {
	return &OapiImpaasNewretailSendstaffgroupmessageRequest{}
}

type OapiImpaasNewretailSendstaffgroupmessageRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasNewretailSendstaffgroupmessageResponse
	ChatId          string
	Content         string
	MsgType         int64
	Sender          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) SetChatId(chatId2 string) {
	this.ChatId = chatId2
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) GetChatId() string {
	return this.ChatId
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) SetContent(content2 string) {
	this.Content = content2
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) SetContentString(content2 string) {
	this.Content = content2
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) GetContent() string {
	return this.Content
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) SetMsgType(msgType2 int64) {
	this.MsgType = msgType2
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) GetMsgType() int64 {
	return this.MsgType
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) SetSender(sender2 string) {
	this.Sender = sender2
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) GetSender() string {
	return this.Sender
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.newretail.sendstaffgroupmessage"
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chat_id"] = this.ChatId
	txtParams[taobao.DATA_CONTENT] = this.Content
	txtParams["msg_type"] = this.MsgType
	txtParams["sender"] = this.Sender
	return txtParams
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasNewretailSendstaffgroupmessageRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImpaasNewretailSendstaffgroupmessageResponse struct {
	taobao.TaobaoResponse
	Result int64 `json:"result,omitempty"`
}
