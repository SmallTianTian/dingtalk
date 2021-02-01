package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpProjectBasicUpdateRequest() *OapiTdpProjectBasicUpdateRequest {
	return &OapiTdpProjectBasicUpdateRequest{}
}

type OapiTdpProjectBasicUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpProjectBasicUpdateResponse
	MicroappAgentId int64
	OperatorUserid  string
	Project         string
	ProjectId       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiTdpProjectBasicUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpProjectBasicUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpProjectBasicUpdateRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpProjectBasicUpdateRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpProjectBasicUpdateRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpProjectBasicUpdateRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpProjectBasicUpdateRequest) SetProject(project2 string) {
	this.Project = project2
}
func (this *OapiTdpProjectBasicUpdateRequest) GetProject() string {
	return this.Project
}
func (this *OapiTdpProjectBasicUpdateRequest) SetProjectId(projectId2 string) {
	this.ProjectId = projectId2
}
func (this *OapiTdpProjectBasicUpdateRequest) GetProjectId() string {
	return this.ProjectId
}
func (this *OapiTdpProjectBasicUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.project.basic.update"
}
func (this *OapiTdpProjectBasicUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpProjectBasicUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpProjectBasicUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpProjectBasicUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpProjectBasicUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["project"] = this.Project
	txtParams["project_id"] = this.ProjectId
	return txtParams
}
func (this *OapiTdpProjectBasicUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProjectId, "projectId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpProjectBasicUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpProjectBasicUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ProjectUpdate struct {
	Description    string    `json:"description,omitempty"`
	GmtModified    time.Time `json:"gmt_modified,omitempty"`
	Icon           string    `json:"icon,omitempty"`
	Identifier     string    `json:"identifier,omitempty"`
	IsArchived     bool      `json:"is_archived,omitempty"`
	IsRecycled     bool      `json:"is_recycled,omitempty"`
	ModifierUserid string    `json:"modifier_userid,omitempty"`
	Name           string    `json:"name,omitempty"`
	ParentId       string    `json:"parent_id,omitempty"`
	Visibility     string    `json:"visibility,omitempty"`
}
type OapiTdpProjectBasicUpdateResponse struct {
	taobao.TaobaoResponse
}
