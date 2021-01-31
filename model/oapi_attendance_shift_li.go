package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceShiftListRequest() *OapiAttendanceShiftListRequest {
	return &OapiAttendanceShiftListRequest{}
}

type OapiAttendanceShiftListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceShiftListResponse
	Cursor          int64
	OpUserId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceShiftListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceShiftListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceShiftListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiAttendanceShiftListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiAttendanceShiftListRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceShiftListRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceShiftListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.shift.list"
}
func (this *OapiAttendanceShiftListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceShiftListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceShiftListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceShiftListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceShiftListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["op_user_id"] = this.OpUserId
	return txtParams
}
func (this *OapiAttendanceShiftListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserId, "opUserId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceShiftListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceShiftListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceShiftListResponse struct {
	taobao.TaobaoResponse
	Errcode int64      `json:"errcode,omitempty"`
	Errmsg  string     `json:"errmsg,omitempty"`
	Result  PageResult `json:"result,omitempty"`
	Success bool       `json:"success,omitempty"`
}
type TopMinimalismShiftVo struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
