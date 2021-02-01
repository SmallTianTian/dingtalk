package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMessageCorpconversationRecallRequest() *OapiMessageCorpconversationRecallRequest {
	return &OapiMessageCorpconversationRecallRequest{}
}

type OapiMessageCorpconversationRecallRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMessageCorpconversationRecallResponse
	AgentId         int64
	MsgTaskId       int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMessageCorpconversationRecallRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMessageCorpconversationRecallRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMessageCorpconversationRecallRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMessageCorpconversationRecallRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMessageCorpconversationRecallRequest) SetMsgTaskId(msgTaskId2 int64) {
	this.MsgTaskId = msgTaskId2
}
func (this *OapiMessageCorpconversationRecallRequest) GetMsgTaskId() int64 {
	return this.MsgTaskId
}
func (this *OapiMessageCorpconversationRecallRequest) GetApiMethodName() string {
	return "dingtalk.oapi.message.corpconversation.recall"
}
func (this *OapiMessageCorpconversationRecallRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMessageCorpconversationRecallRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMessageCorpconversationRecallRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMessageCorpconversationRecallRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMessageCorpconversationRecallRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["msg_task_id"] = this.MsgTaskId
	return txtParams
}
func (this *OapiMessageCorpconversationRecallRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMessageCorpconversationRecallRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMessageCorpconversationRecallRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMessageCorpconversationRecallResponse struct {
	taobao.TaobaoResponse
}
