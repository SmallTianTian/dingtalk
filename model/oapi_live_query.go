package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiLiveQueryRequest() *OapiLiveQueryRequest {
	return &OapiLiveQueryRequest{}
}

type OapiLiveQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiLiveQueryResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiLiveQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiLiveQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiLiveQueryRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiLiveQueryRequest) GetRequest() string {
	return this.Request
}
func (this *OapiLiveQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.live.query"
}
func (this *OapiLiveQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiLiveQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiLiveQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiLiveQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiLiveQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiLiveQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiLiveQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiLiveQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GetDetailReqModel struct {
	Uuid string `json:"uuid,omitempty"`
}
type OapiLiveQueryResponse struct {
	taobao.TaobaoResponse
	Result GetDetailRespModel `json:"result,omitempty"`
}
type GetDetailRespModel struct {
	AppointmentTime time.Time       `json:"appointment_time,omitempty"`
	ApptBeginTime   time.Time       `json:"appt_begin_time,omitempty"`
	ApptEndTime     time.Time       `json:"appt_end_time,omitempty"`
	CoverUrl        string          `json:"cover_url,omitempty"`
	EndTime         time.Time       `json:"end_time,omitempty"`
	InputStreamUrl  string          `json:"input_stream_url,omitempty"`
	Intro           string          `json:"intro,omitempty"`
	LandScape       bool            `json:"land_scape,omitempty"`
	LiveUrl         string          `json:"live_url,omitempty"`
	LiveUrlExt      LiveUrlExtModel `json:"live_url_ext,omitempty"`
	LiveUrlHls      string          `json:"live_url_hls,omitempty"`
	PlaybackUrl     string          `json:"playback_url,omitempty"`
	PreVideoPlayUrl string          `json:"pre_video_play_url,omitempty"`
	Shared          bool            `json:"shared,omitempty"`
	StartTime       time.Time       `json:"start_time,omitempty"`
	Status          int64           `json:"status,omitempty"`
	TimeLength      int64           `json:"time_length,omitempty"`
	Title           string          `json:"title,omitempty"`
	TotalJoinCount  int64           `json:"total_join_count,omitempty"`
	TotalViewCount  int64           `json:"total_view_count,omitempty"`
	UserNick        string          `json:"user_nick,omitempty"`
	Userid          string          `json:"userid,omitempty"`
	Uuid            string          `json:"uuid,omitempty"`
}
