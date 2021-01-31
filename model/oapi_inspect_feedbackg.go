package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiInspectFeedbackGetRequest() *OapiInspectFeedbackGetRequest {
	return &OapiInspectFeedbackGetRequest{}
}

type OapiInspectFeedbackGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiInspectFeedbackGetResponse
	FormId          string
	TaskId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiInspectFeedbackGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiInspectFeedbackGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiInspectFeedbackGetRequest) SetFormId(formId2 string) {
	this.FormId = formId2
}
func (this *OapiInspectFeedbackGetRequest) GetFormId() string {
	return this.FormId
}
func (this *OapiInspectFeedbackGetRequest) SetTaskId(taskId2 string) {
	this.TaskId = taskId2
}
func (this *OapiInspectFeedbackGetRequest) GetTaskId() string {
	return this.TaskId
}
func (this *OapiInspectFeedbackGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.inspect.feedback.get"
}
func (this *OapiInspectFeedbackGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiInspectFeedbackGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiInspectFeedbackGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiInspectFeedbackGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiInspectFeedbackGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["form_id"] = this.FormId
	txtParams["task_id"] = this.TaskId
	return txtParams
}
func (this *OapiInspectFeedbackGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.TaskId, "taskId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiInspectFeedbackGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiInspectFeedbackGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiInspectFeedbackGetResponse struct {
	taobao.TaobaoResponse
	Errcode      int64                `json:"errcode,omitempty"`
	Errmsg       string               `json:"errmsg,omitempty"`
	FeedbackForm TopInspectFeedbackVo `json:"feedback_form,omitempty"`
}
type TopInspectFeedbackFormItemVO struct {
	BizAlias string `json:"biz_alias,omitempty"`
	Id       string `json:"id,omitempty"`
	Label    string `json:"label,omitempty"`
	Type     string `json:"type,omitempty"`
	Value    string `json:"value,omitempty"`
}
type TopInspectFeedbackVo struct {
	Content    []TopInspectFeedbackFormItemVO `json:"content,omitempty"`
	FormInstId string                         `json:"form_inst_id,omitempty"`
	Title      string                         `json:"title,omitempty"`
}
