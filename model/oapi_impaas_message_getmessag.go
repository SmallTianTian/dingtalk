package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasMessageGetmessageRequest() *OapiImpaasMessageGetmessageRequest {
	return &OapiImpaasMessageGetmessageRequest{}
}

type OapiImpaasMessageGetmessageRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasMessageGetmessageResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasMessageGetmessageRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasMessageGetmessageRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasMessageGetmessageRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasMessageGetmessageRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasMessageGetmessageRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.message.getmessage"
}
func (this *OapiImpaasMessageGetmessageRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasMessageGetmessageRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasMessageGetmessageRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasMessageGetmessageRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasMessageGetmessageRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasMessageGetmessageRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasMessageGetmessageRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasMessageGetmessageRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GetMessageRequest struct {
	BegTime    int64       `json:"beg_time,omitempty"`
	GroupId    string      `json:"group_id,omitempty"`
	LimitNum   int64       `json:"limit_num,omitempty"`
	ReceiverId AccountInfo `json:"receiver_id,omitempty"`
	SenderId   AccountInfo `json:"sender_id,omitempty"`
}
type OapiImpaasMessageGetmessageResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
