package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiReportCommentListRequest() *OapiReportCommentListRequest {
	return &OapiReportCommentListRequest{}
}

type OapiReportCommentListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiReportCommentListResponse
	Offset          int64
	ReportId        string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiReportCommentListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiReportCommentListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiReportCommentListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiReportCommentListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiReportCommentListRequest) SetReportId(reportId2 string) {
	this.ReportId = reportId2
}
func (this *OapiReportCommentListRequest) GetReportId() string {
	return this.ReportId
}
func (this *OapiReportCommentListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiReportCommentListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiReportCommentListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.report.comment.list"
}
func (this *OapiReportCommentListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiReportCommentListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiReportCommentListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiReportCommentListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiReportCommentListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["report_id"] = this.ReportId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiReportCommentListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ReportId, "reportId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiReportCommentListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiReportCommentListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiReportCommentListResponse struct {
	taobao.TaobaoResponse
	Result  ReportPageVo `json:"result,omitempty"`
	Success bool         `json:"success,omitempty"`
}
type ReportCommentVo struct {
	Content    string    `json:"content,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	Userid     string    `json:"userid,omitempty"`
}
type ReportPageVo struct {
	Comments   []ReportCommentVo `json:"comments,omitempty"`
	HasMore    bool              `json:"has_more,omitempty"`
	NextCursor int64             `json:"next_cursor,omitempty"`
}
