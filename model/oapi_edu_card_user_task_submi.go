package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduCardUserTaskSubmitRequest() *OapiEduCardUserTaskSubmitRequest {
	return &OapiEduCardUserTaskSubmitRequest{}
}

type OapiEduCardUserTaskSubmitRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCardUserTaskSubmitResponse
	Taskparam       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduCardUserTaskSubmitRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCardUserTaskSubmitRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCardUserTaskSubmitRequest) SetTaskparam(taskparam2 string) {
	this.Taskparam = taskparam2
}
func (this *OapiEduCardUserTaskSubmitRequest) GetTaskparam() string {
	return this.Taskparam
}
func (this *OapiEduCardUserTaskSubmitRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.card.user.task.submit"
}
func (this *OapiEduCardUserTaskSubmitRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCardUserTaskSubmitRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCardUserTaskSubmitRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCardUserTaskSubmitRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCardUserTaskSubmitRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["taskparam"] = this.Taskparam
	return txtParams
}
func (this *OapiEduCardUserTaskSubmitRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduCardUserTaskSubmitRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCardUserTaskSubmitRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenCardTaskSubmitParam struct {
	CardBizId         string `json:"card_biz_id,omitempty"`
	CardBizcode       string `json:"card_bizcode,omitempty"`
	CardId            int64  `json:"card_id,omitempty"`
	CardTaskCode      string `json:"card_task_code,omitempty"`
	CardTaskId        int64  `json:"card_task_id,omitempty"`
	Content           string `json:"content,omitempty"`
	DetailUrl         string `json:"detail_url,omitempty"`
	EditUrl           string `json:"edit_url,omitempty"`
	Medias            string `json:"medias,omitempty"`
	MeteringNumber    string `json:"metering_number,omitempty"`
	ReissueCard       bool   `json:"reissue_card,omitempty"`
	ResultEvaluation  string `json:"result_evaluation,omitempty"`
	SourceType        string `json:"source_type,omitempty"`
	UnitOfMeasurement string `json:"unit_of_measurement,omitempty"`
	Userid            string `json:"userid,omitempty"`
}
type OapiEduCardUserTaskSubmitResponse struct {
	taobao.TaobaoResponse
	Errcode int64                  `json:"errcode,omitempty"`
	Errmsg  string                 `json:"errmsg,omitempty"`
	Result  CardTaskSubmitResponse `json:"result,omitempty"`
	Success bool                   `json:"success,omitempty"`
}
type CardTaskSubmitResponse struct {
	Id int64 `json:"id,omitempty"`
}
