package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiServicegroupMessageSendRequest() *OapiServicegroupMessageSendRequest {
	return &OapiServicegroupMessageSendRequest{}
}

type OapiServicegroupMessageSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                OapiServicegroupMessageSendResponse
	ConversationMessage string
	OrderId             int64
	TopHttpMethod       string
	TopResponseType     string
}

func (this *OapiServicegroupMessageSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServicegroupMessageSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServicegroupMessageSendRequest) SetConversationMessage(conversationMessage2 string) {
	this.ConversationMessage = conversationMessage2
}
func (this *OapiServicegroupMessageSendRequest) GetConversationMessage() string {
	return this.ConversationMessage
}
func (this *OapiServicegroupMessageSendRequest) SetOrderId(orderId2 int64) {
	this.OrderId = orderId2
}
func (this *OapiServicegroupMessageSendRequest) GetOrderId() int64 {
	return this.OrderId
}
func (this *OapiServicegroupMessageSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.servicegroup.message.send"
}
func (this *OapiServicegroupMessageSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServicegroupMessageSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServicegroupMessageSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServicegroupMessageSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServicegroupMessageSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["conversation_message"] = this.ConversationMessage
	txtParams["order_id"] = this.OrderId
	return txtParams
}
func (this *OapiServicegroupMessageSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OrderId, "orderId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiServicegroupMessageSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServicegroupMessageSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ConversationMessage struct {
	Text  string `json:"text,omitempty"`
	Title string `json:"title,omitempty"`
}
type OapiServicegroupMessageSendResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
