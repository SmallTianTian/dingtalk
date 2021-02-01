package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupMemberUpdateRequest() *OapiAttendanceGroupMemberUpdateRequest {
	return &OapiAttendanceGroupMemberUpdateRequest{}
}

type OapiAttendanceGroupMemberUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupMemberUpdateResponse
	GroupId         int64
	OpUserId        string
	ScheduleFlag    int64
	TopHttpMethod   string
	TopResponseType string
	UpdateParam     string
}

func (this *OapiAttendanceGroupMemberUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupMemberUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupMemberUpdateRequest) SetGroupId(groupId2 int64) {
	this.GroupId = groupId2
}
func (this *OapiAttendanceGroupMemberUpdateRequest) GetGroupId() int64 {
	return this.GroupId
}
func (this *OapiAttendanceGroupMemberUpdateRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceGroupMemberUpdateRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceGroupMemberUpdateRequest) SetScheduleFlag(scheduleFlag2 int64) {
	this.ScheduleFlag = scheduleFlag2
}
func (this *OapiAttendanceGroupMemberUpdateRequest) GetScheduleFlag() int64 {
	return this.ScheduleFlag
}
func (this *OapiAttendanceGroupMemberUpdateRequest) SetUpdateParam(updateParam2 string) {
	this.UpdateParam = updateParam2
}
func (this *OapiAttendanceGroupMemberUpdateRequest) GetUpdateParam() string {
	return this.UpdateParam
}
func (this *OapiAttendanceGroupMemberUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.member.update"
}
func (this *OapiAttendanceGroupMemberUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupMemberUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupMemberUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupMemberUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupMemberUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_id"] = this.GroupId
	txtParams["op_user_id"] = this.OpUserId
	txtParams["schedule_flag"] = this.ScheduleFlag
	txtParams["update_param"] = this.UpdateParam
	return txtParams
}
func (this *OapiAttendanceGroupMemberUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupId, "groupId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupMemberUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupMemberUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopGroupMemberUpdateParam struct {
	AddDepts         []string `json:"add_depts,omitempty"`
	AddExtraUsers    []string `json:"add_extra_users,omitempty"`
	AddUsers         []string `json:"add_users,omitempty"`
	RemoveDepts      []string `json:"remove_depts,omitempty"`
	RemoveExtraUsers []string `json:"remove_extra_users,omitempty"`
	RemoveUsers      []string `json:"remove_users,omitempty"`
}
type OapiAttendanceGroupMemberUpdateResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
