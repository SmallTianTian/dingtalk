package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiDingmiO2oSendRequest() *OapiDingmiO2oSendRequest {
	return &OapiDingmiO2oSendRequest{}
}

type OapiDingmiO2oSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingmiO2oSendResponse
	MsgKey          string
	MsgParam        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiDingmiO2oSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingmiO2oSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingmiO2oSendRequest) SetMsgKey(msgKey2 string) {
	this.MsgKey = msgKey2
}
func (this *OapiDingmiO2oSendRequest) GetMsgKey() string {
	return this.MsgKey
}
func (this *OapiDingmiO2oSendRequest) SetMsgParam(msgParam2 string) {
	this.MsgParam = msgParam2
}
func (this *OapiDingmiO2oSendRequest) GetMsgParam() string {
	return this.MsgParam
}
func (this *OapiDingmiO2oSendRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiDingmiO2oSendRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiDingmiO2oSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingmi.o2o.send"
}
func (this *OapiDingmiO2oSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingmiO2oSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingmiO2oSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingmiO2oSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingmiO2oSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["msg_key"] = this.MsgKey
	txtParams["msg_param"] = this.MsgParam
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiDingmiO2oSendRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiDingmiO2oSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingmiO2oSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingmiO2oSendResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
