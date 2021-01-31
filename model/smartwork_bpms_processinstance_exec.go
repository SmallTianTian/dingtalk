package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewSmartworkBpmsProcessinstanceExecuteRequest() *SmartworkBpmsProcessinstanceExecuteRequest {
	return &SmartworkBpmsProcessinstanceExecuteRequest{}
}

type SmartworkBpmsProcessinstanceExecuteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              SmartworkBpmsProcessinstanceExecuteResponse
	ActionerUserid    string
	ProcessInstanceId string
	Remark            string
	Result            string
	TaskId            int64
	TopHttpMethod     string
	TopResponseType   string
}

func (this *SmartworkBpmsProcessinstanceExecuteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) SetActionerUserid(actionerUserid2 string) {
	this.ActionerUserid = actionerUserid2
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) GetActionerUserid() string {
	return this.ActionerUserid
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) SetProcessInstanceId(processInstanceId2 string) {
	this.ProcessInstanceId = processInstanceId2
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) GetProcessInstanceId() string {
	return this.ProcessInstanceId
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) SetRemark(remark2 string) {
	this.Remark = remark2
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) GetRemark() string {
	return this.Remark
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) SetResult(result2 string) {
	this.Result = result2
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) GetResult() string {
	return this.Result
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) SetTaskId(taskId2 int64) {
	this.TaskId = taskId2
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) GetTaskId() int64 {
	return this.TaskId
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) GetApiMethodName() string {
	return "dingtalk.smartwork.bpms.processinstance.execute"
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) GetTopApiCallType() string {
	return "top"
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["actioner_userid"] = this.ActionerUserid
	txtParams["process_instance_id"] = this.ProcessInstanceId
	txtParams["remark"] = this.Remark
	txtParams["result"] = this.Result
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ActionerUserid, "actionerUserid"); err != nil {
		return err
	}
	return nil
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *SmartworkBpmsProcessinstanceExecuteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SmartworkBpmsProcessinstanceExecuteResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
