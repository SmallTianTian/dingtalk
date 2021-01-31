package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCustomerserviceMessageSendRequest() *OapiCustomerserviceMessageSendRequest {
	return &OapiCustomerserviceMessageSendRequest{}
}

type OapiCustomerserviceMessageSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCustomerserviceMessageSendResponse
	Message         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCustomerserviceMessageSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCustomerserviceMessageSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCustomerserviceMessageSendRequest) SetMessage(message2 string) {
	this.Message = message2
}
func (this *OapiCustomerserviceMessageSendRequest) GetMessage() string {
	return this.Message
}
func (this *OapiCustomerserviceMessageSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.customerservice.message.send"
}
func (this *OapiCustomerserviceMessageSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCustomerserviceMessageSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCustomerserviceMessageSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCustomerserviceMessageSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCustomerserviceMessageSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["message"] = this.Message
	return txtParams
}
func (this *OapiCustomerserviceMessageSendRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCustomerserviceMessageSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCustomerserviceMessageSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type MessageDTO struct {
	BuId            string `json:"bu_id,omitempty"`
	CmsId           string `json:"cms_id,omitempty"`
	Content         string `json:"content,omitempty"`
	ContentType     int64  `json:"content_type,omitempty"`
	MessageCreateAt int64  `json:"message_create_at,omitempty"`
	MessageId       string `json:"message_id,omitempty"`
	SenderId        string `json:"sender_id,omitempty"`
	SenderType      string `json:"sender_type,omitempty"`
	SessionSource   string `json:"session_source,omitempty"`
	Sid             string `json:"sid,omitempty"`
}
type OapiCustomerserviceMessageSendResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
