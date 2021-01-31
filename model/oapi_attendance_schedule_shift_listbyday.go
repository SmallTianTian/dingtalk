package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceScheduleShiftListbydaysRequest() *OapiAttendanceScheduleShiftListbydaysRequest {
	return &OapiAttendanceScheduleShiftListbydaysRequest{}
}

type OapiAttendanceScheduleShiftListbydaysRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceScheduleShiftListbydaysResponse
	FromDateTime    int64
	OpUserId        string
	ToDateTime      int64
	TopHttpMethod   string
	TopResponseType string
	Userids         string
}

func (this *OapiAttendanceScheduleShiftListbydaysRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) SetFromDateTime(fromDateTime2 int64) {
	this.FromDateTime = fromDateTime2
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) GetFromDateTime() int64 {
	return this.FromDateTime
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) SetToDateTime(toDateTime2 int64) {
	this.ToDateTime = toDateTime2
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) GetToDateTime() int64 {
	return this.ToDateTime
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) SetUserids(userids2 string) {
	this.Userids = userids2
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) GetUserids() string {
	return this.Userids
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.schedule.shift.listbydays"
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["from_date_time"] = this.FromDateTime
	txtParams["op_user_id"] = this.OpUserId
	txtParams["to_date_time"] = this.ToDateTime
	txtParams["userids"] = this.Userids
	return txtParams
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.FromDateTime, "fromDateTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceScheduleShiftListbydaysRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceScheduleShiftListbydaysResponse struct {
	taobao.TaobaoResponse
	Errcode int64                   `json:"errcode,omitempty"`
	Errmsg  string                  `json:"errmsg,omitempty"`
	Result  []TopDayScheduleShiftVo `json:"result,omitempty"`
	Success bool                    `json:"success,omitempty"`
}
type TopDayScheduleShiftVo struct {
	CorpId        string    `json:"corp_id,omitempty"`
	GroupId       int64     `json:"group_id,omitempty"`
	ShiftIds      []int64   `json:"shift_ids,omitempty"`
	ShiftNames    []string  `json:"shift_names,omitempty"`
	ShiftVersions []int64   `json:"shift_versions,omitempty"`
	Userid        string    `json:"userid,omitempty"`
	WorkDate      time.Time `json:"work_date,omitempty"`
}
