package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpDingTaskCreateRequest() *CorpDingTaskCreateRequest {
	return &CorpDingTaskCreateRequest{}
}

type CorpDingTaskCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpDingTaskCreateResponse
	TaskSendVO      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpDingTaskCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpDingTaskCreateRequest) SetTaskSendVO(taskSendVO2 string) {
	this.TaskSendVO = taskSendVO2
}
func (this *CorpDingTaskCreateRequest) GetTaskSendVO() string {
	return this.TaskSendVO
}
func (this *CorpDingTaskCreateRequest) GetApiMethodName() string {
	return "dingtalk.corp.ding.task.create"
}
func (this *CorpDingTaskCreateRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpDingTaskCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpDingTaskCreateRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpDingTaskCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpDingTaskCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpDingTaskCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["task_send_v_o"] = this.TaskSendVO
	return txtParams
}
func (this *CorpDingTaskCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpDingTaskCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpDingTaskCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type TaskSendVo struct {
	Content        string `json:"content,omitempty"`
	ContentType    int64  `json:"content_type,omitempty"`
	DeadLine       int64  `json:"dead_line,omitempty"`
	ReceiverUserid string `json:"receiver_userid,omitempty"`
	RemindTime     int64  `json:"remind_time,omitempty"`
	RemindType     int64  `json:"remind_type,omitempty"`
	SendUserid     string `json:"send_userid,omitempty"`
}
type CorpDingTaskCreateResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type TaskSendResult struct {
	DingId string `json:"ding_id,omitempty"`
}
