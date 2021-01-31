package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiImpaasConversaionChangegroupownerRequest() *OapiImpaasConversaionChangegroupownerRequest {
	return &OapiImpaasConversaionChangegroupownerRequest{}
}

type OapiImpaasConversaionChangegroupownerRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiImpaasConversaionChangegroupownerResponse
	Channel         string
	Chatid          string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiImpaasConversaionChangegroupownerRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiImpaasConversaionChangegroupownerRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiImpaasConversaionChangegroupownerRequest) SetChannel(channel2 string) {
	this.Channel = channel2
}
func (this *OapiImpaasConversaionChangegroupownerRequest) GetChannel() string {
	return this.Channel
}
func (this *OapiImpaasConversaionChangegroupownerRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiImpaasConversaionChangegroupownerRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiImpaasConversaionChangegroupownerRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiImpaasConversaionChangegroupownerRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiImpaasConversaionChangegroupownerRequest) GetApiMethodName() string {
	return "dingtalk.oapi.impaas.conversaion.changegroupowner"
}
func (this *OapiImpaasConversaionChangegroupownerRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiImpaasConversaionChangegroupownerRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiImpaasConversaionChangegroupownerRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiImpaasConversaionChangegroupownerRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiImpaasConversaionChangegroupownerRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["channel"] = this.Channel
	txtParams["chatid"] = this.Chatid
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiImpaasConversaionChangegroupownerRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiImpaasConversaionChangegroupownerRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiImpaasConversaionChangegroupownerRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiImpaasConversaionChangegroupownerResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
}
