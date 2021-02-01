package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceVacationQuotaListRequest() *OapiAttendanceVacationQuotaListRequest {
	return &OapiAttendanceVacationQuotaListRequest{}
}

type OapiAttendanceVacationQuotaListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceVacationQuotaListResponse
	LeaveCode       string
	Offset          int64
	OpUserid        string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
	Userids         string
}

func (this *OapiAttendanceVacationQuotaListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceVacationQuotaListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceVacationQuotaListRequest) SetLeaveCode(leaveCode2 string) {
	this.LeaveCode = leaveCode2
}
func (this *OapiAttendanceVacationQuotaListRequest) GetLeaveCode() string {
	return this.LeaveCode
}
func (this *OapiAttendanceVacationQuotaListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiAttendanceVacationQuotaListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiAttendanceVacationQuotaListRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceVacationQuotaListRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceVacationQuotaListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAttendanceVacationQuotaListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAttendanceVacationQuotaListRequest) SetUserids(userids2 string) {
	this.Userids = userids2
}
func (this *OapiAttendanceVacationQuotaListRequest) GetUserids() string {
	return this.Userids
}
func (this *OapiAttendanceVacationQuotaListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.vacation.quota.list"
}
func (this *OapiAttendanceVacationQuotaListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceVacationQuotaListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceVacationQuotaListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceVacationQuotaListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceVacationQuotaListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["leave_code"] = this.LeaveCode
	txtParams["offset"] = this.Offset
	txtParams["op_userid"] = this.OpUserid
	txtParams["size"] = this.Size
	txtParams["userids"] = this.Userids
	return txtParams
}
func (this *OapiAttendanceVacationQuotaListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.LeaveCode, "leaveCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceVacationQuotaListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceVacationQuotaListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceVacationQuotaListResponse struct {
	taobao.TaobaoResponse
	Result  OapiLeaveQuotaUserListVo `json:"result,omitempty"`
	Success bool                     `json:"success,omitempty"`
}
type Leavequotas struct {
	EndTime         int64  `json:"end_time,omitempty"`
	LeaveCode       string `json:"leave_code,omitempty"`
	QuotaCycle      string `json:"quota_cycle,omitempty"`
	QuotaId         string `json:"quota_id,omitempty"`
	QuotaNumPerDay  int64  `json:"quota_num_per_day,omitempty"`
	QuotaNumPerHour int64  `json:"quota_num_per_hour,omitempty"`
	StartTime       int64  `json:"start_time,omitempty"`
	UsedNumPerDay   int64  `json:"used_num_per_day,omitempty"`
	UsedNumPerHour  int64  `json:"used_num_per_hour,omitempty"`
	Userid          string `json:"userid,omitempty"`
}
type OapiLeaveQuotaUserListVo struct {
	HasMore     bool          `json:"has_more,omitempty"`
	LeaveQuotas []Leavequotas `json:"leave_quotas,omitempty"`
}
