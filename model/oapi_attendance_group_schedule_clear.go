package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAttendanceGroupScheduleClearRequest() *OapiAttendanceGroupScheduleClearRequest {
	return &OapiAttendanceGroupScheduleClearRequest{}
}

type OapiAttendanceGroupScheduleClearRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupScheduleClearResponse
	OpUserid        string
	Param           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupScheduleClearRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupScheduleClearRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupScheduleClearRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupScheduleClearRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupScheduleClearRequest) SetParam(param2 string) {
	this.Param = param2
}
func (this *OapiAttendanceGroupScheduleClearRequest) GetParam() string {
	return this.Param
}
func (this *OapiAttendanceGroupScheduleClearRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.schedule.clear"
}
func (this *OapiAttendanceGroupScheduleClearRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupScheduleClearRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupScheduleClearRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupScheduleClearRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupScheduleClearRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["op_userid"] = this.OpUserid
	txtParams["param"] = this.Param
	return txtParams
}
func (this *OapiAttendanceGroupScheduleClearRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAttendanceGroupScheduleClearRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupScheduleClearRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TopClearScheduleParam struct {
	FromDate int64    `json:"from_date,omitempty"`
	GroupId  string   `json:"group_id,omitempty"`
	ToDate   int64    `json:"to_date,omitempty"`
	UserIds  []string `json:"user_ids,omitempty"`
}
type OapiAttendanceGroupScheduleClearResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
