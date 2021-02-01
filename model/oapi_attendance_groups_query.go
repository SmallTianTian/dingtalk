package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAttendanceGroupsQueryRequest() *OapiAttendanceGroupsQueryRequest {
	return &OapiAttendanceGroupsQueryRequest{}
}

type OapiAttendanceGroupsQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGroupsQueryResponse
	Cursor          string
	OpUserid        string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGroupsQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGroupsQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGroupsQueryRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiAttendanceGroupsQueryRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiAttendanceGroupsQueryRequest) SetOpUserid(opUserid2 string) {
	this.OpUserid = opUserid2
}
func (this *OapiAttendanceGroupsQueryRequest) GetOpUserid() string {
	return this.OpUserid
}
func (this *OapiAttendanceGroupsQueryRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAttendanceGroupsQueryRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAttendanceGroupsQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.groups.query"
}
func (this *OapiAttendanceGroupsQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGroupsQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGroupsQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGroupsQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGroupsQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["op_userid"] = this.OpUserid
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiAttendanceGroupsQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAttendanceGroupsQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGroupsQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGroupsQueryResponse struct {
	taobao.TaobaoResponse
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
