package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiReportTemplateListbyuseridRequest() *OapiReportTemplateListbyuseridRequest {
	return &OapiReportTemplateListbyuseridRequest{}
}

type OapiReportTemplateListbyuseridRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiReportTemplateListbyuseridResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiReportTemplateListbyuseridRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiReportTemplateListbyuseridRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiReportTemplateListbyuseridRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiReportTemplateListbyuseridRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiReportTemplateListbyuseridRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiReportTemplateListbyuseridRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiReportTemplateListbyuseridRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiReportTemplateListbyuseridRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiReportTemplateListbyuseridRequest) GetApiMethodName() string {
	return "dingtalk.oapi.report.template.listbyuserid"
}
func (this *OapiReportTemplateListbyuseridRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiReportTemplateListbyuseridRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiReportTemplateListbyuseridRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiReportTemplateListbyuseridRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiReportTemplateListbyuseridRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiReportTemplateListbyuseridRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiReportTemplateListbyuseridRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiReportTemplateListbyuseridRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiReportTemplateListbyuseridResponse struct {
	taobao.TaobaoResponse
	Result HomePageReportTemplateVo `json:"result,omitempty"`
}
type ReportTemplateTopVo struct {
	IconUrl    string `json:"icon_url,omitempty"`
	Name       string `json:"name,omitempty"`
	ReportCode string `json:"report_code,omitempty"`
	Url        string `json:"url,omitempty"`
}
type HomePageReportTemplateVo struct {
	NextCursor   int64                 `json:"next_cursor,omitempty"`
	TemplateList []ReportTemplateTopVo `json:"template_list,omitempty"`
}
