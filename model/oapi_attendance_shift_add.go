package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiAttendanceShiftAddRequest() *OapiAttendanceShiftAddRequest {
	return &OapiAttendanceShiftAddRequest{}
}

type OapiAttendanceShiftAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceShiftAddResponse
	OpUserId        string
	Shift           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceShiftAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceShiftAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceShiftAddRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceShiftAddRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceShiftAddRequest) SetShift(shift2 string) {
	this.Shift = shift2
}
func (this *OapiAttendanceShiftAddRequest) GetShift() string {
	return this.Shift
}
func (this *OapiAttendanceShiftAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.shift.add"
}
func (this *OapiAttendanceShiftAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceShiftAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceShiftAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceShiftAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceShiftAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["op_user_id"] = this.OpUserId
	txtParams["shift"] = this.Shift
	return txtParams
}
func (this *OapiAttendanceShiftAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserId, "opUserId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceShiftAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceShiftAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopAtTimeVo struct {
	Across      int64     `json:"across,omitempty"`
	BeginMin    int64     `json:"begin_min,omitempty"`
	CheckTime   time.Time `json:"check_time,omitempty"`
	CheckType   string    `json:"check_type,omitempty"`
	EndMin      int64     `json:"end_min,omitempty"`
	FlexMinutes []int64   `json:"flex_minutes,omitempty"`
	FreeCheck   bool      `json:"free_check,omitempty"`
}
type TopAtSectionVo struct {
	Times []TopAtTimeVo `json:"times,omitempty"`
}
type TopAtClassSettingVo struct {
	AbsenteeismLateMinutes int64       `json:"absenteeism_late_minutes,omitempty"`
	ClassId                int64       `json:"class_id,omitempty"`
	CorpId                 string      `json:"corp_id,omitempty"`
	DemandWorkTimeMinutes  int64       `json:"demand_work_time_minutes,omitempty"`
	Extras                 string      `json:"extras,omitempty"`
	IsDeleted              string      `json:"is_deleted,omitempty"`
	IsFlexible             bool        `json:"is_flexible,omitempty"`
	RestBeginTime          TopAtTimeVo `json:"rest_begin_time,omitempty"`
	RestEndTime            TopAtTimeVo `json:"rest_end_time,omitempty"`
	SeriousLateMinutes     int64       `json:"serious_late_minutes,omitempty"`
	Tags                   string      `json:"tags,omitempty"`
}
type TopAtClassVo struct {
	ClassGroupName string              `json:"class_group_name,omitempty"`
	CorpId         string              `json:"corp_id,omitempty"`
	Id             int64               `json:"id,omitempty"`
	Name           string              `json:"name,omitempty"`
	Owner          string              `json:"owner,omitempty"`
	Sections       []TopAtSectionVo    `json:"sections,omitempty"`
	ServiceId      int64               `json:"service_id,omitempty"`
	Setting        TopAtClassSettingVo `json:"setting,omitempty"`
}
type OapiAttendanceShiftAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64        `json:"errcode,omitempty"`
	Errmsg  string       `json:"errmsg,omitempty"`
	Result  TopAtClassVo `json:"result,omitempty"`
	Success bool         `json:"success,omitempty"`
}
