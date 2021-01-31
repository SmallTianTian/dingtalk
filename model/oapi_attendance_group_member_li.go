package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupMemberListRequest() *OapiAttendanceGroupMemberListRequest {
	return &OapiAttendanceGroupMemberListRequest{}
}

type OapiAttendanceGroupMemberListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupMemberListResponse
	Cursor          int64
	GroupId         int64
	OpUserId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupMemberListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupMemberListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupMemberListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiAttendanceGroupMemberListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiAttendanceGroupMemberListRequest) SetGroupId(groupId2 int64) {
	this.GroupId = groupId2
}
func (this *OapiAttendanceGroupMemberListRequest) GetGroupId() int64 {
	return this.GroupId
}
func (this *OapiAttendanceGroupMemberListRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceGroupMemberListRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceGroupMemberListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.member.list"
}
func (this *OapiAttendanceGroupMemberListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupMemberListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupMemberListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupMemberListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupMemberListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["group_id"] = this.GroupId
	txtParams["op_user_id"] = this.OpUserId
	return txtParams
}
func (this *OapiAttendanceGroupMemberListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupId, "groupId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupMemberListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupMemberListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupMemberListResponse struct {
	taobao.TaobaoResponse
	Errcode int64      `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Result  PageResult `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
type TopGroupMemberVo struct {
	AtcFlag  string `json:"atc_flag,omitempty"`
	MemberId string `json:"member_id,omitempty"`
	Type     string `json:"type,omitempty"`
}
