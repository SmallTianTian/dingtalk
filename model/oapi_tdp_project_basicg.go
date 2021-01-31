package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"time"
)

func NewOapiTdpProjectBasicGetRequest() *OapiTdpProjectBasicGetRequest {
	return &OapiTdpProjectBasicGetRequest{}
}

type OapiTdpProjectBasicGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpProjectBasicGetResponse
	MicroappAgentId int64
	ProjectId       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiTdpProjectBasicGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpProjectBasicGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpProjectBasicGetRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpProjectBasicGetRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpProjectBasicGetRequest) SetProjectId(projectId2 string) {
	this.ProjectId = projectId2
}
func (this *OapiTdpProjectBasicGetRequest) GetProjectId() string {
	return this.ProjectId
}
func (this *OapiTdpProjectBasicGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.project.basic.get"
}
func (this *OapiTdpProjectBasicGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpProjectBasicGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpProjectBasicGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpProjectBasicGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpProjectBasicGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["project_id"] = this.ProjectId
	return txtParams
}
func (this *OapiTdpProjectBasicGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiTdpProjectBasicGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpProjectBasicGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiTdpProjectBasicGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  Task   `json:"result,omitempty"`
}
type Task struct {
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
	SourceId         string    `json:"source_id,omitempty"`
	VirtualDingOrgid string    `json:"virtual_ding_orgid,omitempty"`
	Visibility       string    `json:"visibility,omitempty"`
}
