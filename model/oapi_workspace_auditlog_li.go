package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceAuditlogListRequest() *OapiWorkspaceAuditlogListRequest {
	return &OapiWorkspaceAuditlogListRequest{}
}

type OapiWorkspaceAuditlogListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp              OapiWorkspaceAuditlogListResponse
	EndDate           int64
	LoadMoreBizId     int64
	LoadMoreGmtCreate int64
	PageSize          int64
	StartDate         int64
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiWorkspaceAuditlogListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceAuditlogListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceAuditlogListRequest) SetEndDate(endDate2 int64) {
	this.EndDate = endDate2
}
func (this *OapiWorkspaceAuditlogListRequest) GetEndDate() int64 {
	return this.EndDate
}
func (this *OapiWorkspaceAuditlogListRequest) SetLoadMoreBizId(loadMoreBizId2 int64) {
	this.LoadMoreBizId = loadMoreBizId2
}
func (this *OapiWorkspaceAuditlogListRequest) GetLoadMoreBizId() int64 {
	return this.LoadMoreBizId
}
func (this *OapiWorkspaceAuditlogListRequest) SetLoadMoreGmtCreate(loadMoreGmtCreate2 int64) {
	this.LoadMoreGmtCreate = loadMoreGmtCreate2
}
func (this *OapiWorkspaceAuditlogListRequest) GetLoadMoreGmtCreate() int64 {
	return this.LoadMoreGmtCreate
}
func (this *OapiWorkspaceAuditlogListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiWorkspaceAuditlogListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiWorkspaceAuditlogListRequest) SetStartDate(startDate2 int64) {
	this.StartDate = startDate2
}
func (this *OapiWorkspaceAuditlogListRequest) GetStartDate() int64 {
	return this.StartDate
}
func (this *OapiWorkspaceAuditlogListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.auditlog.list"
}
func (this *OapiWorkspaceAuditlogListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceAuditlogListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceAuditlogListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceAuditlogListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceAuditlogListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["end_date"] = this.EndDate
	txtParams["load_more_bizId"] = this.LoadMoreBizId
	txtParams["load_more_gmt_create"] = this.LoadMoreGmtCreate
	txtParams["page_size"] = this.PageSize
	txtParams["start_date"] = this.StartDate
	return txtParams
}
func (this *OapiWorkspaceAuditlogListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.EndDate, "endDate"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceAuditlogListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceAuditlogListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceAuditlogListResponse struct {
	taobao.TaobaoResponse
	Result OpenAuditLogDto `json:"result,omitempty"`
}
type EventAuditLogDto struct {
	Action            string `json:"action,omitempty"`
	BizId             string `json:"biz_id,omitempty"`
	Browser           string `json:"browser,omitempty"`
	DingTalkId        string `json:"ding_talk_id,omitempty"`
	EmpId             string `json:"emp_id,omitempty"`
	GmtCreate         string `json:"gmt_create,omitempty"`
	IpAddress         string `json:"ip_address,omitempty"`
	OperatorName      string `json:"operator_name,omitempty"`
	OrgName           string `json:"org_name,omitempty"`
	Platform          string `json:"platform,omitempty"`
	ProjectName       string `json:"project_name,omitempty"`
	ReceiverName      string `json:"receiver_name,omitempty"`
	Resource          string `json:"resource,omitempty"`
	ResourceExtension string `json:"resource_extension,omitempty"`
	ResourceSize      string `json:"resource_size,omitempty"`
	TaskName          string `json:"task_name,omitempty"`
}
type OpenAuditLogDto struct {
	LogList []EventAuditLogDto `json:"log_list,omitempty"`
}
