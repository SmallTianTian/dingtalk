package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiTdpTaskBasicUpdateRequest() *OapiTdpTaskBasicUpdateRequest {
	return &OapiTdpTaskBasicUpdateRequest{}
}

type OapiTdpTaskBasicUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpTaskBasicUpdateResponse
	MicroappAgentId int64
	OperatorUserid  string
	Task            string
	TaskId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiTdpTaskBasicUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpTaskBasicUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpTaskBasicUpdateRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpTaskBasicUpdateRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpTaskBasicUpdateRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpTaskBasicUpdateRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpTaskBasicUpdateRequest) SetTask(task2 string) {
	this.Task = task2
}
func (this *OapiTdpTaskBasicUpdateRequest) GetTask() string {
	return this.Task
}
func (this *OapiTdpTaskBasicUpdateRequest) SetTaskId(taskId2 string) {
	this.TaskId = taskId2
}
func (this *OapiTdpTaskBasicUpdateRequest) GetTaskId() string {
	return this.TaskId
}
func (this *OapiTdpTaskBasicUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.task.basic.update"
}
func (this *OapiTdpTaskBasicUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpTaskBasicUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpTaskBasicUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpTaskBasicUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpTaskBasicUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["task"] = this.Task
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *OapiTdpTaskBasicUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.TaskId, "taskId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpTaskBasicUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpTaskBasicUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TaskUpdate struct {
	Description      string        `json:"description,omitempty"`
	ExecutorUserid   string        `json:"executor_userid,omitempty"`
	Extension        TaskExtension `json:"extension,omitempty"`
	FinishDate       time.Time     `json:"finish_date,omitempty"`
	GmtModified      time.Time     `json:"gmt_modified,omitempty"`
	IsArchived       bool          `json:"is_archived,omitempty"`
	IsRecycled       bool          `json:"is_recycled,omitempty"`
	ModifierUserid   string        `json:"modifier_userid,omitempty"`
	ParentId         string        `json:"parent_id,omitempty"`
	PlanFinishDate   time.Time     `json:"plan_finish_date,omitempty"`
	PlanStartDate    time.Time     `json:"plan_start_date,omitempty"`
	Priority         int64         `json:"priority,omitempty"`
	ProjectId        string        `json:"project_id,omitempty"`
	Source           string        `json:"source,omitempty"`
	SourceId         string        `json:"source_id,omitempty"`
	StartDate        time.Time     `json:"start_date,omitempty"`
	StatusId         int64         `json:"status_id,omitempty"`
	StatusStage      int64         `json:"status_stage,omitempty"`
	Subject          string        `json:"subject,omitempty"`
	TaskTypeCategory string        `json:"task_type_category,omitempty"`
	TaskTypeId       int64         `json:"task_type_id,omitempty"`
	TrackerUserids   []string      `json:"tracker_userids,omitempty"`
}
type OapiTdpTaskBasicUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
