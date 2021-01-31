package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAttendanceGroupUpdateRequest() *OapiAttendanceGroupUpdateRequest {
	return &OapiAttendanceGroupUpdateRequest{}
}

type OapiAttendanceGroupUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupUpdateResponse
	Group           string
	OpUserid        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupUpdateRequest) SetGroup(group2 string) {
	this.Group = group2
}
func (this *OapiAttendanceGroupUpdateRequest) GetGroup() string {
	return this.Group
}
func (this *OapiAttendanceGroupUpdateRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupUpdateRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.group.update"
}
func (this *OapiAttendanceGroupUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["group"] = this.Group
	txtParams["op_userid"] = this.OpUserid
	return txtParams
}
func (this *OapiAttendanceGroupUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAttendanceGroupUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  Group  `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
