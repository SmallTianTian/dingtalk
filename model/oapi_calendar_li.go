package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCalendarListRequest() *OapiCalendarListRequest {
	return &OapiCalendarListRequest{}
}

type OapiCalendarListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp             OapiCalendarListResponse
	CalendarFolderId string
	MaxResults       int64
	PageToken        string
	SingleEvents     bool
	TimeMax          string
	TimeMin          string
	TopHttpMethod    string
	TopResponseType  string
	UserId           string
}

func (this *OapiCalendarListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCalendarListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCalendarListRequest) SetCalendarFolderId(calendarFolderId2 string) {
	this.CalendarFolderId = calendarFolderId2
}
func (this *OapiCalendarListRequest) GetCalendarFolderId() string {
	return this.CalendarFolderId
}
func (this *OapiCalendarListRequest) SetMaxResults(maxResults2 int64) {
	this.MaxResults = maxResults2
}
func (this *OapiCalendarListRequest) GetMaxResults() int64 {
	return this.MaxResults
}
func (this *OapiCalendarListRequest) SetPageToken(pageToken2 string) {
	this.PageToken = pageToken2
}
func (this *OapiCalendarListRequest) GetPageToken() string {
	return this.PageToken
}
func (this *OapiCalendarListRequest) SetSingleEvents(singleEvents2 bool) {
	this.SingleEvents = singleEvents2
}
func (this *OapiCalendarListRequest) GetSingleEvents() bool {
	return this.SingleEvents
}
func (this *OapiCalendarListRequest) SetTimeMax(timeMax2 string) {
	this.TimeMax = timeMax2
}
func (this *OapiCalendarListRequest) GetTimeMax() string {
	return this.TimeMax
}
func (this *OapiCalendarListRequest) SetTimeMin(timeMin2 string) {
	this.TimeMin = timeMin2
}
func (this *OapiCalendarListRequest) GetTimeMin() string {
	return this.TimeMin
}
func (this *OapiCalendarListRequest) SetUserId(userId2 string) {
	this.UserId = userId2
}
func (this *OapiCalendarListRequest) GetUserId() string {
	return this.UserId
}
func (this *OapiCalendarListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.calendar.list"
}
func (this *OapiCalendarListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCalendarListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCalendarListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCalendarListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCalendarListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["calendar_folder_id"] = this.CalendarFolderId
	txtParams["max_results"] = this.MaxResults
	txtParams["page_token"] = this.PageToken
	txtParams["single_events"] = this.SingleEvents
	txtParams["time_max"] = this.TimeMax
	txtParams["time_min"] = this.TimeMin
	txtParams["user_id"] = this.UserId
	return txtParams
}
func (this *OapiCalendarListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.UserId, "userId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCalendarListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCalendarListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DateTime struct {
	DateTime string `json:"date_time,omitempty"`
	TimeZone string `json:"time_zone,omitempty"`
}
type OapiCalendarListResponse struct {
	taobao.TaobaoResponse
	Errcode int64                    `json:"errcode,omitempty"`
	Errmsg  string                   `json:"errmsg,omitempty"`
	Result  OpenCalendarListResponse `json:"result,omitempty"`
	Success bool                     `json:"success,omitempty"`
}
type Attendees struct {
	DisplayName    string `json:"display_name,omitempty"`
	Organizer      bool   `json:"organizer,omitempty"`
	ResponseStatus string `json:"response_status,omitempty"`
	Self           bool   `json:"self,omitempty"`
	Userid         string `json:"userid,omitempty"`
}
type User struct {
	DisplayName string `json:"display_name,omitempty"`
	Self        bool   `json:"self,omitempty"`
	Userid      string `json:"userid,omitempty"`
}
type Items struct {
	Attendees      []Attendees `json:"attendees,omitempty"`
	Created        DateTime    `json:"created,omitempty"`
	Description    string      `json:"description,omitempty"`
	End            DateTime    `json:"end,omitempty"`
	Id             string      `json:"id,omitempty"`
	Location       string      `json:"location,omitempty"`
	Organizer      User        `json:"organizer,omitempty"`
	Recurrence     []string    `json:"recurrence,omitempty"`
	RecurrenceId   string      `json:"recurrence_id,omitempty"`
	ResponseStatus string      `json:"response_status,omitempty"`
	Start          DateTime    `json:"start,omitempty"`
	Status         string      `json:"status,omitempty"`
	Summary        string      `json:"summary,omitempty"`
	UniqueId       string      `json:"unique_id,omitempty"`
	Updated        DateTime    `json:"updated,omitempty"`
}
type OpenCalendarListResponse struct {
	Items         []Items `json:"items,omitempty"`
	NextPageToken string  `json:"next_page_token,omitempty"`
	Summary       string  `json:"summary,omitempty"`
}
