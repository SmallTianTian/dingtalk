package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewCorpReportListRequest() *CorpReportListRequest {
	return &CorpReportListRequest{}
}

type CorpReportListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpReportListResponse
	Cursor          int64
	EndTime         int64
	Size            int64
	StartTime       int64
	TemplateName    string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *CorpReportListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpReportListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *CorpReportListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *CorpReportListRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *CorpReportListRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *CorpReportListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *CorpReportListRequest) GetSize() int64 {
	return this.Size
}
func (this *CorpReportListRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *CorpReportListRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *CorpReportListRequest) SetTemplateName(templateName2 string) {
	this.TemplateName = templateName2
}
func (this *CorpReportListRequest) GetTemplateName() string {
	return this.TemplateName
}
func (this *CorpReportListRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *CorpReportListRequest) GetUserid() string {
	return this.Userid
}
func (this *CorpReportListRequest) GetApiMethodName() string {
	return "dingtalk.corp.report.list"
}
func (this *CorpReportListRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpReportListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpReportListRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpReportListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpReportListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpReportListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["end_time"] = this.EndTime
	txtParams["size"] = this.Size
	txtParams["start_time"] = this.StartTime
	txtParams["template_name"] = this.TemplateName
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *CorpReportListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *CorpReportListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpReportListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpReportListResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type JsonObject struct {
	Key   string `json:"key,omitempty"`
	Sort  string `json:"sort,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
type ReportOapiVo struct {
	Contents     []JsonObject `json:"contents,omitempty"`
	CreateTime   int64        `json:"create_time,omitempty"`
	CreatorId    string       `json:"creator_id,omitempty"`
	CreatorName  string       `json:"creator_name,omitempty"`
	DeptName     string       `json:"dept_name,omitempty"`
	Remark       string       `json:"remark,omitempty"`
	ReportId     string       `json:"report_id,omitempty"`
	TemplateName string       `json:"template_name,omitempty"`
}
