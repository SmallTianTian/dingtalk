package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessPrinterTemplateDeleteRequest() *OapiProcessPrinterTemplateDeleteRequest {
	return &OapiProcessPrinterTemplateDeleteRequest{}
}

type OapiProcessPrinterTemplateDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessPrinterTemplateDeleteResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessPrinterTemplateDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessPrinterTemplateDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessPrinterTemplateDeleteRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessPrinterTemplateDeleteRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessPrinterTemplateDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.printer.template.delete"
}
func (this *OapiProcessPrinterTemplateDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessPrinterTemplateDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessPrinterTemplateDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessPrinterTemplateDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessPrinterTemplateDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessPrinterTemplateDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessPrinterTemplateDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessPrinterTemplateDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PrintVmClearRequest struct {
	AgentId     int64  `json:"agent_id,omitempty"`
	ProcessCode string `json:"process_code,omitempty"`
}
type OapiProcessPrinterTemplateDeleteResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
