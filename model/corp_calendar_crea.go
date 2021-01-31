package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpCalendarCreateRequest() *CorpCalendarCreateRequest {
	return &CorpCalendarCreateRequest{}
}

type CorpCalendarCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpCalendarCreateResponse
	CreateVo        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpCalendarCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpCalendarCreateRequest) SetCreateVo(createVo2 string) {
	this.CreateVo = createVo2
}
func (this *CorpCalendarCreateRequest) GetCreateVo() string {
	return this.CreateVo
}
func (this *CorpCalendarCreateRequest) GetApiMethodName() string {
	return "dingtalk.corp.calendar.create"
}
func (this *CorpCalendarCreateRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpCalendarCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpCalendarCreateRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpCalendarCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpCalendarCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpCalendarCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["create_vo"] = this.CreateVo
	return txtParams
}
func (this *CorpCalendarCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpCalendarCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpCalendarCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenCalendarReminderVo struct {
	Minutes    int64  `json:"minutes,omitempty"`
	RemindType string `json:"remind_type,omitempty"`
}
type DatetimeVo struct {
	Timezone      string `json:"timezone,omitempty"`
	UnixTimestamp int64  `json:"unix_timestamp,omitempty"`
}
type OpenCalendarSourceVo struct {
	Title string `json:"title,omitempty"`
	Url   string `json:"url,omitempty"`
}
type OpenCalendarCreateVo struct {
	BizId           string                 `json:"biz_id,omitempty"`
	CalendarType    string                 `json:"calendar_type,omitempty"`
	CreatorUserid   string                 `json:"creator_userid,omitempty"`
	Description     string                 `json:"description,omitempty"`
	EndTime         DatetimeVo             `json:"end_time,omitempty"`
	Location        string                 `json:"location,omitempty"`
	ReceiverUserids []string               `json:"receiver_userids,omitempty"`
	Reminder        OpenCalendarReminderVo `json:"reminder,omitempty"`
	Source          OpenCalendarSourceVo   `json:"source,omitempty"`
	StartTime       DatetimeVo             `json:"start_time,omitempty"`
	Summary         string                 `json:"summary,omitempty"`
	Uuid            string                 `json:"uuid,omitempty"`
}
type CorpCalendarCreateResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type CorpCalendarCreateResult struct {
	DingtalkCalendarId string   `json:"dingtalk_calendar_id,omitempty"`
	InvalidUserids     []string `json:"invalid_userids,omitempty"`
}
