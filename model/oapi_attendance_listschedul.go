package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceListscheduleRequest() *OapiAttendanceListscheduleRequest {
	return &OapiAttendanceListscheduleRequest{}
}

type OapiAttendanceListscheduleRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceListscheduleResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
	WorkDate        time.Time
}

func (this *OapiAttendanceListscheduleRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceListscheduleRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceListscheduleRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiAttendanceListscheduleRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiAttendanceListscheduleRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAttendanceListscheduleRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAttendanceListscheduleRequest) SetWorkDate(workDate2 time.Time) {
	this.WorkDate = workDate2
}
func (this *OapiAttendanceListscheduleRequest) GetWorkDate() time.Time {
	return this.WorkDate
}
func (this *OapiAttendanceListscheduleRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.listschedule"
}
func (this *OapiAttendanceListscheduleRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceListscheduleRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceListscheduleRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceListscheduleRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceListscheduleRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	txtParams["workDate"] = this.WorkDate
	return txtParams
}
func (this *OapiAttendanceListscheduleRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.WorkDate, "workDate"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceListscheduleRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceListscheduleRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceListscheduleResponse struct {
	taobao.TaobaoResponse
	Result AtScheduleListForTopVo `json:"result,omitempty"`
}
type AtScheduleForTopVo struct {
	ApproveId        int64     `json:"approve_id,omitempty"`
	ChangedCheckTime time.Time `json:"changed_check_time,omitempty"`
	CheckType        string    `json:"check_type,omitempty"`
	ClassId          int64     `json:"class_id,omitempty"`
	ClassSettingId   int64     `json:"class_setting_id,omitempty"`
	GroupId          int64     `json:"group_id,omitempty"`
	PlanCheckTime    time.Time `json:"plan_check_time,omitempty"`
	PlanId           int64     `json:"plan_id,omitempty"`
	Userid           string    `json:"userid,omitempty"`
}
type AtScheduleListForTopVo struct {
	HasMore   bool                 `json:"has_more,omitempty"`
	Schedules []AtScheduleForTopVo `json:"schedules,omitempty"`
}
