package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessTemplateSaveRequest() *OapiProcessTemplateSaveRequest {
	return &OapiProcessTemplateSaveRequest{}
}

type OapiProcessTemplateSaveRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessTemplateSaveResponse
	Font            string
	ProcessCode     string
	TopHttpMethod   string
	TopResponseType string
	Vm              string
}

func (this *OapiProcessTemplateSaveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessTemplateSaveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessTemplateSaveRequest) SetFont(font2 string) {
	this.Font = font2
}
func (this *OapiProcessTemplateSaveRequest) GetFont() string {
	return this.Font
}
func (this *OapiProcessTemplateSaveRequest) SetProcessCode(processCode2 string) {
	this.ProcessCode = processCode2
}
func (this *OapiProcessTemplateSaveRequest) GetProcessCode() string {
	return this.ProcessCode
}
func (this *OapiProcessTemplateSaveRequest) SetVm(vm2 string) {
	this.Vm = vm2
}
func (this *OapiProcessTemplateSaveRequest) GetVm() string {
	return this.Vm
}
func (this *OapiProcessTemplateSaveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.template.save"
}
func (this *OapiProcessTemplateSaveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessTemplateSaveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessTemplateSaveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessTemplateSaveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessTemplateSaveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["font"] = this.Font
	txtParams["process_code"] = this.ProcessCode
	txtParams["vm"] = this.Vm
	return txtParams
}
func (this *OapiProcessTemplateSaveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Font, "font"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessTemplateSaveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessTemplateSaveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessTemplateSaveResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
