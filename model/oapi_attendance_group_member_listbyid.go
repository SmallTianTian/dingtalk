package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupMemberListbyidsRequest() *OapiAttendanceGroupMemberListbyidsRequest {
	return &OapiAttendanceGroupMemberListbyidsRequest{}
}

type OapiAttendanceGroupMemberListbyidsRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupMemberListbyidsResponse
	GroupId         int64
	MemberIds       string
	MemberType      int64
	OpUserId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupMemberListbyidsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) SetGroupId(groupId2 int64) {
	this.GroupId = groupId2
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) GetGroupId() int64 {
	return this.GroupId
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) SetMemberIds(memberIds2 string) {
	this.MemberIds = memberIds2
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) GetMemberIds() string {
	return this.MemberIds
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) SetMemberType(memberType2 int64) {
	this.MemberType = memberType2
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) GetMemberType() int64 {
	return this.MemberType
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.member.listbyids"
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_id"] = this.GroupId
	txtParams["member_ids"] = this.MemberIds
	txtParams["member_type"] = this.MemberType
	txtParams["op_user_id"] = this.OpUserId
	return txtParams
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GroupId, "groupId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupMemberListbyidsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupMemberListbyidsResponse struct {
	taobao.TaobaoResponse
	Result  []string `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
