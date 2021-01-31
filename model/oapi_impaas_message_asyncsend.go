package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasMessageAsyncsendRequest() *OapiImpaasMessageAsyncsendRequest {
	return &OapiImpaasMessageAsyncsendRequest{}
}

type OapiImpaasMessageAsyncsendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasMessageAsyncsendResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiImpaasMessageAsyncsendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasMessageAsyncsendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasMessageAsyncsendRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiImpaasMessageAsyncsendRequest) GetRequest() string {
	return this.Request
}
func (this *OapiImpaasMessageAsyncsendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.message.asyncsend"
}
func (this *OapiImpaasMessageAsyncsendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasMessageAsyncsendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasMessageAsyncsendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasMessageAsyncsendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasMessageAsyncsendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiImpaasMessageAsyncsendRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasMessageAsyncsendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasMessageAsyncsendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type XpnContentModel struct {
	AlertContent string `json:"alert_content,omitempty"`
	Params       string `json:"params,omitempty"`
}
type AsyncSendMessageRequest struct {
	GroupId        string          `json:"group_id,omitempty"`
	MsgContent     string          `json:"msg_content,omitempty"`
	MsgExtension   string          `json:"msg_extension,omitempty"`
	MsgFeature     string          `json:"msg_feature,omitempty"`
	MsgType        string          `json:"msg_type,omitempty"`
	ReceiveridList []string        `json:"receiverid_list,omitempty"`
	Senderid       AccountInfo     `json:"senderid,omitempty"`
	XpnModel       XpnContentModel `json:"xpn_model,omitempty"`
}
type OapiImpaasMessageAsyncsendResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Taskid  int64  `json:"taskid,omitempty"`
}
