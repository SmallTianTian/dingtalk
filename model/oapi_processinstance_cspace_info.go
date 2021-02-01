package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessinstanceCspaceInfoRequest() *OapiProcessinstanceCspaceInfoRequest {
	return &OapiProcessinstanceCspaceInfoRequest{}
}

type OapiProcessinstanceCspaceInfoRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessinstanceCspaceInfoResponse
	AgentId         int64
	TopHttpMethod   string
	TopResponseType string
	UserId          string
}

func (this *OapiProcessinstanceCspaceInfoRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessinstanceCspaceInfoRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessinstanceCspaceInfoRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiProcessinstanceCspaceInfoRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiProcessinstanceCspaceInfoRequest) SetUserId(userId2 string) {
	this.UserId = userId2
}
func (this *OapiProcessinstanceCspaceInfoRequest) GetUserId() string {
	return this.UserId
}
func (this *OapiProcessinstanceCspaceInfoRequest) GetApiMethodName() string {
	return "dingtalk.oapi.processinstance.cspace.info"
}
func (this *OapiProcessinstanceCspaceInfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessinstanceCspaceInfoRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessinstanceCspaceInfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessinstanceCspaceInfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessinstanceCspaceInfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["user_id"] = this.UserId
	return txtParams
}
func (this *OapiProcessinstanceCspaceInfoRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.UserId, "userId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessinstanceCspaceInfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessinstanceCspaceInfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessinstanceCspaceInfoResponse struct {
	taobao.TaobaoResponse
	Result  AppSpaceResponse `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
type AppSpaceResponse struct {
	SpaceId int64 `json:"space_id,omitempty"`
}
