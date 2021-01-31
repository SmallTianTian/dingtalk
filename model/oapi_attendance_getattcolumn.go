package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAttendanceGetattcolumnsRequest() *OapiAttendanceGetattcolumnsRequest {
	return &OapiAttendanceGetattcolumnsRequest{}
}

type OapiAttendanceGetattcolumnsRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAttendanceGetattcolumnsResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAttendanceGetattcolumnsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAttendanceGetattcolumnsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAttendanceGetattcolumnsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.attendance.getattcolumns"
}
func (this *OapiAttendanceGetattcolumnsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAttendanceGetattcolumnsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAttendanceGetattcolumnsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAttendanceGetattcolumnsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAttendanceGetattcolumnsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiAttendanceGetattcolumnsRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAttendanceGetattcolumnsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAttendanceGetattcolumnsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAttendanceGetattcolumnsResponse struct {
	taobao.TaobaoResponse
	Result AttColumnsForTopVo `json:"result,omitempty"`
}
type ColumnForTopVo struct {
	Alias        string `json:"alias,omitempty"`
	ExpressionId int64  `json:"expression_id,omitempty"`
	Extension    string `json:"extension,omitempty"`
	Id           int64  `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Status       int64  `json:"status,omitempty"`
	SubType      int64  `json:"sub_type,omitempty"`
	Type         int64  `json:"type,omitempty"`
}
type AttColumnsForTopVo struct {
	Columns []ColumnForTopVo `json:"columns,omitempty"`
}
