package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiLivePlaybackRequest() *OapiLivePlaybackRequest {
	return &OapiLivePlaybackRequest{}
}

type OapiLivePlaybackRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiLivePlaybackResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiLivePlaybackRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiLivePlaybackRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiLivePlaybackRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiLivePlaybackRequest) GetRequest() string {
	return this.Request
}
func (this *OapiLivePlaybackRequest) GetApiMethodName() string {
	return "dingtalk.oapi.live.playback"
}
func (this *OapiLivePlaybackRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiLivePlaybackRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiLivePlaybackRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiLivePlaybackRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiLivePlaybackRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiLivePlaybackRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiLivePlaybackRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiLivePlaybackRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PlayBackReqModel struct {
	EndTime   time.Time `json:"end_time,omitempty"`
	Offset    int64     `json:"offset,omitempty"`
	PageSize  int64     `json:"page_size,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
}
type OapiLivePlaybackResponse struct {
	taobao.TaobaoResponse
	Result PlayBackRespModel `json:"result,omitempty"`
}
type PlayBackModel struct {
	CoverUrl       string    `json:"cover_url,omitempty"`
	EndTime        time.Time `json:"end_time,omitempty"`
	Intro          string    `json:"intro,omitempty"`
	LandScape      bool      `json:"land_scape,omitempty"`
	PlaybackUrl    string    `json:"playback_url,omitempty"`
	Shared         bool      `json:"shared,omitempty"`
	StartTime      time.Time `json:"start_time,omitempty"`
	TimeLength     int64     `json:"time_length,omitempty"`
	Title          string    `json:"title,omitempty"`
	TotalJoinCount int64     `json:"total_join_count,omitempty"`
	TotalViewCount int64     `json:"total_view_count,omitempty"`
	UserNick       string    `json:"user_nick,omitempty"`
	Userid         string    `json:"userid,omitempty"`
	Uuid           string    `json:"uuid,omitempty"`
}
type PlayBackRespModel struct {
	AllCount     int64           `json:"all_count,omitempty"`
	HasMore      bool            `json:"has_more,omitempty"`
	Offset       int64           `json:"offset,omitempty"`
	PageSize     int64           `json:"page_size,omitempty"`
	PlayBackList []PlayBackModel `json:"play_back_list,omitempty"`
}
