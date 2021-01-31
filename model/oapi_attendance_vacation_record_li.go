package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceVacationRecordListRequest() *OapiAttendanceVacationRecordListRequest {
	return &OapiAttendanceVacationRecordListRequest{}
}

type OapiAttendanceVacationRecordListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceVacationRecordListResponse
	LeaveCode       string
	Offset          int64
	OpUserid        string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
	Userids         string
}

func (this *OapiAttendanceVacationRecordListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceVacationRecordListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceVacationRecordListRequest) SetLeaveCode(leaveCode2 string) {
	this.LeaveCode = leaveCode2
}
func (this *OapiAttendanceVacationRecordListRequest) GetLeaveCode() string {
	return this.LeaveCode
}
func (this *OapiAttendanceVacationRecordListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiAttendanceVacationRecordListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiAttendanceVacationRecordListRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceVacationRecordListRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceVacationRecordListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAttendanceVacationRecordListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAttendanceVacationRecordListRequest) SetUserids(userids2 string) {
	this.Userids = userids2
}
func (this *OapiAttendanceVacationRecordListRequest) GetUserids() string {
	return this.Userids
}
func (this *OapiAttendanceVacationRecordListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.vacation.record.list"
}
func (this *OapiAttendanceVacationRecordListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceVacationRecordListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceVacationRecordListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceVacationRecordListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceVacationRecordListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["leave_code"] = this.LeaveCode
	txtParams["offset"] = this.Offset
	txtParams["op_userid"] = this.OpUserid
	txtParams["size"] = this.Size
	txtParams["userids"] = this.Userids
	return txtParams
}
func (this *OapiAttendanceVacationRecordListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.LeaveCode, "leaveCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceVacationRecordListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceVacationRecordListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceVacationRecordListResponse struct {
	taobao.TaobaoResponse
	Errcode int64                 `json:"errcode,omitempty"`
	Errmsg  string                `json:"errmsg,omitempty"`
	Result  OapiLeaveRecordListVo `json:"result,omitempty"`
	Success bool                  `json:"success,omitempty"`
}
type OapiLeaveRecordVo struct {
	CalType          string `json:"cal_type,omitempty"`
	EndTime          int64  `json:"end_time,omitempty"`
	LeaveCode        string `json:"leave_code,omitempty"`
	LeaveReason      string `json:"leave_reason,omitempty"`
	LeaveRecordType  string `json:"leave_record_type,omitempty"`
	LeaveStatus      string `json:"leave_status,omitempty"`
	LeaveViewUnit    string `json:"leave_view_unit,omitempty"`
	ParentRecordId   string `json:"parent_record_id,omitempty"`
	QuotaId          string `json:"quota_id,omitempty"`
	RecordId         string `json:"record_id,omitempty"`
	RecordNumPerDay  int64  `json:"record_num_per_day,omitempty"`
	RecordNumPerHour int64  `json:"record_num_per_hour,omitempty"`
	StartTime        int64  `json:"start_time,omitempty"`
	Userid           string `json:"userid,omitempty"`
}
type OapiLeaveRecordListVo struct {
	HasMore      bool                `json:"has_more,omitempty"`
	LeaveRecords []OapiLeaveRecordVo `json:"leave_records,omitempty"`
}
