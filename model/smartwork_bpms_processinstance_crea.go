package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewSmartworkBpmsProcessinstanceCreateRequest() *SmartworkBpmsProcessinstanceCreateRequest {
	return &SmartworkBpmsProcessinstanceCreateRequest{}
}

type SmartworkBpmsProcessinstanceCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                SmartworkBpmsProcessinstanceCreateResponse
	AgentId             int64
	Approvers           string
	CcList              string
	CcPosition          string
	DeptId              int64
	FormComponentValues string
	OriginatorUserId    string
	ProcessCode         string
	TopHttpMethod       string
	TopResponseType     string
}

func (this *SmartworkBpmsProcessinstanceCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) SetApprovers(approvers2 string) {
	this.Approvers = approvers2
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) GetApprovers() string {
	return this.Approvers
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) SetCcList(ccList2 string) {
	this.CcList = ccList2
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) GetCcList() string {
	return this.CcList
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) SetCcPosition(ccPosition2 string) {
	this.CcPosition = ccPosition2
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) GetCcPosition() string {
	return this.CcPosition
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) SetFormComponentValues(formComponentValues2 string) {
	this.FormComponentValues = formComponentValues2
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) GetFormComponentValues() string {
	return this.FormComponentValues
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) SetOriginatorUserId(originatorUserId2 string) {
	this.OriginatorUserId = originatorUserId2
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) GetOriginatorUserId() string {
	return this.OriginatorUserId
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) SetProcessCode(processCode2 string) {
	this.ProcessCode = processCode2
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) GetProcessCode() string {
	return this.ProcessCode
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) GetApiMethodName() string {
	return "dingtalk.smartwork.bpms.processinstance.create"
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) GetTopApiCallType() string {
	return "top"
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["approvers"] = this.Approvers
	txtParams["cc_list"] = this.CcList
	txtParams["cc_position"] = this.CcPosition
	txtParams["dept_id"] = this.DeptId
	txtParams["form_component_values"] = this.FormComponentValues
	txtParams["originator_user_id"] = this.OriginatorUserId
	txtParams["process_code"] = this.ProcessCode
	return txtParams
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Approvers, "approvers"); err != nil {
		return err
	}
	return nil
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *SmartworkBpmsProcessinstanceCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SmartworkBpmsProcessinstanceCreateResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
