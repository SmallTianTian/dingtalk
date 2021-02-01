package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAttendanceGroupDeletebyidRequest() *OapiAttendanceGroupDeletebyidRequest {
	return &OapiAttendanceGroupDeletebyidRequest{}
}

type OapiAttendanceGroupDeletebyidRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupDeletebyidResponse
	GroupId         int64
	OpUserId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupDeletebyidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupDeletebyidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupDeletebyidRequest) SetGroupId(groupId2 int64) {
	this.GroupId = groupId2
}
func (this *OapiAttendanceGroupDeletebyidRequest) GetGroupId() int64 {
	return this.GroupId
}
func (this *OapiAttendanceGroupDeletebyidRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceGroupDeletebyidRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceGroupDeletebyidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.deletebyid"
}
func (this *OapiAttendanceGroupDeletebyidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupDeletebyidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupDeletebyidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupDeletebyidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupDeletebyidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group_id"] = this.GroupId
	txtParams["op_user_id"] = this.OpUserId
	return txtParams
}
func (this *OapiAttendanceGroupDeletebyidRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAttendanceGroupDeletebyidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupDeletebyidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupDeletebyidResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
