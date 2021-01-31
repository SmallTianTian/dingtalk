package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiReportSavecontentRequest() *OapiReportSavecontentRequest {
	return &OapiReportSavecontentRequest{}
}

type OapiReportSavecontentRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiReportSavecontentResponse
	CreateReportParam string
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiReportSavecontentRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiReportSavecontentRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiReportSavecontentRequest) SetCreateReportParam(createReportParam2 string) {
	this.CreateReportParam = createReportParam2
}
func (this *OapiReportSavecontentRequest) GetCreateReportParam() string {
	return this.CreateReportParam
}
func (this *OapiReportSavecontentRequest) GetApiMethodName() string {
	return "dingtalk.oapi.report.savecontent"
}
func (this *OapiReportSavecontentRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiReportSavecontentRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiReportSavecontentRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiReportSavecontentRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiReportSavecontentRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["create_report_param"] = this.CreateReportParam
	return txtParams
}
func (this *OapiReportSavecontentRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiReportSavecontentRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiReportSavecontentRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiReportSavecontentResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
}
