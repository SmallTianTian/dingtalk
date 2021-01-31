package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripTrainCitySuggestRequest() *OapiAlitripBtripTrainCitySuggestRequest {
	return &OapiAlitripBtripTrainCitySuggestRequest{}
}

type OapiAlitripBtripTrainCitySuggestRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripTrainCitySuggestResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripTrainCitySuggestRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripTrainCitySuggestRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripTrainCitySuggestRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripTrainCitySuggestRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripTrainCitySuggestRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.train.city.suggest"
}
func (this *OapiAlitripBtripTrainCitySuggestRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripTrainCitySuggestRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripTrainCitySuggestRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripTrainCitySuggestRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripTrainCitySuggestRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripTrainCitySuggestRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripTrainCitySuggestRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripTrainCitySuggestRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAlitripBtripTrainCitySuggestResponse struct {
	taobao.TaobaoResponse
	Errcode int64     `json:"errcode,omitempty"`
	Errmsg  string    `json:"errmsg,omitempty"`
	Result  SuggestRs `json:"result,omitempty"`
	Success bool      `json:"success,omitempty"`
}
