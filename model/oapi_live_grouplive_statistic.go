package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiLiveGroupliveStatisticsRequest() *OapiLiveGroupliveStatisticsRequest {
	return &OapiLiveGroupliveStatisticsRequest{}
}

type OapiLiveGroupliveStatisticsRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiLiveGroupliveStatisticsResponse
	Cid             int64
	LiveUuid        string
	OpenId          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiLiveGroupliveStatisticsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiLiveGroupliveStatisticsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiLiveGroupliveStatisticsRequest) SetCid(cid2 int64) {
	this.Cid = cid2
}
func (this *OapiLiveGroupliveStatisticsRequest) GetCid() int64 {
	return this.Cid
}
func (this *OapiLiveGroupliveStatisticsRequest) SetLiveUuid(liveUuid2 string) {
	this.LiveUuid = liveUuid2
}
func (this *OapiLiveGroupliveStatisticsRequest) GetLiveUuid() string {
	return this.LiveUuid
}
func (this *OapiLiveGroupliveStatisticsRequest) SetOpenId(openId2 int64) {
	this.OpenId = openId2
}
func (this *OapiLiveGroupliveStatisticsRequest) GetOpenId() int64 {
	return this.OpenId
}
func (this *OapiLiveGroupliveStatisticsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.live.grouplive.statistics"
}
func (this *OapiLiveGroupliveStatisticsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiLiveGroupliveStatisticsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiLiveGroupliveStatisticsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiLiveGroupliveStatisticsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiLiveGroupliveStatisticsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cid"] = this.Cid
	txtParams["live_uuid"] = this.LiveUuid
	txtParams["open_id"] = this.OpenId
	return txtParams
}
func (this *OapiLiveGroupliveStatisticsRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.LiveUuid, "liveUuid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiLiveGroupliveStatisticsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiLiveGroupliveStatisticsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiLiveGroupliveStatisticsResponse struct {
	taobao.TaobaoResponse
	Result GroupLiveStatistics `json:"result,omitempty"`
}
