package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiTdpTaskBasicCreateRequest() *OapiTdpTaskBasicCreateRequest {
	return &OapiTdpTaskBasicCreateRequest{}
}

type OapiTdpTaskBasicCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiTdpTaskBasicCreateResponse
	MicroappAgentId int64
	OperatorUserid  string
	Task            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiTdpTaskBasicCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiTdpTaskBasicCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiTdpTaskBasicCreateRequest) SetMicroappAgentId(microappAgentId2 int64) {
	this.MicroappAgentId = microappAgentId2
}
func (this *OapiTdpTaskBasicCreateRequest) GetMicroappAgentId() int64 {
	return this.MicroappAgentId
}
func (this *OapiTdpTaskBasicCreateRequest) SetOperatorUserid(operatorUserid2 string) {
	this.OperatorUserid = operatorUserid2
}
func (this *OapiTdpTaskBasicCreateRequest) GetOperatorUserid() string {
	return this.OperatorUserid
}
func (this *OapiTdpTaskBasicCreateRequest) SetTask(task2 string) {
	this.Task = task2
}
func (this *OapiTdpTaskBasicCreateRequest) GetTask() string {
	return this.Task
}
func (this *OapiTdpTaskBasicCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.tdp.task.basic.create"
}
func (this *OapiTdpTaskBasicCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiTdpTaskBasicCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiTdpTaskBasicCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiTdpTaskBasicCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiTdpTaskBasicCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["microapp_agent_id"] = this.MicroappAgentId
	txtParams["operator_userid"] = this.OperatorUserid
	txtParams["task"] = this.Task
	return txtParams
}
func (this *OapiTdpTaskBasicCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OperatorUserid, "operatorUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiTdpTaskBasicCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiTdpTaskBasicCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TaskExtension struct {
	CommentCount int64  `json:"comment_count,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
	Other        string `json:"other,omitempty"`
}
type TaskCreate struct {
	CreatorUserid    string        `json:"creator_userid,omitempty"`
	Description      string        `json:"description,omitempty"`
	ExecutorUserid   string        `json:"executor_userid,omitempty"`
	Extension        TaskExtension `json:"extension,omitempty"`
	FinishDate       time.Time     `json:"finish_date,omitempty"`
	GmtCreate        time.Time     `json:"gmt_create,omitempty"`
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
type OapiTdpTaskBasicCreateResponse struct {
	taobao.TaobaoResponse
	Result OrgTask `json:"result,omitempty"`
}
type OrgTask struct {
	BelongCorpId     string        `json:"belong_corp_id,omitempty"`
	BizTag           string        `json:"biz_tag,omitempty"`
	CreatorUserid    string        `json:"creator_userid,omitempty"`
	Description      string        `json:"description,omitempty"`
	ExecutorUserid   string        `json:"executor_userid,omitempty"`
	Extension        TaskExtension `json:"extension,omitempty"`
	FinishDate       time.Time     `json:"finish_date,omitempty"`
	GmtCreate        time.Time     `json:"gmt_create,omitempty"`
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
	TaskId           string        `json:"task_id,omitempty"`
	TaskTypeCategory string        `json:"task_type_category,omitempty"`
	TaskTypeId       int64         `json:"task_type_id,omitempty"`
	TrackerUserids   []string      `json:"tracker_userids,omitempty"`
}
