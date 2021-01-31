package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProjectPointHistoryPageRequest() *OapiProjectPointHistoryPageRequest {
	return &OapiProjectPointHistoryPageRequest{}
}

type OapiProjectPointHistoryPageRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProjectPointHistoryPageResponse
	Cursor          int64
	PageSize        int64
	TenantId        int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiProjectPointHistoryPageRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProjectPointHistoryPageRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProjectPointHistoryPageRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiProjectPointHistoryPageRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiProjectPointHistoryPageRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiProjectPointHistoryPageRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiProjectPointHistoryPageRequest) SetTenantId(tenantId2 int64) {
	this.TenantId = tenantId2
}
func (this *OapiProjectPointHistoryPageRequest) GetTenantId() int64 {
	return this.TenantId
}
func (this *OapiProjectPointHistoryPageRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiProjectPointHistoryPageRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiProjectPointHistoryPageRequest) GetApiMethodName() string {
	return "dingtalk.oapi.project.point.history.page"
}
func (this *OapiProjectPointHistoryPageRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProjectPointHistoryPageRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProjectPointHistoryPageRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProjectPointHistoryPageRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProjectPointHistoryPageRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["page_size"] = this.PageSize
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiProjectPointHistoryPageRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProjectPointHistoryPageRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProjectPointHistoryPageRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProjectPointHistoryPageResponse struct {
	taobao.TaobaoResponse
	Result  PageResult `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
type PointHistoryDTO struct {
	CorpId   string `json:"corp_id,omitempty"`
	CreateAt int64  `json:"create_at,omitempty"`
	RuleCode string `json:"rule_code,omitempty"`
	RuleName string `json:"rule_name,omitempty"`
	Score    int64  `json:"score,omitempty"`
	Userid   string `json:"userid,omitempty"`
	Uuid     string `json:"uuid,omitempty"`
}
