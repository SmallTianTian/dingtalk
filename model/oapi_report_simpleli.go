package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiReportSimplelistRequest() *OapiReportSimplelistRequest {
	return &OapiReportSimplelistRequest{}
}

type OapiReportSimplelistRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiReportSimplelistResponse
	Cursor          int64
	EndTime         int64
	Size            int64
	StartTime       int64
	TemplateName    string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiReportSimplelistRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiReportSimplelistRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiReportSimplelistRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiReportSimplelistRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiReportSimplelistRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *OapiReportSimplelistRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *OapiReportSimplelistRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiReportSimplelistRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiReportSimplelistRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *OapiReportSimplelistRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *OapiReportSimplelistRequest) SetTemplateName(templateName2 string) {
	this.TemplateName = templateName2
}
func (this *OapiReportSimplelistRequest) GetTemplateName() string {
	return this.TemplateName
}
func (this *OapiReportSimplelistRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiReportSimplelistRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiReportSimplelistRequest) GetApiMethodName() string {
	return "dingtalk.oapi.report.simplelist"
}
func (this *OapiReportSimplelistRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiReportSimplelistRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiReportSimplelistRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiReportSimplelistRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiReportSimplelistRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["end_time"] = this.EndTime
	txtParams["size"] = this.Size
	txtParams["start_time"] = this.StartTime
	txtParams["template_name"] = this.TemplateName
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiReportSimplelistRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *OapiReportSimplelistRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiReportSimplelistRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiReportSimplelistResponse struct {
	taobao.TaobaoResponse
	Result PageVo `json:"result,omitempty"`
}
