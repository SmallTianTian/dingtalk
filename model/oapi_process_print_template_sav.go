package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessPrintTemplateSaveRequest() *OapiProcessPrintTemplateSaveRequest {
	return &OapiProcessPrintTemplateSaveRequest{}
}

type OapiProcessPrintTemplateSaveRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiProcessPrintTemplateSaveResponse
	Font               string
	OpenCustomizePrint bool
	ProcessCode        string
	TopHttpMethod      string
	TopResponseType    string
	Vm                 string
}

func (this *OapiProcessPrintTemplateSaveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessPrintTemplateSaveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessPrintTemplateSaveRequest) SetFont(font2 string) {
	this.Font = font2
}
func (this *OapiProcessPrintTemplateSaveRequest) GetFont() string {
	return this.Font
}
func (this *OapiProcessPrintTemplateSaveRequest) SetOpenCustomizePrint(openCustomizePrint2 bool) {
	this.OpenCustomizePrint = openCustomizePrint2
}
func (this *OapiProcessPrintTemplateSaveRequest) GetOpenCustomizePrint() bool {
	return this.OpenCustomizePrint
}
func (this *OapiProcessPrintTemplateSaveRequest) SetProcessCode(processCode2 string) {
	this.ProcessCode = processCode2
}
func (this *OapiProcessPrintTemplateSaveRequest) GetProcessCode() string {
	return this.ProcessCode
}
func (this *OapiProcessPrintTemplateSaveRequest) SetVm(vm2 string) {
	this.Vm = vm2
}
func (this *OapiProcessPrintTemplateSaveRequest) GetVm() string {
	return this.Vm
}
func (this *OapiProcessPrintTemplateSaveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.print.template.save"
}
func (this *OapiProcessPrintTemplateSaveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessPrintTemplateSaveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessPrintTemplateSaveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessPrintTemplateSaveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessPrintTemplateSaveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["font"] = this.Font
	txtParams["open_customize_print"] = this.OpenCustomizePrint
	txtParams["process_code"] = this.ProcessCode
	txtParams["vm"] = this.Vm
	return txtParams
}
func (this *OapiProcessPrintTemplateSaveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Font, "font"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessPrintTemplateSaveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessPrintTemplateSaveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessPrintTemplateSaveResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
