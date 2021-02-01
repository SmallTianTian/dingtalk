package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAttendanceGroupMemberusersListRequest() *OapiAttendanceGroupMemberusersListRequest {
	return &OapiAttendanceGroupMemberusersListRequest{}
}

type OapiAttendanceGroupMemberusersListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupMemberusersListResponse
	Cursor          int64
	GroupId         int64
	OpUserId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupMemberusersListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupMemberusersListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupMemberusersListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiAttendanceGroupMemberusersListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiAttendanceGroupMemberusersListRequest) SetGroupId(groupId2 int64) {
	this.GroupId = groupId2
}
func (this *OapiAttendanceGroupMemberusersListRequest) GetGroupId() int64 {
	return this.GroupId
}
func (this *OapiAttendanceGroupMemberusersListRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceGroupMemberusersListRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceGroupMemberusersListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.memberusers.list"
}
func (this *OapiAttendanceGroupMemberusersListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupMemberusersListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupMemberusersListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupMemberusersListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupMemberusersListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["group_id"] = this.GroupId
	txtParams["op_user_id"] = this.OpUserId
	return txtParams
}
func (this *OapiAttendanceGroupMemberusersListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAttendanceGroupMemberusersListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupMemberusersListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupMemberusersListResponse struct {
	taobao.TaobaoResponse
	Result  PageResult `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
