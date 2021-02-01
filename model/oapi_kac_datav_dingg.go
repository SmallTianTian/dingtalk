package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavDingGetRequest() *OapiKacDatavDingGetRequest {
	return &OapiKacDatavDingGetRequest{}
}

type OapiKacDatavDingGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                    OapiKacDatavDingGetResponse
	DingUsageSummaryRequest string
	TopHttpMethod           string
	TopResponseType         string
}

func (this *OapiKacDatavDingGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavDingGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavDingGetRequest) SetDingUsageSummaryRequest(dingUsageSummaryRequest2 string) {
	this.DingUsageSummaryRequest = dingUsageSummaryRequest2
}
func (this *OapiKacDatavDingGetRequest) GetDingUsageSummaryRequest() string {
	return this.DingUsageSummaryRequest
}
func (this *OapiKacDatavDingGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.ding.get"
}
func (this *OapiKacDatavDingGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavDingGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavDingGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavDingGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavDingGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["ding_usage_summary_request"] = this.DingUsageSummaryRequest
	return txtParams
}
func (this *OapiKacDatavDingGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavDingGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavDingGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiKacDatavDingGetResponse struct {
	taobao.TaobaoResponse
	Result DingUsageSummaryResponse `json:"result,omitempty"`
}
type DingUsageSummaryResponse struct {
	DingAppCnt      int64 `json:"ding_app_cnt,omitempty"`
	DingAppUserCnt  int64 `json:"ding_app_user_cnt,omitempty"`
	DingCallCnt     int64 `json:"ding_call_cnt,omitempty"`
	DingCallUserCnt int64 `json:"ding_call_user_cnt,omitempty"`
	DingCnt         int64 `json:"ding_cnt,omitempty"`
	DingSmsCnt      int64 `json:"ding_sms_cnt,omitempty"`
	DingSmsUserCnt  int64 `json:"ding_sms_user_cnt,omitempty"`
	DingUserCnt     int64 `json:"ding_user_cnt,omitempty"`
	DingVoipCnt     int64 `json:"ding_voip_cnt,omitempty"`
	DingVoipUserCnt int64 `json:"ding_voip_user_cnt,omitempty"`
}
