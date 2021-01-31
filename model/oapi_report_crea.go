package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiReportCreateRequest() *OapiReportCreateRequest {
	return &OapiReportCreateRequest{}
}

type OapiReportCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiReportCreateResponse
	CreateReportParam string
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiReportCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiReportCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiReportCreateRequest) SetCreateReportParam(createReportParam2 string) {
	this.CreateReportParam = createReportParam2
}
func (this *OapiReportCreateRequest) GetCreateReportParam() string {
	return this.CreateReportParam
}
func (this *OapiReportCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.report.create"
}
func (this *OapiReportCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiReportCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiReportCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiReportCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiReportCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["create_report_param"] = this.CreateReportParam
	return txtParams
}
func (this *OapiReportCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiReportCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiReportCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiReportContentVo struct {
	Content     string `json:"content,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	Key         string `json:"key,omitempty"`
	Sort        int64  `json:"sort,omitempty"`
	Type        int64  `json:"type,omitempty"`
}
type OapiCreateReportParam struct {
	Contents   []OapiReportContentVo `json:"contents,omitempty"`
	DdFrom     string                `json:"dd_from,omitempty"`
	TemplateId string                `json:"template_id,omitempty"`
	ToChat     bool                  `json:"to_chat,omitempty"`
	ToCids     []string              `json:"to_cids,omitempty"`
	ToUserids  []string              `json:"to_userids,omitempty"`
	Userid     string                `json:"userid,omitempty"`
}
type OapiReportCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
}
