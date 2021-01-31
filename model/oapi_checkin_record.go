package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCheckinRecordRequest() *OapiCheckinRecordRequest {
	return &OapiCheckinRecordRequest{}
}

type OapiCheckinRecordRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCheckinRecordResponse
	DepartmentId    string
	EndTime         int64
	Offset          int64
	Order           string
	Size            int64
	StartTime       int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCheckinRecordRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiCheckinRecordRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCheckinRecordRequest) SetDepartmentId(departmentId2 string) {
	this.DepartmentId = departmentId2
}
func (this *OapiCheckinRecordRequest) GetDepartmentId() string {
	return this.DepartmentId
}
func (this *OapiCheckinRecordRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *OapiCheckinRecordRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *OapiCheckinRecordRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiCheckinRecordRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiCheckinRecordRequest) SetOrder(order2 string) {
	this.Order = order2
}
func (this *OapiCheckinRecordRequest) GetOrder() string {
	return this.Order
}
func (this *OapiCheckinRecordRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiCheckinRecordRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiCheckinRecordRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *OapiCheckinRecordRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *OapiCheckinRecordRequest) GetApiMethodName() string {
	return "dingtalk.oapi.checkin.record"
}
func (this *OapiCheckinRecordRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCheckinRecordRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCheckinRecordRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCheckinRecordRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCheckinRecordRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["department_id"] = this.DepartmentId
	txtParams["end_time"] = this.EndTime
	txtParams["offset"] = this.Offset
	txtParams["order"] = this.Order
	txtParams["size"] = this.Size
	txtParams["start_time"] = this.StartTime
	return txtParams
}
func (this *OapiCheckinRecordRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCheckinRecordRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCheckinRecordRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCheckinRecordResponse struct {
	taobao.TaobaoResponse
	Data    []Data `json:"data,omitempty"`
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
