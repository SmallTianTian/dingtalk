package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDingmiCommonO2oPushRequest() *OapiDingmiCommonO2oPushRequest {
	return &OapiDingmiCommonO2oPushRequest{}
}

type OapiDingmiCommonO2oPushRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingmiCommonO2oPushResponse
	ChatbotId       string
	MsgKey          string
	MsgParam        string
	StaffId         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDingmiCommonO2oPushRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingmiCommonO2oPushRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingmiCommonO2oPushRequest) SetChatbotId(chatbotId2 string) {
	this.ChatbotId = chatbotId2
}
func (this *OapiDingmiCommonO2oPushRequest) GetChatbotId() string {
	return this.ChatbotId
}
func (this *OapiDingmiCommonO2oPushRequest) SetMsgKey(msgKey2 string) {
	this.MsgKey = msgKey2
}
func (this *OapiDingmiCommonO2oPushRequest) GetMsgKey() string {
	return this.MsgKey
}
func (this *OapiDingmiCommonO2oPushRequest) SetMsgParam(msgParam2 string) {
	this.MsgParam = msgParam2
}
func (this *OapiDingmiCommonO2oPushRequest) GetMsgParam() string {
	return this.MsgParam
}
func (this *OapiDingmiCommonO2oPushRequest) SetStaffId(staffId2 string) {
	this.StaffId = staffId2
}
func (this *OapiDingmiCommonO2oPushRequest) GetStaffId() string {
	return this.StaffId
}
func (this *OapiDingmiCommonO2oPushRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingmi.common.o2o.push"
}
func (this *OapiDingmiCommonO2oPushRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingmiCommonO2oPushRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingmiCommonO2oPushRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingmiCommonO2oPushRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingmiCommonO2oPushRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatbot_id"] = this.ChatbotId
	txtParams["msg_key"] = this.MsgKey
	txtParams["msg_param"] = this.MsgParam
	txtParams["staff_id"] = this.StaffId
	return txtParams
}
func (this *OapiDingmiCommonO2oPushRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MsgKey, "msgKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingmiCommonO2oPushRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingmiCommonO2oPushRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingmiCommonO2oPushResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
}
