package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMessageSendToSingleConversationRequest() *OapiMessageSendToSingleConversationRequest {
	return &OapiMessageSendToSingleConversationRequest{}
}

type OapiMessageSendToSingleConversationRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMessageSendToSingleConversationResponse
	Msg             string
	ReceiverUserid  string
	SenderUserid    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMessageSendToSingleConversationRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMessageSendToSingleConversationRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMessageSendToSingleConversationRequest) SetMsg(msg2 string) {
	this.Msg = msg2
}
func (this *OapiMessageSendToSingleConversationRequest) GetMsg() string {
	return this.Msg
}
func (this *OapiMessageSendToSingleConversationRequest) SetReceiverUserid(receiverUserid2 string) {
	this.ReceiverUserid = receiverUserid2
}
func (this *OapiMessageSendToSingleConversationRequest) GetReceiverUserid() string {
	return this.ReceiverUserid
}
func (this *OapiMessageSendToSingleConversationRequest) SetSenderUserid(senderUserid2 string) {
	this.SenderUserid = senderUserid2
}
func (this *OapiMessageSendToSingleConversationRequest) GetSenderUserid() string {
	return this.SenderUserid
}
func (this *OapiMessageSendToSingleConversationRequest) GetApiMethodName() string {
	return "dingtalk.oapi.message.send_to_single_conversation"
}
func (this *OapiMessageSendToSingleConversationRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMessageSendToSingleConversationRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMessageSendToSingleConversationRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMessageSendToSingleConversationRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMessageSendToSingleConversationRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["msg"] = this.Msg
	txtParams["receiver_userid"] = this.ReceiverUserid
	txtParams["sender_userid"] = this.SenderUserid
	return txtParams
}
func (this *OapiMessageSendToSingleConversationRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ReceiverUserid, "receiverUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMessageSendToSingleConversationRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMessageSendToSingleConversationRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMessageSendToSingleConversationResponse struct {
	taobao.TaobaoResponse
	MsgId string `json:"msg_id,omitempty"`
}
