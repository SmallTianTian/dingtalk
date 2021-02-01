package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiSnsSendMsgRequest() *OapiSnsSendMsgRequest {
	return &OapiSnsSendMsgRequest{}
}

type OapiSnsSendMsgRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiSnsSendMsgResponse
	Code            string
	Msg             string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiSnsSendMsgRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSnsSendMsgRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSnsSendMsgRequest) SetCode(code2 string) {
	this.Code = code2
}
func (this *OapiSnsSendMsgRequest) GetCode() string {
	return this.Code
}
func (this *OapiSnsSendMsgRequest) SetMsg(msg2 string) {
	this.Msg = msg2
}
func (this *OapiSnsSendMsgRequest) GetMsg() string {
	return this.Msg
}
func (this *OapiSnsSendMsgRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sns.send_msg"
}
func (this *OapiSnsSendMsgRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSnsSendMsgRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSnsSendMsgRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSnsSendMsgRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSnsSendMsgRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.ERROR_CODE] = this.Code
	txtParams["msg"] = this.Msg
	return txtParams
}
func (this *OapiSnsSendMsgRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiSnsSendMsgRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSnsSendMsgRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Eapp struct {
	Content string `json:"content,omitempty"`
	Img     string `json:"img,omitempty"`
	Link    string `json:"link,omitempty"`
	Title   string `json:"title,omitempty"`
}
type OapiSnsSendMsgResponse struct {
	taobao.TaobaoResponse
}
