package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMcsConferenceCreateRequest() *OapiMcsConferenceCreateRequest {
	return &OapiMcsConferenceCreateRequest{}
}

type OapiMcsConferenceCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp             OapiMcsConferenceCreateResponse
	BizKey           string
	IsPushRecord     bool
	PreferenceRegion string
	RoomValidTime    int64
	Title            string
	TopHttpMethod    string
	TopResponseType  string
}

func (this *OapiMcsConferenceCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMcsConferenceCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMcsConferenceCreateRequest) SetBizKey(bizKey2 string) {
	this.BizKey = bizKey2
}
func (this *OapiMcsConferenceCreateRequest) GetBizKey() string {
	return this.BizKey
}
func (this *OapiMcsConferenceCreateRequest) SetIsPushRecord(isPushRecord2 bool) {
	this.IsPushRecord = isPushRecord2
}
func (this *OapiMcsConferenceCreateRequest) GetIsPushRecord() bool {
	return this.IsPushRecord
}
func (this *OapiMcsConferenceCreateRequest) SetPreferenceRegion(preferenceRegion2 string) {
	this.PreferenceRegion = preferenceRegion2
}
func (this *OapiMcsConferenceCreateRequest) GetPreferenceRegion() string {
	return this.PreferenceRegion
}
func (this *OapiMcsConferenceCreateRequest) SetRoomValidTime(roomValidTime2 int64) {
	this.RoomValidTime = roomValidTime2
}
func (this *OapiMcsConferenceCreateRequest) GetRoomValidTime() int64 {
	return this.RoomValidTime
}
func (this *OapiMcsConferenceCreateRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiMcsConferenceCreateRequest) GetTitle() string {
	return this.Title
}
func (this *OapiMcsConferenceCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.mcs.conference.create"
}
func (this *OapiMcsConferenceCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMcsConferenceCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMcsConferenceCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMcsConferenceCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMcsConferenceCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_key"] = this.BizKey
	txtParams["is_push_record"] = this.IsPushRecord
	txtParams["preference_region"] = this.PreferenceRegion
	txtParams["room_valid_time"] = this.RoomValidTime
	txtParams["title"] = this.Title
	return txtParams
}
func (this *OapiMcsConferenceCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizKey, "bizKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMcsConferenceCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMcsConferenceCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMcsConferenceCreateResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
