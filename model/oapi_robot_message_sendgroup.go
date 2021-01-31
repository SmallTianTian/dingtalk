package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRobotMessageSendgroupRequest() *OapiRobotMessageSendgroupRequest {
	return &OapiRobotMessageSendgroupRequest{}
}

type OapiRobotMessageSendgroupRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRobotMessageSendgroupResponse
	MsgKey          string
	MsgParam        string
	Token           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRobotMessageSendgroupRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRobotMessageSendgroupRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRobotMessageSendgroupRequest) SetMsgKey(msgKey2 string) {
	this.MsgKey = msgKey2
}
func (this *OapiRobotMessageSendgroupRequest) GetMsgKey() string {
	return this.MsgKey
}
func (this *OapiRobotMessageSendgroupRequest) SetMsgParam(msgParam2 string) {
	this.MsgParam = msgParam2
}
func (this *OapiRobotMessageSendgroupRequest) GetMsgParam() string {
	return this.MsgParam
}
func (this *OapiRobotMessageSendgroupRequest) SetToken(token2 string) {
	this.Token = token2
}
func (this *OapiRobotMessageSendgroupRequest) GetToken() string {
	return this.Token
}
func (this *OapiRobotMessageSendgroupRequest) GetApiMethodName() string {
	return "dingtalk.oapi.robot.message.sendgroup"
}
func (this *OapiRobotMessageSendgroupRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRobotMessageSendgroupRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRobotMessageSendgroupRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRobotMessageSendgroupRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRobotMessageSendgroupRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["msg_key"] = this.MsgKey
	txtParams["msg_param"] = this.MsgParam
	txtParams["token"] = this.Token
	return txtParams
}
func (this *OapiRobotMessageSendgroupRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.MsgKey, "msgKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRobotMessageSendgroupRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRobotMessageSendgroupRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRobotMessageSendgroupResponse struct {
	taobao.TaobaoResponse
	Errcode int64                       `json:"errcode,omitempty"`
	Errmsg  string                      `json:"errmsg,omitempty"`
	Result  GroupMessageSendTopResponse `json:"result,omitempty"`
	Success bool                        `json:"success,omitempty"`
}
