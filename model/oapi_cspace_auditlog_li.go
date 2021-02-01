package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCspaceAuditlogListRequest() *OapiCspaceAuditlogListRequest {
	return &OapiCspaceAuditlogListRequest{}
}

type OapiCspaceAuditlogListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp              OapiCspaceAuditlogListResponse
	EndDate           int64
	LoadMoreBizId     int64
	LoadMoreGmtCreate int64
	PageSize          int64
	StartDate         int64
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiCspaceAuditlogListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCspaceAuditlogListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCspaceAuditlogListRequest) SetEndDate(endDate2 int64) {
	this.EndDate = endDate2
}
func (this *OapiCspaceAuditlogListRequest) GetEndDate() int64 {
	return this.EndDate
}
func (this *OapiCspaceAuditlogListRequest) SetLoadMoreBizId(loadMoreBizId2 int64) {
	this.LoadMoreBizId = loadMoreBizId2
}
func (this *OapiCspaceAuditlogListRequest) GetLoadMoreBizId() int64 {
	return this.LoadMoreBizId
}
func (this *OapiCspaceAuditlogListRequest) SetLoadMoreGmtCreate(loadMoreGmtCreate2 int64) {
	this.LoadMoreGmtCreate = loadMoreGmtCreate2
}
func (this *OapiCspaceAuditlogListRequest) GetLoadMoreGmtCreate() int64 {
	return this.LoadMoreGmtCreate
}
func (this *OapiCspaceAuditlogListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiCspaceAuditlogListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiCspaceAuditlogListRequest) SetStartDate(startDate2 int64) {
	this.StartDate = startDate2
}
func (this *OapiCspaceAuditlogListRequest) GetStartDate() int64 {
	return this.StartDate
}
func (this *OapiCspaceAuditlogListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.cspace.auditlog.list"
}
func (this *OapiCspaceAuditlogListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCspaceAuditlogListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCspaceAuditlogListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCspaceAuditlogListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCspaceAuditlogListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["end_date"] = this.EndDate
	txtParams["load_more_biz_id"] = this.LoadMoreBizId
	txtParams["load_more_gmt_create"] = this.LoadMoreGmtCreate
	txtParams["page_size"] = this.PageSize
	txtParams["start_date"] = this.StartDate
	return txtParams
}
func (this *OapiCspaceAuditlogListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.EndDate, "endDate"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCspaceAuditlogListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCspaceAuditlogListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCspaceAuditlogListResponse struct {
	taobao.TaobaoResponse
	Result  ListAuditLogResult `json:"result,omitempty"`
	Success bool               `json:"success,omitempty"`
}
type AuditLogVO struct {
	Action            int64  `json:"action,omitempty"`
	ActionView        string `json:"action_view,omitempty"`
	BizId             string `json:"biz_id,omitempty"`
	GmtCreate         int64  `json:"gmt_create,omitempty"`
	GmtModified       int64  `json:"gmt_modified,omitempty"`
	IpAddress         string `json:"ip_address,omitempty"`
	OperateModule     int64  `json:"operate_module,omitempty"`
	OperateModuleView string `json:"operate_module_view,omitempty"`
	OperatorName      string `json:"operator_name,omitempty"`
	OrgName           string `json:"org_name,omitempty"`
	Platform          int64  `json:"platform,omitempty"`
	PlatformView      string `json:"platform_view,omitempty"`
	ReceiverName      string `json:"receiver_name,omitempty"`
	ReceiverType      int64  `json:"receiver_type,omitempty"`
	ReceiverTypeView  string `json:"receiver_type_view,omitempty"`
	Resource          string `json:"resource,omitempty"`
	ResourceExtension string `json:"resource_extension,omitempty"`
	ResourceSize      int64  `json:"resource_size,omitempty"`
	Status            int64  `json:"status,omitempty"`
	Userid            string `json:"userid,omitempty"`
}
type ListAuditLogResult struct {
	List []AuditLogVO `json:"list,omitempty"`
}
