package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupAddRequest() *OapiAttendanceGroupAddRequest {
	return &OapiAttendanceGroupAddRequest{}
}

type OapiAttendanceGroupAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupAddResponse
	OpUserId        string
	TopGroup        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupAddRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceGroupAddRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceGroupAddRequest) SetTopGroup(topGroup2 string) {
	this.TopGroup = topGroup2
}
func (this *OapiAttendanceGroupAddRequest) GetTopGroup() string {
	return this.TopGroup
}
func (this *OapiAttendanceGroupAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.add"
}
func (this *OapiAttendanceGroupAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["op_user_id"] = this.OpUserId
	txtParams["top_group"] = this.TopGroup
	return txtParams
}
func (this *OapiAttendanceGroupAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserId, "opUserId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopPositionVo struct {
	Accuracy  string `json:"accuracy,omitempty"`
	Address   string `json:"address,omitempty"`
	CorpId    string `json:"corp_id,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Title     string `json:"title,omitempty"`
}
type TopShiftVo struct {
	Id int64 `json:"id,omitempty"`
}
type TopMemberVo struct {
	CorpId string `json:"corp_id,omitempty"`
	Role   string `json:"role,omitempty"`
	Type   string `json:"type,omitempty"`
	UserId string `json:"user_id,omitempty"`
}
type TopGroupManageRolePermissionVo struct {
	CameraCheck       string `json:"camera_check,omitempty"`
	CheckPositionType string `json:"check_position_type,omitempty"`
	CheckTime         string `json:"check_time,omitempty"`
	GroupMember       string `json:"group_member,omitempty"`
	GroupType         string `json:"group_type,omitempty"`
	OutSideCheck      string `json:"out_side_check,omitempty"`
	OverTimeRule      string `json:"over_time_rule,omitempty"`
	Schedule          string `json:"schedule,omitempty"`
}
type TopWifiVo struct {
	CorpId  string `json:"corp_id,omitempty"`
	MacAddr string `json:"mac_addr,omitempty"`
	Ssid    string `json:"ssid,omitempty"`
}
type TopGroupVo struct {
	CheckNeedHealthyCode        bool                           `json:"check_need_healthy_code,omitempty"`
	CorpId                      string                         `json:"corp_id,omitempty"`
	DefaultClassId              int64                          `json:"default_class_id,omitempty"`
	DisableCheckWithoutSchedule bool                           `json:"disable_check_without_schedule,omitempty"`
	EnableCameraCheck           bool                           `json:"enable_camera_check,omitempty"`
	EnableEmpSelectClass        bool                           `json:"enable_emp_select_class,omitempty"`
	EnableFaceCheck             bool                           `json:"enable_face_check,omitempty"`
	EnableNextDay               bool                           `json:"enable_next_day,omitempty"`
	EnableOutsideCameraCheck    bool                           `json:"enable_outside_camera_check,omitempty"`
	EnableOutsideCheck          bool                           `json:"enable_outside_check,omitempty"`
	FreecheckDayStartMinOffset  int64                          `json:"freecheck_day_start_min_offset,omitempty"`
	FreecheckWorkDays           []int64                        `json:"freecheck_work_days,omitempty"`
	Id                          int64                          `json:"id,omitempty"`
	ManagerList                 []string                       `json:"manager_list,omitempty"`
	Members                     []TopMemberVo                  `json:"members,omitempty"`
	ModifyMember                bool                           `json:"modify_member,omitempty"`
	Name                        string                         `json:"name,omitempty"`
	Offset                      int64                          `json:"offset,omitempty"`
	Owner                       string                         `json:"owner,omitempty"`
	Positions                   []TopPositionVo                `json:"positions,omitempty"`
	ResourcePermissionMap       TopGroupManageRolePermissionVo `json:"resource_permission_map,omitempty"`
	ShiftVoList                 []TopShiftVo                   `json:"shift_vo_list,omitempty"`
	SkipHolidays                bool                           `json:"skip_holidays,omitempty"`
	SpecialDays                 string                         `json:"special_days,omitempty"`
	Type                        string                         `json:"type,omitempty"`
	Wifis                       []TopWifiVo                    `json:"wifis,omitempty"`
	WorkdayClassList            []int64                        `json:"workday_class_list,omitempty"`
}
type OapiAttendanceGroupAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64      `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Result  TopGroupVo `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
