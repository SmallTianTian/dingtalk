package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessFormGetRequest() *OapiProcessFormGetRequest {
	return &OapiProcessFormGetRequest{}
}

type OapiProcessFormGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessFormGetResponse
	ProcessCode     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessFormGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessFormGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessFormGetRequest) SetProcessCode(processCode2 string) {
	this.ProcessCode = processCode2
}
func (this *OapiProcessFormGetRequest) GetProcessCode() string {
	return this.ProcessCode
}
func (this *OapiProcessFormGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.form.get"
}
func (this *OapiProcessFormGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessFormGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessFormGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessFormGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessFormGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["process_code"] = this.ProcessCode
	return txtParams
}
func (this *OapiProcessFormGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProcessCode, "processCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessFormGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessFormGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessFormGetResponse struct {
	taobao.TaobaoResponse
	Result  ProcessTopVo `json:"result,omitempty"`
	Success bool         `json:"success,omitempty"`
}
type FormComponentStatVo struct {
	Id    string `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Unit  string `json:"unit,omitempty"`
	Upper bool   `json:"upper,omitempty"`
}
type LinkageTargetVo struct {
	Behavior string `json:"behavior,omitempty"`
	FieldId  string `json:"field_id,omitempty"`
}
type FormComponentPropVo struct {
	AttendTypeLabel   string                `json:"attend_type_label,omitempty"`
	BehaviorLinkage   []BehaviorLinkageVo   `json:"behavior_linkage,omitempty"`
	BizAlias          string                `json:"biz_alias,omitempty"`
	BizType           string                `json:"biz_type,omitempty"`
	ChildFieldVisible string                `json:"child_field_visible,omitempty"`
	Disable           bool                  `json:"disable,omitempty"`
	Duration          bool                  `json:"duration,omitempty"`
	DurationLabel     string                `json:"duration_label,omitempty"`
	Format            string                `json:"format,omitempty"`
	Id                string                `json:"id,omitempty"`
	Invisible         bool                  `json:"invisible,omitempty"`
	Label             string                `json:"label,omitempty"`
	NotPrint          string                `json:"not_print,omitempty"`
	NotUpper          string                `json:"not_upper,omitempty"`
	Options           []string              `json:"options,omitempty"`
	Required          bool                  `json:"required,omitempty"`
	StatField         []FormComponentStatVo `json:"stat_field,omitempty"`
}
type FormComponent2Vo struct {
	ComponentName string              `json:"component_name,omitempty"`
	Props         FormComponentPropVo `json:"props,omitempty"`
}
type FormComponent1Vo struct {
	Children      []FormComponent2Vo  `json:"children,omitempty"`
	ComponentName string              `json:"component_name,omitempty"`
	Props         FormComponentPropVo `json:"props,omitempty"`
}
