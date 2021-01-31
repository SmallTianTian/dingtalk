package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavVideoliveGetRequest() *OapiKacDatavVideoliveGetRequest {
	return &OapiKacDatavVideoliveGetRequest{}
}

type OapiKacDatavVideoliveGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                         OapiKacDatavVideoliveGetResponse
	ParamVideoLiveSummaryRequest string
	TopHttpMethod                string
	TopResponseType              string
}

func (this *OapiKacDatavVideoliveGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavVideoliveGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavVideoliveGetRequest) SetParamVideoLiveSummaryRequest(paramVideoLiveSummaryRequest2 string) {
	this.ParamVideoLiveSummaryRequest = paramVideoLiveSummaryRequest2
}
func (this *OapiKacDatavVideoliveGetRequest) GetParamVideoLiveSummaryRequest() string {
	return this.ParamVideoLiveSummaryRequest
}
func (this *OapiKacDatavVideoliveGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.videolive.get"
}
func (this *OapiKacDatavVideoliveGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavVideoliveGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavVideoliveGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavVideoliveGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavVideoliveGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["param_video_live_summary_request"] = this.ParamVideoLiveSummaryRequest
	return txtParams
}
func (this *OapiKacDatavVideoliveGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavVideoliveGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavVideoliveGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type VideoLiveSummaryRequest struct {
	DataId string `json:"data_id,omitempty"`
}
type OapiKacDatavVideoliveGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64                    `json:"errcode,omitempty"`
	Errmsg  string                   `json:"errmsg,omitempty"`
	Result  VideoLiveSummaryResponse `json:"result,omitempty"`
}
type VideoLiveSummaryResponse struct {
	LiveLaunchSucc5minCnt int64  `json:"live_launch_succ5min_cnt,omitempty"`
	LiveLaunchSuccCnt     int64  `json:"live_launch_succ_cnt,omitempty"`
	LivePlayCnt           int64  `json:"live_play_cnt,omitempty"`
	LivePlayUserCnt       int64  `json:"live_play_user_cnt,omitempty"`
	LiveSuccTimeLen       string `json:"live_succ_time_len,omitempty"`
	MaxUserCnt            int64  `json:"max_user_cnt,omitempty"`
	WatchGroupLiveUserCnt int64  `json:"watch_group_live_user_cnt,omitempty"`
}
