package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiLiveCreateRequest() *OapiLiveCreateRequest {
	return &OapiLiveCreateRequest{}
}

type OapiLiveCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiLiveCreateResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiLiveCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiLiveCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiLiveCreateRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiLiveCreateRequest) GetRequest() string {
	return this.Request
}
func (this *OapiLiveCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.live.create"
}
func (this *OapiLiveCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiLiveCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiLiveCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiLiveCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiLiveCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiLiveCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiLiveCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiLiveCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CreateLiveReqModel struct {
	ApptBeginTime   time.Time `json:"appt_begin_time,omitempty"`
	ApptEndTime     time.Time `json:"appt_end_time,omitempty"`
	CoverUrl        string    `json:"cover_url,omitempty"`
	Intro           string    `json:"intro,omitempty"`
	LandScape       bool      `json:"land_scape,omitempty"`
	PreVideoPlayUrl string    `json:"pre_video_play_url,omitempty"`
	Shared          bool      `json:"shared,omitempty"`
	Title           string    `json:"title,omitempty"`
	UserNick        string    `json:"user_nick,omitempty"`
	Userid          string    `json:"userid,omitempty"`
}
type OapiLiveCreateResponse struct {
	taobao.TaobaoResponse
	Result CreateLiveRespModel `json:"result,omitempty"`
}
type LiveUrlExtModel struct {
	LiveUrlHigh     string `json:"live_url_high,omitempty"`
	LiveUrlLow      string `json:"live_url_low,omitempty"`
	LiveUrlNormal   string `json:"live_url_normal,omitempty"`
	LiveUrlUltraLow string `json:"live_url_ultra_low,omitempty"`
	LiveUrlVeryLow  string `json:"live_url_very_low,omitempty"`
}
type CreateLiveRespModel struct {
	AppointmentTime time.Time       `json:"appointment_time,omitempty"`
	InputStreamUrl  string          `json:"input_stream_url,omitempty"`
	LiveUrl         string          `json:"live_url,omitempty"`
	LiveUrlExt      LiveUrlExtModel `json:"live_url_ext,omitempty"`
	LiveUrlHls      string          `json:"live_url_hls,omitempty"`
	PlaybackUrl     string          `json:"playback_url,omitempty"`
	Uuid            string          `json:"uuid,omitempty"`
}
