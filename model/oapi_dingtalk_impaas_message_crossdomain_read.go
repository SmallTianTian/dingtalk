package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDingtalkImpaasMessageCrossdomainReadRequest() *OapiDingtalkImpaasMessageCrossdomainReadRequest {
	return &OapiDingtalkImpaasMessageCrossdomainReadRequest{}
}

type OapiDingtalkImpaasMessageCrossdomainReadRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp             OapiDingtalkImpaasMessageCrossdomainReadResponse
	MessageReadModel string
	TopHttpMethod    string
	TopResponseType  string
}

func (this *OapiDingtalkImpaasMessageCrossdomainReadRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingtalkImpaasMessageCrossdomainReadRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingtalkImpaasMessageCrossdomainReadRequest) SetMessageReadModel(messageReadModel2 string) {
	this.MessageReadModel = messageReadModel2
}
func (this *OapiDingtalkImpaasMessageCrossdomainReadRequest) GetMessageReadModel() string {
	return this.MessageReadModel
}
func (this *OapiDingtalkImpaasMessageCrossdomainReadRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingtalk.impaas.message.crossdomain.read"
}
func (this *OapiDingtalkImpaasMessageCrossdomainReadRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingtalkImpaasMessageCrossdomainReadRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingtalkImpaasMessageCrossdomainReadRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingtalkImpaasMessageCrossdomainReadRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingtalkImpaasMessageCrossdomainReadRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["message_read_model"] = this.MessageReadModel
	return txtParams
}
func (this *OapiDingtalkImpaasMessageCrossdomainReadRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDingtalkImpaasMessageCrossdomainReadRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingtalkImpaasMessageCrossdomainReadRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopCrossDomainMessageReadModel struct {
	MessageIds  []string `json:"message_ids,omitempty"`
	OperatorUid string   `json:"operator_uid,omitempty"`
}
type OapiDingtalkImpaasMessageCrossdomainReadResponse struct {
	taobao.TaobaoResponse
}
