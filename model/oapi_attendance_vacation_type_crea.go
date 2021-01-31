package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceVacationTypeCreateRequest() *OapiAttendanceVacationTypeCreateRequest {
	return &OapiAttendanceVacationTypeCreateRequest{}
}

type OapiAttendanceVacationTypeCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceVacationTypeCreateResponse
	BizType         string
	Extras          string
	HoursInPerDay   int64
	LeaveName       string
	LeaveViewUnit   string
	NaturalDayLeave bool
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceVacationTypeCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceVacationTypeCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceVacationTypeCreateRequest) SetBizType(bizType2 string) {
	this.BizType = bizType2
}
func (this *OapiAttendanceVacationTypeCreateRequest) GetBizType() string {
	return this.BizType
}
func (this *OapiAttendanceVacationTypeCreateRequest) SetExtras(extras2 string) {
	this.Extras = extras2
}
func (this *OapiAttendanceVacationTypeCreateRequest) GetExtras() string {
	return this.Extras
}
func (this *OapiAttendanceVacationTypeCreateRequest) SetHoursInPerDay(hoursInPerDay2 int64) {
	this.HoursInPerDay = hoursInPerDay2
}
func (this *OapiAttendanceVacationTypeCreateRequest) GetHoursInPerDay() int64 {
	return this.HoursInPerDay
}
func (this *OapiAttendanceVacationTypeCreateRequest) SetLeaveName(leaveName2 string) {
	this.LeaveName = leaveName2
}
func (this *OapiAttendanceVacationTypeCreateRequest) GetLeaveName() string {
	return this.LeaveName
}
func (this *OapiAttendanceVacationTypeCreateRequest) SetLeaveViewUnit(leaveViewUnit2 string) {
	this.LeaveViewUnit = leaveViewUnit2
}
func (this *OapiAttendanceVacationTypeCreateRequest) GetLeaveViewUnit() string {
	return this.LeaveViewUnit
}
func (this *OapiAttendanceVacationTypeCreateRequest) SetNaturalDayLeave(naturalDayLeave2 bool) {
	this.NaturalDayLeave = naturalDayLeave2
}
func (this *OapiAttendanceVacationTypeCreateRequest) GetNaturalDayLeave() bool {
	return this.NaturalDayLeave
}
func (this *OapiAttendanceVacationTypeCreateRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceVacationTypeCreateRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceVacationTypeCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.vacation.type.create"
}
func (this *OapiAttendanceVacationTypeCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceVacationTypeCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceVacationTypeCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceVacationTypeCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceVacationTypeCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_type"] = this.BizType
	txtParams["extras"] = this.Extras
	txtParams["hours_in_per_day"] = this.HoursInPerDay
	txtParams["leave_name"] = this.LeaveName
	txtParams["leave_view_unit"] = this.LeaveViewUnit
	txtParams["natural_day_leave"] = this.NaturalDayLeave
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *OapiAttendanceVacationTypeCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizType, "bizType"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceVacationTypeCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceVacationTypeCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceVacationTypeCreateResponse struct {
	taobao.TaobaoResponse
	Result  LeaveTypeVo `json:"result,omitempty"`
	Success bool        `json:"success,omitempty"`
}
type LeaveTypeVo struct {
	BizType         string `json:"biz_type,omitempty"`
	HoursInPerDay   int64  `json:"hours_in_per_day,omitempty"`
	LeaveCode       string `json:"leave_code,omitempty"`
	LeaveName       string `json:"leave_name,omitempty"`
	LeaveViewUnit   string `json:"leave_view_unit,omitempty"`
	NaturalDayLeave bool   `json:"natural_day_leave,omitempty"`
}
