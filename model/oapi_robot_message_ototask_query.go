package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRobotMessageOtotaskQueryRequest() *OapiRobotMessageOtotaskQueryRequest {
	return &OapiRobotMessageOtotaskQueryRequest{}
}

type OapiRobotMessageOtotaskQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRobotMessageOtotaskQueryResponse
	ChatbotId       string
	ProcessQueryKey string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRobotMessageOtotaskQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRobotMessageOtotaskQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRobotMessageOtotaskQueryRequest) SetChatbotId(chatbotId2 string) {
	this.ChatbotId = chatbotId2
}
func (this *OapiRobotMessageOtotaskQueryRequest) GetChatbotId() string {
	return this.ChatbotId
}
func (this *OapiRobotMessageOtotaskQueryRequest) SetProcessQueryKey(processQueryKey2 string) {
	this.ProcessQueryKey = processQueryKey2
}
func (this *OapiRobotMessageOtotaskQueryRequest) GetProcessQueryKey() string {
	return this.ProcessQueryKey
}
func (this *OapiRobotMessageOtotaskQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.robot.message.ototask.query"
}
func (this *OapiRobotMessageOtotaskQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRobotMessageOtotaskQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRobotMessageOtotaskQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRobotMessageOtotaskQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRobotMessageOtotaskQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatbot_id"] = this.ChatbotId
	txtParams["process_query_key"] = this.ProcessQueryKey
	return txtParams
}
func (this *OapiRobotMessageOtotaskQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ChatbotId, "chatbotId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRobotMessageOtotaskQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRobotMessageOtotaskQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRobotMessageOtotaskQueryResponse struct {
	taobao.TaobaoResponse
	Result  OtoMessageSendTopResponse `json:"result,omitempty"`
	Success bool                      `json:"success,omitempty"`
}
type OtoMessageSendTopResponse struct {
	ReadStatus    string `json:"read_status,omitempty"`
	ReadTimestamp int64  `json:"read_timestamp,omitempty"`
	SendStatus    string `json:"send_status,omitempty"`
}
