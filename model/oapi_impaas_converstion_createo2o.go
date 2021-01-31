package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasConverstionCreateo2oRequest() *OapiImpaasConverstionCreateo2oRequest {
	return &OapiImpaasConverstionCreateo2oRequest{}
}

type OapiImpaasConverstionCreateo2oRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasConverstionCreateo2oResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasConverstionCreateo2oRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasConverstionCreateo2oRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasConverstionCreateo2oRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasConverstionCreateo2oRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasConverstionCreateo2oRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.converstion.createo2o"
}
func (this *OapiImpaasConverstionCreateo2oRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasConverstionCreateo2oRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasConverstionCreateo2oRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasConverstionCreateo2oRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasConverstionCreateo2oRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasConverstionCreateo2oRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasConverstionCreateo2oRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasConverstionCreateo2oRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ConversationAccountInfo struct {
	Id   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}
type CreateConversationRequest struct {
	Channel            string                  `json:"channel,omitempty"`
	Extension          string                  `json:"extension,omitempty"`
	Receiver           ConversationAccountInfo `json:"receiver,omitempty"`
	ReceiverEntranceId int64                   `json:"receiver_entrance_id,omitempty"`
	Sender             ConversationAccountInfo `json:"sender,omitempty"`
	SenderEntranceId   int64                   `json:"sender_entrance_id,omitempty"`
	Uuid               string                  `json:"uuid,omitempty"`
}
type OapiImpaasConverstionCreateo2oResponse struct {
	taobao.TaobaoResponse
	Chatid  string `json:"chatid,omitempty"`
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
