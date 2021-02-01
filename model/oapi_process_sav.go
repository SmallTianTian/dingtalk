package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessSaveRequest() *OapiProcessSaveRequest {
	return &OapiProcessSaveRequest{}
}

type OapiProcessSaveRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiProcessSaveResponse
	SaveProcessRequest string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiProcessSaveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessSaveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessSaveRequest) SetSaveProcessRequest(saveProcessRequest2 string) {
	this.SaveProcessRequest = saveProcessRequest2
}
func (this *OapiProcessSaveRequest) GetSaveProcessRequest() string {
	return this.SaveProcessRequest
}
func (this *OapiProcessSaveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.save"
}
func (this *OapiProcessSaveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessSaveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessSaveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessSaveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessSaveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["saveProcessRequest"] = this.SaveProcessRequest
	return txtParams
}
func (this *OapiProcessSaveRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessSaveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessSaveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type FormComponentVo2 struct {
	ComponentName string              `json:"component_name,omitempty"`
	Props         FormComponentPropVo `json:"props,omitempty"`
}
type ProcessConfig struct {
	DisableDeleteProcess     bool   `json:"disable_delete_process,omitempty"`
	DisableFormEdit          bool   `json:"disable_form_edit,omitempty"`
	DisableHomepage          bool   `json:"disable_homepage,omitempty"`
	DisableResubmit          bool   `json:"disable_resubmit,omitempty"`
	DisableStopProcessButton bool   `json:"disable_stop_process_button,omitempty"`
	FakeTemplateEditUrl      string `json:"fake_template_edit_url,omitempty"`
	Hidden                   bool   `json:"hidden,omitempty"`
	TemplateEditUrl          string `json:"template_edit_url,omitempty"`
}
type SaveProcessRequest struct {
	Agentid                  int64             `json:"agentid,omitempty"`
	CreateInstanceMobileUrl  string            `json:"create_instance_mobile_url,omitempty"`
	CreateInstancePcUrl      string            `json:"create_instance_pc_url,omitempty"`
	Description              string            `json:"description,omitempty"`
	DirId                    string            `json:"dir_id,omitempty"`
	DisableFormEdit          bool              `json:"disable_form_edit,omitempty"`
	DisableStopProcessButton bool              `json:"disable_stop_process_button,omitempty"`
	FakeMode                 bool              `json:"fake_mode,omitempty"`
	FormComponentList        []FormComponentVo `json:"form_component_list,omitempty"`
	Hidden                   bool              `json:"hidden,omitempty"`
	Icon                     string            `json:"icon,omitempty"`
	Name                     string            `json:"name,omitempty"`
	OriginDirId              string            `json:"origin_dir_id,omitempty"`
	ProcessCode              string            `json:"process_code,omitempty"`
	ProcessConfig            ProcessConfig     `json:"process_config,omitempty"`
	TemplateEditUrl          string            `json:"template_edit_url,omitempty"`
}
type OapiProcessSaveResponse struct {
	taobao.TaobaoResponse
	Result ProcessTopVo `json:"result,omitempty"`
}
