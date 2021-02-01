package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceTaskGetbysourceidRequest() *OapiWorkspaceTaskGetbysourceidRequest {
	return &OapiWorkspaceTaskGetbysourceidRequest{}
}

type OapiWorkspaceTaskGetbysourceidRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceTaskGetbysourceidResponse
	AgentId         int64
	Source          string
	SourceId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceTaskGetbysourceidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) SetSource(source2 string) {
	this.Source = source2
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) GetSource() string {
	return this.Source
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) SetSourceId(sourceId2 string) {
	this.SourceId = sourceId2
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) GetSourceId() string {
	return this.SourceId
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.task.getbysourceid"
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["source"] = this.Source
	txtParams["source_id"] = this.SourceId
	return txtParams
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.SourceId, "sourceId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceTaskGetbysourceidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceTaskGetbysourceidResponse struct {
	taobao.TaobaoResponse
	Result Task `json:"result,omitempty"`
}
