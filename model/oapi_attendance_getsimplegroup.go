package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"time"
)

func NewOapiAttendanceGetsimplegroupsRequest() *OapiAttendanceGetsimplegroupsRequest {
	return &OapiAttendanceGetsimplegroupsRequest{}
}

type OapiAttendanceGetsimplegroupsRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGetsimplegroupsResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGetsimplegroupsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGetsimplegroupsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGetsimplegroupsRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiAttendanceGetsimplegroupsRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiAttendanceGetsimplegroupsRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAttendanceGetsimplegroupsRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAttendanceGetsimplegroupsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.getsimplegroups"
}
func (this *OapiAttendanceGetsimplegroupsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGetsimplegroupsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGetsimplegroupsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGetsimplegroupsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGetsimplegroupsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiAttendanceGetsimplegroupsRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAttendanceGetsimplegroupsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGetsimplegroupsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGetsimplegroupsResponse struct {
	taobao.TaobaoResponse
	Result AtGroupListForTopVo `json:"result,omitempty"`
}
type ClassSettingVo struct {
	AbsenteeismLateMinutes int64    `json:"absenteeism_late_minutes,omitempty"`
	ClassSettingId         int64    `json:"class_setting_id,omitempty"`
	IsOffDutyFreeCheck     string   `json:"is_off_duty_free_check,omitempty"`
	PermitLateMinutes      int64    `json:"permit_late_minutes,omitempty"`
	RestBeginTime          AtTimeVo `json:"rest_begin_time,omitempty"`
	RestEndTime            AtTimeVo `json:"rest_end_time,omitempty"`
	SeriousLateMinutes     int64    `json:"serious_late_minutes,omitempty"`
	WorkTimeMinutes        int64    `json:"work_time_minutes,omitempty"`
}
type SetionTimeVO struct {
	Across    int64     `json:"across,omitempty"`
	CheckTime time.Time `json:"check_time,omitempty"`
	CheckType string    `json:"check_type,omitempty"`
}
type AtClassVo struct {
	ClassId   int64          `json:"class_id,omitempty"`
	ClassName string         `json:"class_name,omitempty"`
	Sections  []AtSectionVo  `json:"sections,omitempty"`
	Setting   ClassSettingVo `json:"setting,omitempty"`
}
type AtGroupForTopVo struct {
	ClassesList                []string    `json:"classes_list,omitempty"`
	DefaultClassId             int64       `json:"default_class_id,omitempty"`
	DeptIds                    []int64     `json:"dept_ids,omitempty"`
	DeptNameList               []string    `json:"dept_name_list,omitempty"`
	FreecheckDayStartMinOffset int64       `json:"freecheck_day_start_min_offset,omitempty"`
	FreecheckWorkDays          []int64     `json:"freecheck_work_days,omitempty"`
	GroupId                    int64       `json:"group_id,omitempty"`
	GroupName                  string      `json:"group_name,omitempty"`
	IsDefault                  bool        `json:"is_default,omitempty"`
	ManagerList                []string    `json:"manager_list,omitempty"`
	MemberCount                int64       `json:"member_count,omitempty"`
	OwnerUserId                string      `json:"owner_user_id,omitempty"`
	SelectedClass              []AtClassVo `json:"selected_class,omitempty"`
	Type                       string      `json:"type,omitempty"`
	UserIds                    []string    `json:"user_ids,omitempty"`
	WorkDayList                []string    `json:"work_day_list,omitempty"`
}
type AtGroupListForTopVo struct {
	Groups  []AtGroupForTopVo `json:"groups,omitempty"`
	HasMore bool              `json:"has_more,omitempty"`
}
