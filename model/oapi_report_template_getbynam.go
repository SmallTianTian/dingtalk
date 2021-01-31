package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiReportTemplateGetbynameRequest() *OapiReportTemplateGetbynameRequest {
	return &OapiReportTemplateGetbynameRequest{}
}

type OapiReportTemplateGetbynameRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiReportTemplateGetbynameResponse
	TemplateName    string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiReportTemplateGetbynameRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiReportTemplateGetbynameRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiReportTemplateGetbynameRequest) SetTemplateName(templateName2 string) {
	this.TemplateName = templateName2
}
func (this *OapiReportTemplateGetbynameRequest) GetTemplateName() string {
	return this.TemplateName
}
func (this *OapiReportTemplateGetbynameRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiReportTemplateGetbynameRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiReportTemplateGetbynameRequest) GetApiMethodName() string {
	return "dingtalk.oapi.report.template.getbyname"
}
func (this *OapiReportTemplateGetbynameRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiReportTemplateGetbynameRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiReportTemplateGetbynameRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiReportTemplateGetbynameRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiReportTemplateGetbynameRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["template_name"] = this.TemplateName
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiReportTemplateGetbynameRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.TemplateName, "templateName"); err != nil {
		return err
	}
	return nil
}
func (this *OapiReportTemplateGetbynameRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiReportTemplateGetbynameRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiReportTemplateGetbynameResponse struct {
	taobao.TaobaoResponse
	Result ReportTemplateResponseVo `json:"result,omitempty"`
}
type DefaultReceivers struct {
	UserName string `json:"user_name,omitempty"`
	Userid   string `json:"userid,omitempty"`
}
type BaseConversationVo struct {
	ConversationId string `json:"conversation_id,omitempty"`
	Title          string `json:"title,omitempty"`
}
type ReportTemplateResponseVo struct {
	DefaultReceivedConvs []BaseConversationVo `json:"default_received_convs,omitempty"`
	DefaultReceivers     []DefaultReceivers   `json:"default_receivers,omitempty"`
	Fields               []Fields             `json:"fields,omitempty"`
	Id                   string               `json:"id,omitempty"`
	Name                 string               `json:"name,omitempty"`
	UserName             string               `json:"user_name,omitempty"`
	Userid               string               `json:"userid,omitempty"`
}
