package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRobotMessageGrouptaskQueryRequest() *OapiRobotMessageGrouptaskQueryRequest {
	return &OapiRobotMessageGrouptaskQueryRequest{}
}

type OapiRobotMessageGrouptaskQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRobotMessageGrouptaskQueryResponse
	ProcessQueryKey string
	Token           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRobotMessageGrouptaskQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRobotMessageGrouptaskQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRobotMessageGrouptaskQueryRequest) SetProcessQueryKey(processQueryKey2 string) {
	this.ProcessQueryKey = processQueryKey2
}
func (this *OapiRobotMessageGrouptaskQueryRequest) GetProcessQueryKey() string {
	return this.ProcessQueryKey
}
func (this *OapiRobotMessageGrouptaskQueryRequest) SetToken(token2 string) {
	this.Token = token2
}
func (this *OapiRobotMessageGrouptaskQueryRequest) GetToken() string {
	return this.Token
}
func (this *OapiRobotMessageGrouptaskQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.robot.message.grouptask.query"
}
func (this *OapiRobotMessageGrouptaskQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRobotMessageGrouptaskQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRobotMessageGrouptaskQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRobotMessageGrouptaskQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRobotMessageGrouptaskQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["process_query_key"] = this.ProcessQueryKey
	txtParams["token"] = this.Token
	return txtParams
}
func (this *OapiRobotMessageGrouptaskQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProcessQueryKey, "processQueryKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRobotMessageGrouptaskQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRobotMessageGrouptaskQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRobotMessageGrouptaskQueryResponse struct {
	taobao.TaobaoResponse
	Result  GroupMessageSendTopResponse `json:"result,omitempty"`
	Success bool                        `json:"success,omitempty"`
}
type GroupMessageSendTopResponse struct {
	SendStatus string `json:"send_status,omitempty"`
}
