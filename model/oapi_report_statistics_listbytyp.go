package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiReportStatisticsListbytypeRequest() *OapiReportStatisticsListbytypeRequest {
	return &OapiReportStatisticsListbytypeRequest{}
}

type OapiReportStatisticsListbytypeRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiReportStatisticsListbytypeResponse
	Offset          int64
	ReportId        string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
	Type            int64
}

func (this *OapiReportStatisticsListbytypeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiReportStatisticsListbytypeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiReportStatisticsListbytypeRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiReportStatisticsListbytypeRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiReportStatisticsListbytypeRequest) SetReportId(reportId2 string) {
	this.ReportId = reportId2
}
func (this *OapiReportStatisticsListbytypeRequest) GetReportId() string {
	return this.ReportId
}
func (this *OapiReportStatisticsListbytypeRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiReportStatisticsListbytypeRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiReportStatisticsListbytypeRequest) SetType(type2 int64) {
	this.Type = type2
}
func (this *OapiReportStatisticsListbytypeRequest) GetType() int64 {
	return this.Type
}
func (this *OapiReportStatisticsListbytypeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.report.statistics.listbytype"
}
func (this *OapiReportStatisticsListbytypeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiReportStatisticsListbytypeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiReportStatisticsListbytypeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiReportStatisticsListbytypeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiReportStatisticsListbytypeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["report_id"] = this.ReportId
	txtParams["size"] = this.Size
	txtParams["type"] = this.Type
	return txtParams
}
func (this *OapiReportStatisticsListbytypeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ReportId, "reportId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiReportStatisticsListbytypeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiReportStatisticsListbytypeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiReportStatisticsListbytypeResponse struct {
	taobao.TaobaoResponse
	Result  ReportPageVo `json:"result,omitempty"`
	Success bool         `json:"success,omitempty"`
}
