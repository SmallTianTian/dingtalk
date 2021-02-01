package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasMessageGetmessagestatusRequest() *OapiImpaasMessageGetmessagestatusRequest {
	return &OapiImpaasMessageGetmessagestatusRequest{}
}

type OapiImpaasMessageGetmessagestatusRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasMessageGetmessagestatusResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasMessageGetmessagestatusRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasMessageGetmessagestatusRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasMessageGetmessagestatusRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasMessageGetmessagestatusRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasMessageGetmessagestatusRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.message.getmessagestatus"
}
func (this *OapiImpaasMessageGetmessagestatusRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasMessageGetmessagestatusRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasMessageGetmessagestatusRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasMessageGetmessagestatusRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasMessageGetmessagestatusRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasMessageGetmessagestatusRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasMessageGetmessagestatusRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasMessageGetmessagestatusRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GetMessageStatusRequest struct {
	Senderid AccountInfo `json:"senderid,omitempty"`
	Taskid   int64       `json:"taskid,omitempty"`
}
type OapiImpaasMessageGetmessagestatusResponse struct {
	taobao.TaobaoResponse
	Result GetMessageStatusResponse `json:"result,omitempty"`
}
type GetMessageStatusResponse struct {
	TaskCode   int64  `json:"task_code,omitempty"`
	TaskMsg    string `json:"task_msg,omitempty"`
	TaskStatus int64  `json:"task_status,omitempty"`
}
