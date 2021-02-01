package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiReportListRequest() *OapiReportListRequest {
	return &OapiReportListRequest{}
}

type OapiReportListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp              OapiReportListResponse
	Cursor            int64
	EndTime           int64
	ModifiedEndTime   int64
	ModifiedStartTime int64
	Size              int64
	StartTime         int64
	TemplateName      string
	TopHttpMethod     string
	TopResponseType   string
	Userid            string
}

func (this *OapiReportListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiReportListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiReportListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiReportListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiReportListRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *OapiReportListRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *OapiReportListRequest) SetModifiedEndTime(modifiedEndTime2 int64) {
	this.ModifiedEndTime = modifiedEndTime2
}
func (this *OapiReportListRequest) GetModifiedEndTime() int64 {
	return this.ModifiedEndTime
}
func (this *OapiReportListRequest) SetModifiedStartTime(modifiedStartTime2 int64) {
	this.ModifiedStartTime = modifiedStartTime2
}
func (this *OapiReportListRequest) GetModifiedStartTime() int64 {
	return this.ModifiedStartTime
}
func (this *OapiReportListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiReportListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiReportListRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *OapiReportListRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *OapiReportListRequest) SetTemplateName(templateName2 string) {
	this.TemplateName = templateName2
}
func (this *OapiReportListRequest) GetTemplateName() string {
	return this.TemplateName
}
func (this *OapiReportListRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiReportListRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiReportListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.report.list"
}
func (this *OapiReportListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiReportListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiReportListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiReportListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiReportListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["end_time"] = this.EndTime
	txtParams["modified_end_time"] = this.ModifiedEndTime
	txtParams["modified_start_time"] = this.ModifiedStartTime
	txtParams["size"] = this.Size
	txtParams["start_time"] = this.StartTime
	txtParams["template_name"] = this.TemplateName
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiReportListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *OapiReportListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiReportListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiReportListResponse struct {
	taobao.TaobaoResponse
	Result PageVo `json:"result,omitempty"`
}
