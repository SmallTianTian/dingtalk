package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpProjectMemberGetbyprojectRequest() *OapiTdpProjectMemberGetbyprojectRequest {
	return &OapiTdpProjectMemberGetbyprojectRequest{}
}

type OapiTdpProjectMemberGetbyprojectRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpProjectMemberGetbyprojectResponse
	MicroappAgentId int64
	PageRequest     string
	ProjectId       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiTdpProjectMemberGetbyprojectRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) SetPageRequest(pageRequest2 string) {
	this.PageRequest = pageRequest2
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) GetPageRequest() string {
	return this.PageRequest
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) SetProjectId(projectId2 string) {
	this.ProjectId = projectId2
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) GetProjectId() string {
	return this.ProjectId
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.project.member.getbyproject"
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["page_request"] = this.PageRequest
	txtParams["project_id"] = this.ProjectId
	return txtParams
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProjectId, "projectId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpProjectMemberGetbyprojectRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PageRequest struct {
	OrderBy        string `json:"order_by,omitempty"`
	OrderDirection string `json:"order_direction,omitempty"`
	Page           int64  `json:"page,omitempty"`
	PageSize       int64  `json:"page_size,omitempty"`
}
type OapiTdpProjectMemberGetbyprojectResponse struct {
	taobao.TaobaoResponse
	Result PageResponse `json:"result,omitempty"`
}
type Member struct {
	BelongCorpId   string    `json:"belong_corp_id,omitempty"`
	BizTag         string    `json:"biz_tag,omitempty"`
	CreatorUserid  string    `json:"creator_userid,omitempty"`
	GmtCreate      time.Time `json:"gmt_create,omitempty"`
	GmtModified    time.Time `json:"gmt_modified,omitempty"`
	MemberId       string    `json:"member_id,omitempty"`
	ModifierUserid string    `json:"modifier_userid,omitempty"`
	TargetId       string    `json:"target_id,omitempty"`
	TargetType     string    `json:"target_type,omitempty"`
	Userid         string    `json:"userid,omitempty"`
}
type PageResponse struct {
	Data       []Member `json:"data,omitempty"`
	Page       int64    `json:"page,omitempty"`
	PageSize   int64    `json:"page_size,omitempty"`
	TotalCount int64    `json:"total_count,omitempty"`
}
