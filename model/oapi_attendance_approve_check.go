package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAttendanceApproveCheckRequest() *OapiAttendanceApproveCheckRequest {
	return &OapiAttendanceApproveCheckRequest{}
}

type OapiAttendanceApproveCheckRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceApproveCheckResponse
	ApproveId       string
	JumpUrl         string
	PunchCheckTime  string
	PunchId         int64
	TagName         string
	TopHttpMethod   string
	TopResponseType string
	UserCheckTime   string
	Userid          string
	WorkDate        string
}

func (this *OapiAttendanceApproveCheckRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceApproveCheckRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceApproveCheckRequest) SetApproveId(approveId2 string) {
	this.ApproveId = approveId2
}
func (this *OapiAttendanceApproveCheckRequest) GetApproveId() string {
	return this.ApproveId
}
func (this *OapiAttendanceApproveCheckRequest) SetJumpUrl(jumpUrl2 string) {
	this.JumpUrl = jumpUrl2
}
func (this *OapiAttendanceApproveCheckRequest) GetJumpUrl() string {
	return this.JumpUrl
}
func (this *OapiAttendanceApproveCheckRequest) SetPunchCheckTime(punchCheckTime2 string) {
	this.PunchCheckTime = punchCheckTime2
}
func (this *OapiAttendanceApproveCheckRequest) GetPunchCheckTime() string {
	return this.PunchCheckTime
}
func (this *OapiAttendanceApproveCheckRequest) SetPunchId(punchId2 int64) {
	this.PunchId = punchId2
}
func (this *OapiAttendanceApproveCheckRequest) GetPunchId() int64 {
	return this.PunchId
}
func (this *OapiAttendanceApproveCheckRequest) SetTagName(tagName2 string) {
	this.TagName = tagName2
}
func (this *OapiAttendanceApproveCheckRequest) GetTagName() string {
	return this.TagName
}
func (this *OapiAttendanceApproveCheckRequest) SetUserCheckTime(userCheckTime2 string) {
	this.UserCheckTime = userCheckTime2
}
func (this *OapiAttendanceApproveCheckRequest) GetUserCheckTime() string {
	return this.UserCheckTime
}
func (this *OapiAttendanceApproveCheckRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAttendanceApproveCheckRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAttendanceApproveCheckRequest) SetWorkDate(workDate2 string) {
	this.WorkDate = workDate2
}
func (this *OapiAttendanceApproveCheckRequest) GetWorkDate() string {
	return this.WorkDate
}
func (this *OapiAttendanceApproveCheckRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.approve.check"
}
func (this *OapiAttendanceApproveCheckRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceApproveCheckRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceApproveCheckRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceApproveCheckRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceApproveCheckRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["approve_id"] = this.ApproveId
	txtParams["jump_url"] = this.JumpUrl
	txtParams["punch_check_time"] = this.PunchCheckTime
	txtParams["punch_id"] = this.PunchId
	txtParams["tag_name"] = this.TagName
	txtParams["user_check_time"] = this.UserCheckTime
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["work_date"] = this.WorkDate
	return txtParams
}
func (this *OapiAttendanceApproveCheckRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAttendanceApproveCheckRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceApproveCheckRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceApproveCheckResponse struct {
	taobao.TaobaoResponse
}
