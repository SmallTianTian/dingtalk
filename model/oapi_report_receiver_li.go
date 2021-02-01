package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiReportReceiverListRequest() *OapiReportReceiverListRequest {
	return &OapiReportReceiverListRequest{}
}

type OapiReportReceiverListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiReportReceiverListResponse
	Offset          int64
	ReportId        string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiReportReceiverListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiReportReceiverListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiReportReceiverListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiReportReceiverListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiReportReceiverListRequest) SetReportId(reportId2 string) {
	this.ReportId = reportId2
}
func (this *OapiReportReceiverListRequest) GetReportId() string {
	return this.ReportId
}
func (this *OapiReportReceiverListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiReportReceiverListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiReportReceiverListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.report.receiver.list"
}
func (this *OapiReportReceiverListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiReportReceiverListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiReportReceiverListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiReportReceiverListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiReportReceiverListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["report_id"] = this.ReportId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiReportReceiverListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ReportId, "reportId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiReportReceiverListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiReportReceiverListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiReportReceiverListResponse struct {
	taobao.TaobaoResponse
	Result  ReportPageVo `json:"result,omitempty"`
	Success bool         `json:"success,omitempty"`
}
