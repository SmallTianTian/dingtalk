package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiReportStatisticsRequest() *OapiReportStatisticsRequest {
	return &OapiReportStatisticsRequest{}
}

type OapiReportStatisticsRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiReportStatisticsResponse
	ReportId        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiReportStatisticsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiReportStatisticsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiReportStatisticsRequest) SetReportId(reportId2 string) {
	this.ReportId = reportId2
}
func (this *OapiReportStatisticsRequest) GetReportId() string {
	return this.ReportId
}
func (this *OapiReportStatisticsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.report.statistics"
}
func (this *OapiReportStatisticsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiReportStatisticsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiReportStatisticsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiReportStatisticsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiReportStatisticsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["report_id"] = this.ReportId
	return txtParams
}
func (this *OapiReportStatisticsRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ReportId, "reportId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiReportStatisticsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiReportStatisticsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiReportStatisticsResponse struct {
	taobao.TaobaoResponse
	Result  ReportStatisticsVo `json:"result,omitempty"`
	Success bool               `json:"success,omitempty"`
}
type ReportStatisticsVo struct {
	CommentNum     int64 `json:"comment_num,omitempty"`
	CommentUserNum int64 `json:"comment_user_num,omitempty"`
	LikeNum        int64 `json:"like_num,omitempty"`
	ReadNum        int64 `json:"read_num,omitempty"`
}
