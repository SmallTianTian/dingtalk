package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiLiveGroupliveViewrecordRequest() *OapiLiveGroupliveViewrecordRequest {
	return &OapiLiveGroupliveViewrecordRequest{}
}

type OapiLiveGroupliveViewrecordRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiLiveGroupliveViewrecordResponse
	DeptId          int64
	LiveUuid        string
	PageIndex       int64
	PageSize        int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiLiveGroupliveViewrecordRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiLiveGroupliveViewrecordRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiLiveGroupliveViewrecordRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiLiveGroupliveViewrecordRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiLiveGroupliveViewrecordRequest) SetLiveUuid(liveUuid2 string) {
	this.LiveUuid = liveUuid2
}
func (this *OapiLiveGroupliveViewrecordRequest) GetLiveUuid() string {
	return this.LiveUuid
}
func (this *OapiLiveGroupliveViewrecordRequest) SetPageIndex(pageIndex2 int64) {
	this.PageIndex = pageIndex2
}
func (this *OapiLiveGroupliveViewrecordRequest) GetPageIndex() int64 {
	return this.PageIndex
}
func (this *OapiLiveGroupliveViewrecordRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiLiveGroupliveViewrecordRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiLiveGroupliveViewrecordRequest) GetApiMethodName() string {
	return "dingtalk.oapi.live.grouplive.viewrecord"
}
func (this *OapiLiveGroupliveViewrecordRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiLiveGroupliveViewrecordRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiLiveGroupliveViewrecordRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiLiveGroupliveViewrecordRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiLiveGroupliveViewrecordRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dept_id"] = this.DeptId
	txtParams["live_uuid"] = this.LiveUuid
	txtParams["page_index"] = this.PageIndex
	txtParams["page_size"] = this.PageSize
	return txtParams
}
func (this *OapiLiveGroupliveViewrecordRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.DeptId, "deptId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiLiveGroupliveViewrecordRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiLiveGroupliveViewrecordRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiLiveGroupliveViewrecordResponse struct {
	taobao.TaobaoResponse
	Result QueryLiveViewRecordDetail `json:"result,omitempty"`
}
type UserViewRecordDetailModelList struct {
	LiveBeginUnixTime int64  `json:"live_begin_unix_time,omitempty"`
	LiveEndUnixTime   int64  `json:"live_end_unix_time,omitempty"`
	UserId            string `json:"user_id,omitempty"`
	UserName          string `json:"user_name,omitempty"`
	ViewType          int64  `json:"view_type,omitempty"`
	WatchTimeSecond   int64  `json:"watch_time_second,omitempty"`
}
type QueryLiveViewRecordDetail struct {
	DeptId     int64                           `json:"dept_id,omitempty"`
	GroupName  string                          `json:"group_name,omitempty"`
	HasMore    bool                            `json:"has_more,omitempty"`
	LiveUuid   string                          `json:"live_uuid,omitempty"`
	RecordList []UserViewRecordDetailModelList `json:"record_list,omitempty"`
	Title      string                          `json:"title,omitempty"`
}
