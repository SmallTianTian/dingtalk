package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCheckinRecordGetRequest() *OapiCheckinRecordGetRequest {
	return &OapiCheckinRecordGetRequest{}
}

type OapiCheckinRecordGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCheckinRecordGetResponse
	Cursor          int64
	EndTime         int64
	Size            int64
	StartTime       int64
	TopHttpMethod   string
	TopResponseType string
	UseridList      string
}

func (this *OapiCheckinRecordGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCheckinRecordGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCheckinRecordGetRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiCheckinRecordGetRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiCheckinRecordGetRequest) SetEndTime(endTime2 int64) {
	this.EndTime = endTime2
}
func (this *OapiCheckinRecordGetRequest) GetEndTime() int64 {
	return this.EndTime
}
func (this *OapiCheckinRecordGetRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiCheckinRecordGetRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiCheckinRecordGetRequest) SetStartTime(startTime2 int64) {
	this.StartTime = startTime2
}
func (this *OapiCheckinRecordGetRequest) GetStartTime() int64 {
	return this.StartTime
}
func (this *OapiCheckinRecordGetRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiCheckinRecordGetRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiCheckinRecordGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.checkin.record.get"
}
func (this *OapiCheckinRecordGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCheckinRecordGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCheckinRecordGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCheckinRecordGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCheckinRecordGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["end_time"] = this.EndTime
	txtParams["size"] = this.Size
	txtParams["start_time"] = this.StartTime
	txtParams["userid_list"] = this.UseridList
	return txtParams
}
func (this *OapiCheckinRecordGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cursor, "cursor"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCheckinRecordGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCheckinRecordGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCheckinRecordGetResponse struct {
	taobao.TaobaoResponse
	Result PageResult `json:"result,omitempty"`
}
type CheckinRecordVo struct {
	CheckinTime int64    `json:"checkin_time,omitempty"`
	DetailPlace string   `json:"detail_place,omitempty"`
	ImageList   []string `json:"image_list,omitempty"`
	Latitude    string   `json:"latitude,omitempty"`
	Longitude   string   `json:"longitude,omitempty"`
	Place       string   `json:"place,omitempty"`
	Remark      string   `json:"remark,omitempty"`
	Userid      string   `json:"userid,omitempty"`
	VisitUser   string   `json:"visit_user,omitempty"`
}
