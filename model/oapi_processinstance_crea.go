package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessinstanceCreateRequest() *OapiProcessinstanceCreateRequest {
	return &OapiProcessinstanceCreateRequest{}
}

type OapiProcessinstanceCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                OapiProcessinstanceCreateResponse
	AgentId             int64
	Approvers           string
	ApproversV2         string
	CcList              string
	CcPosition          string
	DeptId              int64
	FormComponentValues string
	OriginatorUserId    string
	ProcessCode         string
	TopHttpMethod       string
	TopResponseType     string
}

func (this *OapiProcessinstanceCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessinstanceCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessinstanceCreateRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiProcessinstanceCreateRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiProcessinstanceCreateRequest) SetApprovers(approvers2 string) {
	this.Approvers = approvers2
}
func (this *OapiProcessinstanceCreateRequest) GetApprovers() string {
	return this.Approvers
}
func (this *OapiProcessinstanceCreateRequest) SetApproversV2(approversV22 string) {
	this.ApproversV2 = approversV22
}
func (this *OapiProcessinstanceCreateRequest) GetApproversV2() string {
	return this.ApproversV2
}
func (this *OapiProcessinstanceCreateRequest) SetCcList(ccList2 string) {
	this.CcList = ccList2
}
func (this *OapiProcessinstanceCreateRequest) GetCcList() string {
	return this.CcList
}
func (this *OapiProcessinstanceCreateRequest) SetCcPosition(ccPosition2 string) {
	this.CcPosition = ccPosition2
}
func (this *OapiProcessinstanceCreateRequest) GetCcPosition() string {
	return this.CcPosition
}
func (this *OapiProcessinstanceCreateRequest) SetDeptId(deptId2 int64) {
	this.DeptId = deptId2
}
func (this *OapiProcessinstanceCreateRequest) GetDeptId() int64 {
	return this.DeptId
}
func (this *OapiProcessinstanceCreateRequest) SetFormComponentValues(formComponentValues2 string) {
	this.FormComponentValues = formComponentValues2
}
func (this *OapiProcessinstanceCreateRequest) GetFormComponentValues() string {
	return this.FormComponentValues
}
func (this *OapiProcessinstanceCreateRequest) SetOriginatorUserId(originatorUserId2 string) {
	this.OriginatorUserId = originatorUserId2
}
func (this *OapiProcessinstanceCreateRequest) GetOriginatorUserId() string {
	return this.OriginatorUserId
}
func (this *OapiProcessinstanceCreateRequest) SetProcessCode(processCode2 string) {
	this.ProcessCode = processCode2
}
func (this *OapiProcessinstanceCreateRequest) GetProcessCode() string {
	return this.ProcessCode
}
func (this *OapiProcessinstanceCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.processinstance.create"
}
func (this *OapiProcessinstanceCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessinstanceCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessinstanceCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessinstanceCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessinstanceCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["approvers"] = this.Approvers
	txtParams["approvers_v2"] = this.ApproversV2
	txtParams["cc_list"] = this.CcList
	txtParams["cc_position"] = this.CcPosition
	txtParams["dept_id"] = this.DeptId
	txtParams["form_component_values"] = this.FormComponentValues
	txtParams["originator_user_id"] = this.OriginatorUserId
	txtParams["process_code"] = this.ProcessCode
	return txtParams
}
func (this *OapiProcessinstanceCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.Approvers, 20, "approvers"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessinstanceCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessinstanceCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type ProcessInstanceApproverVo struct {
	TaskActionType string   `json:"task_action_type,omitempty"`
	UserIds        []string `json:"user_ids,omitempty"`
}
type OapiProcessinstanceCreateResponse struct {
	taobao.TaobaoResponse

	ProcessInstanceId string `json:"process_instance_id,omitempty"`
}
