package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessinstanceVariableUpdateRequest() *OapiProcessinstanceVariableUpdateRequest {
	return &OapiProcessinstanceVariableUpdateRequest{}
}

type OapiProcessinstanceVariableUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiProcessinstanceVariableUpdateResponse
	ProcessInstanceId string
	Remark            string
	TopHttpMethod     string
	TopResponseType   string
	Variables         string
}

func (this *OapiProcessinstanceVariableUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessinstanceVariableUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessinstanceVariableUpdateRequest) SetProcessInstanceId(processInstanceId2 string) {
	this.ProcessInstanceId = processInstanceId2
}
func (this *OapiProcessinstanceVariableUpdateRequest) GetProcessInstanceId() string {
	return this.ProcessInstanceId
}
func (this *OapiProcessinstanceVariableUpdateRequest) SetRemark(remark2 string) {
	this.Remark = remark2
}
func (this *OapiProcessinstanceVariableUpdateRequest) GetRemark() string {
	return this.Remark
}
func (this *OapiProcessinstanceVariableUpdateRequest) SetVariables(variables2 string) {
	this.Variables = variables2
}
func (this *OapiProcessinstanceVariableUpdateRequest) GetVariables() string {
	return this.Variables
}
func (this *OapiProcessinstanceVariableUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.processinstance.variable.update"
}
func (this *OapiProcessinstanceVariableUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessinstanceVariableUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessinstanceVariableUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessinstanceVariableUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessinstanceVariableUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["process_instance_id"] = this.ProcessInstanceId
	txtParams["remark"] = this.Remark
	txtParams["variables"] = this.Variables
	return txtParams
}
func (this *OapiProcessinstanceVariableUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProcessInstanceId, "processInstanceId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessinstanceVariableUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessinstanceVariableUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type BizValueVo struct {
	ExtValue string `json:"ext_value,omitempty"`
	Id       string `json:"id,omitempty"`
	Value    string `json:"value,omitempty"`
}
type OapiProcessinstanceVariableUpdateResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
