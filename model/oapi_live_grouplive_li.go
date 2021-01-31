package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiLiveGroupliveListRequest() *OapiLiveGroupliveListRequest {
	return &OapiLiveGroupliveListRequest{}
}

type OapiLiveGroupliveListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiLiveGroupliveListResponse
	Cid             int64
	FromTime        int64
	OpenId          int64
	ToTime          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiLiveGroupliveListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiLiveGroupliveListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiLiveGroupliveListRequest) SetCid(cid2 int64) {
	this.Cid = cid2
}
func (this *OapiLiveGroupliveListRequest) GetCid() int64 {
	return this.Cid
}
func (this *OapiLiveGroupliveListRequest) SetFromTime(fromTime2 int64) {
	this.FromTime = fromTime2
}
func (this *OapiLiveGroupliveListRequest) GetFromTime() int64 {
	return this.FromTime
}
func (this *OapiLiveGroupliveListRequest) SetOpenId(openId2 int64) {
	this.OpenId = openId2
}
func (this *OapiLiveGroupliveListRequest) GetOpenId() int64 {
	return this.OpenId
}
func (this *OapiLiveGroupliveListRequest) SetToTime(toTime2 int64) {
	this.ToTime = toTime2
}
func (this *OapiLiveGroupliveListRequest) GetToTime() int64 {
	return this.ToTime
}
func (this *OapiLiveGroupliveListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.live.grouplive.list"
}
func (this *OapiLiveGroupliveListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiLiveGroupliveListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiLiveGroupliveListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiLiveGroupliveListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiLiveGroupliveListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cid"] = this.Cid
	txtParams["from_time"] = this.FromTime
	txtParams["open_id"] = this.OpenId
	txtParams["to_time"] = this.ToTime
	return txtParams
}
func (this *OapiLiveGroupliveListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Cid, "cid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiLiveGroupliveListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiLiveGroupliveListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiLiveGroupliveListResponse struct {
	taobao.TaobaoResponse
	Result []Result `json:"result,omitempty"`
}
