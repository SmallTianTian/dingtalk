package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewSmartworkBpmsProcessinstanceGetwithformRequest() *SmartworkBpmsProcessinstanceGetwithformRequest {
	return &SmartworkBpmsProcessinstanceGetwithformRequest{}
}

type SmartworkBpmsProcessinstanceGetwithformRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              SmartworkBpmsProcessinstanceGetwithformResponse
	ProcessInstanceId string
	TopHttpMethod     string
	TopResponseType   string
}

func (this *SmartworkBpmsProcessinstanceGetwithformRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *SmartworkBpmsProcessinstanceGetwithformRequest) SetProcessInstanceId(processInstanceId2 string) {
	this.ProcessInstanceId = processInstanceId2
}
func (this *SmartworkBpmsProcessinstanceGetwithformRequest) GetProcessInstanceId() string {
	return this.ProcessInstanceId
}
func (this *SmartworkBpmsProcessinstanceGetwithformRequest) GetApiMethodName() string {
	return "dingtalk.smartwork.bpms.processinstance.getwithform"
}
func (this *SmartworkBpmsProcessinstanceGetwithformRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *SmartworkBpmsProcessinstanceGetwithformRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *SmartworkBpmsProcessinstanceGetwithformRequest) GetTopApiCallType() string {
	return "top"
}
func (this *SmartworkBpmsProcessinstanceGetwithformRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *SmartworkBpmsProcessinstanceGetwithformRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *SmartworkBpmsProcessinstanceGetwithformRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["process_instance_id"] = this.ProcessInstanceId
	return txtParams
}
func (this *SmartworkBpmsProcessinstanceGetwithformRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProcessInstanceId, "processInstanceId"); err != nil {
		return err
	}
	return nil
}
func (this *SmartworkBpmsProcessinstanceGetwithformRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *SmartworkBpmsProcessinstanceGetwithformRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SmartworkBpmsProcessinstanceGetwithformResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type SubFormValueVo struct {
	DetailComponentType string `json:"detail_component_type,omitempty"`
	DetailFormBizAlias  string `json:"detail_form_biz_alias,omitempty"`
	DetailFormExtValue  string `json:"detail_form_ext_value,omitempty"`
	DetailFormId        string `json:"detail_form_id,omitempty"`
	DetailFormName      string `json:"detail_form_name,omitempty"`
	DetailFormValue     string `json:"detail_form_value,omitempty"`
}
type FormDetailVo struct {
	DetailFormValues []SubFormValueVo `json:"detail_form_values,omitempty"`
}
type FormValueVo struct {
	BizAlias      string         `json:"biz_alias,omitempty"`
	ComponentType string         `json:"component_type,omitempty"`
	Details       []FormDetailVo `json:"details,omitempty"`
	ExtValue      string         `json:"ext_value,omitempty"`
	FormId        string         `json:"form_id,omitempty"`
	Name          string         `json:"name,omitempty"`
	Value         string         `json:"value,omitempty"`
}
type NewProcessInstanceTopVo struct {
	ApproverUserids            []string             `json:"approver_userids,omitempty"`
	AttachedProcessInstanceIds []string             `json:"attached_process_instance_ids,omitempty"`
	BizAction                  string               `json:"biz_action,omitempty"`
	BusinessId                 string               `json:"business_id,omitempty"`
	CcUserids                  string               `json:"cc_userids,omitempty"`
	CreateTime                 time.Time            `json:"create_time,omitempty"`
	FinishTime                 time.Time            `json:"finish_time,omitempty"`
	FormValues                 []FormValueVo        `json:"form_values,omitempty"`
	OperationRecords           []OperationRecordsVo `json:"operation_records,omitempty"`
	OriginatorDeptId           string               `json:"originator_dept_id,omitempty"`
	OriginatorDeptName         string               `json:"originator_dept_name,omitempty"`
	OriginatorUserid           string               `json:"originator_userid,omitempty"`
	ProcessCode                string               `json:"process_code,omitempty"`
	Result                     string               `json:"result,omitempty"`
	Status                     string               `json:"status,omitempty"`
	Tasks                      []TaskTopVo          `json:"tasks,omitempty"`
	Title                      string               `json:"title,omitempty"`
}
