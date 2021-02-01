package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavVideoconfGetRequest() *OapiKacDatavVideoconfGetRequest {
	return &OapiKacDatavVideoconfGetRequest{}
}

type OapiKacDatavVideoconfGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                   OapiKacDatavVideoconfGetResponse
	ParamMcsSummaryRequest string
	TopHttpMethod          string
	TopResponseType        string
}

func (this *OapiKacDatavVideoconfGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavVideoconfGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavVideoconfGetRequest) SetParamMcsSummaryRequest(paramMcsSummaryRequest2 string) {
	this.ParamMcsSummaryRequest = paramMcsSummaryRequest2
}
func (this *OapiKacDatavVideoconfGetRequest) GetParamMcsSummaryRequest() string {
	return this.ParamMcsSummaryRequest
}
func (this *OapiKacDatavVideoconfGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.videoconf.get"
}
func (this *OapiKacDatavVideoconfGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavVideoconfGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavVideoconfGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavVideoconfGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavVideoconfGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["param_mcs_summary_request"] = this.ParamMcsSummaryRequest
	return txtParams
}
func (this *OapiKacDatavVideoconfGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavVideoconfGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavVideoconfGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiKacDatavVideoconfGetResponse struct {
	taobao.TaobaoResponse
	Errcode  int64              `json:"errcode,omitempty"`
	Errormsg string             `json:"errormsg,omitempty"`
	Result   McsSummaryResponse `json:"result,omitempty"`
}
type McsSummaryResponse struct {
	JoinVideoConfLen        int64 `json:"join_video_conf_len,omitempty"`
	JoinVideoConfSeccUsrCnt int64 `json:"join_video_conf_secc_usr_cnt,omitempty"`
	JoinVideoConfSeccUsrNum int64 `json:"join_video_conf_secc_usr_num,omitempty"`
	JoinVideoConfUsrCnt     int64 `json:"join_video_conf_usr_cnt,omitempty"`
	StartVideoConfCnt       int64 `json:"start_video_conf_cnt,omitempty"`
	StartVideoConfSeccCnt   int64 `json:"start_video_conf_secc_cnt,omitempty"`
	StartVideoConfUsrNum    int64 `json:"start_video_conf_usr_num,omitempty"`
	VideoConfAveUsrNum      int64 `json:"video_conf_ave_usr_num,omitempty"`
}
