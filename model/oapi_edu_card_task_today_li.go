package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduCardTaskTodayListRequest() *OapiEduCardTaskTodayListRequest {
	return &OapiEduCardTaskTodayListRequest{}
}

type OapiEduCardTaskTodayListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCardTaskTodayListResponse
	CardType        string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduCardTaskTodayListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCardTaskTodayListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCardTaskTodayListRequest) SetCardType(cardType2 string) {
	this.CardType = cardType2
}
func (this *OapiEduCardTaskTodayListRequest) GetCardType() string {
	return this.CardType
}
func (this *OapiEduCardTaskTodayListRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduCardTaskTodayListRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduCardTaskTodayListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.card.task.today.list"
}
func (this *OapiEduCardTaskTodayListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCardTaskTodayListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCardTaskTodayListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCardTaskTodayListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCardTaskTodayListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["card_type"] = this.CardType
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduCardTaskTodayListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduCardTaskTodayListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCardTaskTodayListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCardTaskTodayListResponse struct {
	taobao.TaobaoResponse
	Result  []CardTaskDTO `json:"result,omitempty"`
	Success bool          `json:"success,omitempty"`
}
type CardTaskDTO struct {
	CardId         int64     `json:"card_id,omitempty"`
	ClassName      string    `json:"class_name,omitempty"`
	Content        string    `json:"content,omitempty"`
	Date           time.Time `json:"date,omitempty"`
	IsFinishTask   string    `json:"is_finish_task,omitempty"`
	StudentName    string    `json:"student_name,omitempty"`
	Title          string    `json:"title,omitempty"`
	UserCardTaskId int64     `json:"user_card_task_id,omitempty"`
}
