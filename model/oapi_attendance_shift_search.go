package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAttendanceShiftSearchRequest() *OapiAttendanceShiftSearchRequest {
	return &OapiAttendanceShiftSearchRequest{}
}

type OapiAttendanceShiftSearchRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceShiftSearchResponse
	OpUserId        string
	ShiftName       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceShiftSearchRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceShiftSearchRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceShiftSearchRequest) SetOpUserId(opUserId2 string) {
	this.OpUserId = opUserId2
}
func (this *OapiAttendanceShiftSearchRequest) GetOpUserId() string {
	return this.OpUserId
}
func (this *OapiAttendanceShiftSearchRequest) SetShiftName(shiftName2 string) {
	this.ShiftName = shiftName2
}
func (this *OapiAttendanceShiftSearchRequest) GetShiftName() string {
	return this.ShiftName
}
func (this *OapiAttendanceShiftSearchRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.shift.search"
}
func (this *OapiAttendanceShiftSearchRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceShiftSearchRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceShiftSearchRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceShiftSearchRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceShiftSearchRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["op_user_id"] = this.OpUserId
	txtParams["shift_name"] = this.ShiftName
	return txtParams
}
func (this *OapiAttendanceShiftSearchRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpUserId, "opUserId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAttendanceShiftSearchRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceShiftSearchRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceShiftSearchResponse struct {
	taobao.TaobaoResponse
	Result  []TopMinimalismShiftVO `json:"result,omitempty"`
	Success bool                   `json:"success,omitempty"`
}
type TopMinimalismShiftVO struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
