package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupModifyRequest() *OapiAttendanceGroupModifyRequest {
	return &OapiAttendanceGroupModifyRequest{}
}

type OapiAttendanceGroupModifyRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupModifyResponse
	OpUserId        string
	TopGroup        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupModifyRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupModifyRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupModifyRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceGroupModifyRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceGroupModifyRequest) SetTopGroup(topGroup2 string) {
	this.TopGroup = topGroup2
}
func (this *OapiAttendanceGroupModifyRequest) GetTopGroup() string {
	return this.TopGroup
}
func (this *OapiAttendanceGroupModifyRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.modify"
}
func (this *OapiAttendanceGroupModifyRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupModifyRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupModifyRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupModifyRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupModifyRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["op_user_id"] = this.OpUserId
	txtParams["top_group"] = this.TopGroup
	return txtParams
}
func (this *OapiAttendanceGroupModifyRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserId, "opUserId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupModifyRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupModifyRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type Resourcepermissionmap struct {
	CameraCheck       string `json:"camera_check,omitempty"`
	CheckPositionType string `json:"check_position_type,omitempty"`
	CheckTime         string `json:"check_time,omitempty"`
	GroupMember       string `json:"group_member,omitempty"`
	GroupType         string `json:"group_type,omitempty"`
	OutSideCheck      string `json:"out_side_check,omitempty"`
	OverTimeRule      string `json:"over_time_rule,omitempty"`
	Schedule          string `json:"schedule,omitempty"`
}
type OapiAttendanceGroupModifyResponse struct {
	taobao.TaobaoResponse
	Result  TopGroupVo `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
