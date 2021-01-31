package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDingtalkImpaasMessageCrossdomainSendRequest() *OapiDingtalkImpaasMessageCrossdomainSendRequest {
	return &OapiDingtalkImpaasMessageCrossdomainSendRequest{}
}

type OapiDingtalkImpaasMessageCrossdomainSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                        OapiDingtalkImpaasMessageCrossdomainSendResponse
	CrossdomainMessageSendModel string
	TopHttpMethod               string
	TopResponseType             string
}

func (this *OapiDingtalkImpaasMessageCrossdomainSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingtalkImpaasMessageCrossdomainSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingtalkImpaasMessageCrossdomainSendRequest) SetCrossdomainMessageSendModel(crossdomainMessageSendModel2 string) {
	this.CrossdomainMessageSendModel = crossdomainMessageSendModel2
}
func (this *OapiDingtalkImpaasMessageCrossdomainSendRequest) GetCrossdomainMessageSendModel() string {
	return this.CrossdomainMessageSendModel
}
func (this *OapiDingtalkImpaasMessageCrossdomainSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingtalk.impaas.message.crossdomain.send"
}
func (this *OapiDingtalkImpaasMessageCrossdomainSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingtalkImpaasMessageCrossdomainSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingtalkImpaasMessageCrossdomainSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingtalkImpaasMessageCrossdomainSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingtalkImpaasMessageCrossdomainSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["crossdomain_message_send_model"] = this.CrossdomainMessageSendModel
	return txtParams
}
func (this *OapiDingtalkImpaasMessageCrossdomainSendRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDingtalkImpaasMessageCrossdomainSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingtalkImpaasMessageCrossdomainSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopSendMessageModel struct {
	Cid              string `json:"cid,omitempty"`
	Content          string `json:"content,omitempty"`
	ContentType      int64  `json:"content_type,omitempty"`
	ConversationType int64  `json:"conversation_type,omitempty"`
	Uuid             string `json:"uuid,omitempty"`
}
type MessageReceiverScopeModel struct {
	ActualReceivers []string `json:"actual_receivers,omitempty"`
}
type TopCrossDomainMessageSendModel struct {
	MessageId                 string                    `json:"message_id,omitempty"`
	MessageReceiverScopeModel MessageReceiverScopeModel `json:"message_receiver_scope_model,omitempty"`
	SendMessageModel          TopSendMessageModel       `json:"send_message_model,omitempty"`
	Sender                    string                    `json:"sender,omitempty"`
}
type OapiDingtalkImpaasMessageCrossdomainSendResponse struct {
	taobao.TaobaoResponse
}
