package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceGroupMinimalismListRequest() *OapiAttendanceGroupMinimalismListRequest {
	return &OapiAttendanceGroupMinimalismListRequest{}
}

type OapiAttendanceGroupMinimalismListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupMinimalismListResponse
	Cursor          int64
	OpUserId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupMinimalismListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupMinimalismListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupMinimalismListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiAttendanceGroupMinimalismListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiAttendanceGroupMinimalismListRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceGroupMinimalismListRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceGroupMinimalismListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.minimalism.list"
}
func (this *OapiAttendanceGroupMinimalismListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupMinimalismListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupMinimalismListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupMinimalismListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupMinimalismListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["op_user_id"] = this.OpUserId
	return txtParams
}
func (this *OapiAttendanceGroupMinimalismListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserId, "opUserId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceGroupMinimalismListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupMinimalismListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupMinimalismListResponse struct {
	taobao.TaobaoResponse
	Result  PageResult `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
type TopMinimalismGroupVo struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
