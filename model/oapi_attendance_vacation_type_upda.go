package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceVacationTypeUpdateRequest() *OapiAttendanceVacationTypeUpdateRequest {
	return &OapiAttendanceVacationTypeUpdateRequest{}
}

type OapiAttendanceVacationTypeUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceVacationTypeUpdateResponse
	BizType         string
	Extras          string
	HoursInPerDay   int64
	LeaveCode       string
	LeaveName       string
	LeaveViewUnit   string
	NaturalDayLeave bool
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceVacationTypeUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceVacationTypeUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceVacationTypeUpdateRequest) SetBizType(bizType2 string) {
	this.BizType = bizType2
}
func (this *OapiAttendanceVacationTypeUpdateRequest) GetBizType() string {
	return this.BizType
}
func (this *OapiAttendanceVacationTypeUpdateRequest) SetExtras(extras2 string) {
	this.Extras = extras2
}
func (this *OapiAttendanceVacationTypeUpdateRequest) GetExtras() string {
	return this.Extras
}
func (this *OapiAttendanceVacationTypeUpdateRequest) SetHoursInPerDay(hoursInPerDay2 int64) {
	this.HoursInPerDay = hoursInPerDay2
}
func (this *OapiAttendanceVacationTypeUpdateRequest) GetHoursInPerDay() int64 {
	return this.HoursInPerDay
}
func (this *OapiAttendanceVacationTypeUpdateRequest) SetLeaveCode(leaveCode2 string) {
	this.LeaveCode = leaveCode2
}
func (this *OapiAttendanceVacationTypeUpdateRequest) GetLeaveCode() string {
	return this.LeaveCode
}
func (this *OapiAttendanceVacationTypeUpdateRequest) SetLeaveName(leaveName2 string) {
	this.LeaveName = leaveName2
}
func (this *OapiAttendanceVacationTypeUpdateRequest) GetLeaveName() string {
	return this.LeaveName
}
func (this *OapiAttendanceVacationTypeUpdateRequest) SetLeaveViewUnit(leaveViewUnit2 string) {
	this.LeaveViewUnit = leaveViewUnit2
}
func (this *OapiAttendanceVacationTypeUpdateRequest) GetLeaveViewUnit() string {
	return this.LeaveViewUnit
}
func (this *OapiAttendanceVacationTypeUpdateRequest) SetNaturalDayLeave(naturalDayLeave2 bool) {
	this.NaturalDayLeave = naturalDayLeave2
}
func (this *OapiAttendanceVacationTypeUpdateRequest) GetNaturalDayLeave() bool {
	return this.NaturalDayLeave
}
func (this *OapiAttendanceVacationTypeUpdateRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceVacationTypeUpdateRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceVacationTypeUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.vacation.type.update"
}
func (this *OapiAttendanceVacationTypeUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceVacationTypeUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceVacationTypeUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceVacationTypeUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceVacationTypeUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_type"] = this.BizType
	txtParams["extras"] = this.Extras
	txtParams["hours_in_per_day"] = this.HoursInPerDay
	txtParams["leave_code"] = this.LeaveCode
	txtParams["leave_name"] = this.LeaveName
	txtParams["leave_view_unit"] = this.LeaveViewUnit
	txtParams["natural_day_leave"] = this.NaturalDayLeave
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *OapiAttendanceVacationTypeUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.LeaveCode, "leaveCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceVacationTypeUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceVacationTypeUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceVacationTypeUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64       `json:"errcode,omitempty"`
	Errmsg  string      `json:"errmsg,omitempty"`
	Result  LeaveTypeVo `json:"result,omitempty"`
	Success bool        `json:"success,omitempty"`
}
