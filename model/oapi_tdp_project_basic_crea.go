package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpProjectBasicCreateRequest() *OapiTdpProjectBasicCreateRequest {
	return &OapiTdpProjectBasicCreateRequest{}
}

type OapiTdpProjectBasicCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpProjectBasicCreateResponse
	MicroappAgentId int64
	OperatorUserid  string
	Project         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiTdpProjectBasicCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpProjectBasicCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpProjectBasicCreateRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpProjectBasicCreateRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpProjectBasicCreateRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpProjectBasicCreateRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpProjectBasicCreateRequest) SetProject(project2 string) {
	this.Project = project2
}
func (this *OapiTdpProjectBasicCreateRequest) GetProject() string {
	return this.Project
}
func (this *OapiTdpProjectBasicCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.project.basic.create"
}
func (this *OapiTdpProjectBasicCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpProjectBasicCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpProjectBasicCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpProjectBasicCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpProjectBasicCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["project"] = this.Project
	return txtParams
}
func (this *OapiTdpProjectBasicCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpProjectBasicCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpProjectBasicCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ProjectCreate struct {
	CreatorUserid  string    `json:"creator_userid,omitempty"`
	Description    string    `json:"description,omitempty"`
	GmtCreate      time.Time `json:"gmt_create,omitempty"`
	GmtModified    time.Time `json:"gmt_modified,omitempty"`
	Icon           string    `json:"icon,omitempty"`
	Identifier     string    `json:"identifier,omitempty"`
	IsArchived     bool      `json:"is_archived,omitempty"`
	IsRecycled     bool      `json:"is_recycled,omitempty"`
	ModifierUserid string    `json:"modifier_userid,omitempty"`
	Name           string    `json:"name,omitempty"`
	ParentId       string    `json:"parent_id,omitempty"`
	Source         string    `json:"source,omitempty"`
	SourceId       string    `json:"source_id,omitempty"`
	Visibility     string    `json:"visibility,omitempty"`
}
type OapiTdpProjectBasicCreateResponse struct {
	taobao.TaobaoResponse
	Result Project `json:"result,omitempty"`
}
type Project struct {
	BelongCorpId     string    `json:"belong_corp_id,omitempty"`
	BizTag           string    `json:"biz_tag,omitempty"`
	CreatorUserid    string    `json:"creator_userid,omitempty"`
	Description      string    `json:"description,omitempty"`
	GmtCreate        time.Time `json:"gmt_create,omitempty"`
	GmtModified      time.Time `json:"gmt_modified,omitempty"`
	Icon             string    `json:"icon,omitempty"`
	IsArchived       bool      `json:"is_archived,omitempty"`
	IsRecycled       bool      `json:"is_recycled,omitempty"`
	ModifierUserid   string    `json:"modifier_userid,omitempty"`
	Name             string    `json:"name,omitempty"`
	ParentId         string    `json:"parent_id,omitempty"`
	ProjectId        string    `json:"project_id,omitempty"`
	Source           string    `json:"source,omitempty"`
	SourceId         string    `json:"source_id,omitempty"`
	VirtualDingOrgid string    `json:"virtual_ding_orgid,omitempty"`
	Visibility       string    `json:"visibility,omitempty"`
}
