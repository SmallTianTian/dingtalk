package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRobotMessageOrggrouptaskQueryRequest() *OapiRobotMessageOrggrouptaskQueryRequest {
	return &OapiRobotMessageOrggrouptaskQueryRequest{}
}

type OapiRobotMessageOrggrouptaskQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiRobotMessageOrggrouptaskQueryResponse
	ChatbotId          string
	OpenConversationId string
	ProcessQueryKey    string
	Token              string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiRobotMessageOrggrouptaskQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) SetChatbotId(chatbotId2 string) {
	this.ChatbotId = chatbotId2
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) GetChatbotId() string {
	return this.ChatbotId
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) SetProcessQueryKey(processQueryKey2 string) {
	this.ProcessQueryKey = processQueryKey2
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) GetProcessQueryKey() string {
	return this.ProcessQueryKey
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) SetToken(token2 string) {
	this.Token = token2
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) GetToken() string {
	return this.Token
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.robot.message.orggrouptask.query"
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["chatbot_id"] = this.ChatbotId
	txtParams["open_conversation_id"] = this.OpenConversationId
	txtParams["process_query_key"] = this.ProcessQueryKey
	txtParams["token"] = this.Token
	return txtParams
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProcessQueryKey, "processQueryKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRobotMessageOrggrouptaskQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRobotMessageOrggrouptaskQueryResponse struct {
	taobao.TaobaoResponse
	Result  GroupMessageTaskQueryTopResponse `json:"result,omitempty"`
	Success bool                             `json:"success,omitempty"`
}
type GroupMessageTaskQueryTopResponse struct {
	ReadStaffIds []string `json:"read_staff_ids,omitempty"`
	SendStatus   string   `json:"send_status,omitempty"`
}
