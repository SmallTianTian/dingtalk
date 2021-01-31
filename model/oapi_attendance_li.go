package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAttendanceListRequest() *OapiAttendanceListRequest {
	return &OapiAttendanceListRequest{}
}

type OapiAttendanceListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceListResponse
	IsI18n          bool
	Limit           int64
	Offset          int64
	TopHttpMethod   string
	TopResponseType string
	UserIdList      []string
	WorkDateFrom    string
	WorkDateTo      string
}

func (this *OapiAttendanceListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceListRequest) SetIsI18n(isI18n2 bool) {
	this.IsI18n = isI18n2
}
func (this *OapiAttendanceListRequest) GetIsI18n() bool {
	return this.IsI18n
}
func (this *OapiAttendanceListRequest) SetLimit(limit2 int64) {
	this.Limit = limit2
}
func (this *OapiAttendanceListRequest) GetLimit() int64 {
	return this.Limit
}
func (this *OapiAttendanceListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiAttendanceListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiAttendanceListRequest) SetUserIdList(userIdList2 []string) {
	this.UserIdList = userIdList2
}
func (this *OapiAttendanceListRequest) GetUserIdList() []string {
	return this.UserIdList
}
func (this *OapiAttendanceListRequest) SetWorkDateFrom(workDateFrom2 string) {
	this.WorkDateFrom = workDateFrom2
}
func (this *OapiAttendanceListRequest) GetWorkDateFrom() string {
	return this.WorkDateFrom
}
func (this *OapiAttendanceListRequest) SetWorkDateTo(workDateTo2 string) {
	this.WorkDateTo = workDateTo2
}
func (this *OapiAttendanceListRequest) GetWorkDateTo() string {
	return this.WorkDateTo
}
func (this *OapiAttendanceListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.list"
}
func (this *OapiAttendanceListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["isI18n"] = this.IsI18n
	txtParams["limit"] = this.Limit
	txtParams["offset"] = this.Offset
	txtParams["userIdList"] = this.UserIdList
	txtParams["workDateFrom"] = this.WorkDateFrom
	txtParams["workDateTo"] = this.WorkDateTo
	return txtParams
}
func (this *OapiAttendanceListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAttendanceListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceListResponse struct {
	taobao.TaobaoResponse
	Errcode      int64          `json:"errcode,omitempty"`
	Errmsg       string         `json:"errmsg,omitempty"`
	HasMore      bool           `json:"hasMore,omitempty"`
	Recordresult []Recordresult `json:"recordresult,omitempty"`
}
