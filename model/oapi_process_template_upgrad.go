package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessTemplateUpgradeRequest() *OapiProcessTemplateUpgradeRequest {
	return &OapiProcessTemplateUpgradeRequest{}
}

type OapiProcessTemplateUpgradeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiProcessTemplateUpgradeResponse
	DetailComponentId string
	FormComponentId   string
	ProcessCode       string
	TopHttpMethod     string
	TopResponseType   string
	Userid            string
}

func (this *OapiProcessTemplateUpgradeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessTemplateUpgradeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessTemplateUpgradeRequest) SetDetailComponentId(detailComponentId2 string) {
	this.DetailComponentId = detailComponentId2
}
func (this *OapiProcessTemplateUpgradeRequest) GetDetailComponentId() string {
	return this.DetailComponentId
}
func (this *OapiProcessTemplateUpgradeRequest) SetFormComponentId(formComponentId2 string) {
	this.FormComponentId = formComponentId2
}
func (this *OapiProcessTemplateUpgradeRequest) GetFormComponentId() string {
	return this.FormComponentId
}
func (this *OapiProcessTemplateUpgradeRequest) SetProcessCode(processCode2 string) {
	this.ProcessCode = processCode2
}
func (this *OapiProcessTemplateUpgradeRequest) GetProcessCode() string {
	return this.ProcessCode
}
func (this *OapiProcessTemplateUpgradeRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiProcessTemplateUpgradeRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiProcessTemplateUpgradeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.template.upgrade"
}
func (this *OapiProcessTemplateUpgradeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessTemplateUpgradeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessTemplateUpgradeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessTemplateUpgradeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessTemplateUpgradeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["detail_component_id"] = this.DetailComponentId
	txtParams["form_component_id"] = this.FormComponentId
	txtParams["process_code"] = this.ProcessCode
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiProcessTemplateUpgradeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.FormComponentId, "formComponentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessTemplateUpgradeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessTemplateUpgradeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessTemplateUpgradeResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
