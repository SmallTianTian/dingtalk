package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessinstanceGetRequest() *OapiProcessinstanceGetRequest {
	return &OapiProcessinstanceGetRequest{}
}

type OapiProcessinstanceGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp              OapiProcessinstanceGetResponse
	ProcessInstanceId string
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiProcessinstanceGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessinstanceGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessinstanceGetRequest) SetProcessInstanceId(processInstanceId2 string) {
	this.ProcessInstanceId = processInstanceId2
}
func (this *OapiProcessinstanceGetRequest) GetProcessInstanceId() string {
	return this.ProcessInstanceId
}
func (this *OapiProcessinstanceGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.processinstance.get"
}
func (this *OapiProcessinstanceGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessinstanceGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessinstanceGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessinstanceGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessinstanceGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["process_instance_id"] = this.ProcessInstanceId
	return txtParams
}
func (this *OapiProcessinstanceGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProcessInstanceId, "processInstanceId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessinstanceGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessinstanceGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessinstanceGetResponse struct {
	taobao.TaobaoResponse

	ProcessInstance ProcessInstanceTopVo `json:"process_instance,omitempty"`
}
type OperationRecordsVo struct {
	Attachments     []Attachment `json:"attachments,omitempty"`
	Date            time.Time    `json:"date,omitempty"`
	OperationResult string       `json:"operation_result,omitempty"`
	OperationType   string       `json:"operation_type,omitempty"`
	Remark          string       `json:"remark,omitempty"`
	Userid          string       `json:"userid,omitempty"`
}
type ProcessInstanceTopVo struct {
	ApproverUserids            []string               `json:"approver_userids,omitempty"`
	AttachedProcessInstanceIds []string               `json:"attached_process_instance_ids,omitempty"`
	BizAction                  string                 `json:"biz_action,omitempty"`
	BusinessId                 string                 `json:"business_id,omitempty"`
	CcUserids                  []string               `json:"cc_userids,omitempty"`
	CreateTime                 time.Time              `json:"create_time,omitempty"`
	FinishTime                 time.Time              `json:"finish_time,omitempty"`
	FormComponentValues        []FormComponentValueVo `json:"form_component_values,omitempty"`
	MainProcessInstanceId      string                 `json:"main_process_instance_id,omitempty"`
	OperationRecords           []OperationRecordsVo   `json:"operation_records,omitempty"`
	OriginatorDeptId           string                 `json:"originator_dept_id,omitempty"`
	OriginatorDeptName         string                 `json:"originator_dept_name,omitempty"`
	OriginatorUserid           string                 `json:"originator_userid,omitempty"`
	Result                     string                 `json:"result,omitempty"`
	Status                     string                 `json:"status,omitempty"`
	Tasks                      []TaskTopVo            `json:"tasks,omitempty"`
	Title                      string                 `json:"title,omitempty"`
}
