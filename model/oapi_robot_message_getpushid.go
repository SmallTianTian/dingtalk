package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRobotMessageGetpushidRequest() *OapiRobotMessageGetpushidRequest {
	return &OapiRobotMessageGetpushidRequest{}
}

type OapiRobotMessageGetpushidRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRobotMessageGetpushidResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRobotMessageGetpushidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRobotMessageGetpushidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRobotMessageGetpushidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.robot.message.getpushid"
}
func (this *OapiRobotMessageGetpushidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRobotMessageGetpushidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRobotMessageGetpushidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRobotMessageGetpushidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRobotMessageGetpushidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiRobotMessageGetpushidRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRobotMessageGetpushidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRobotMessageGetpushidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRobotMessageGetpushidResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
