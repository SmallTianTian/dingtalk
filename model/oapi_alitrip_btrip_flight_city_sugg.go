package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripFlightCitySuggestRequest() *OapiAlitripBtripFlightCitySuggestRequest {
	return &OapiAlitripBtripFlightCitySuggestRequest{}
}

type OapiAlitripBtripFlightCitySuggestRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripFlightCitySuggestResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripFlightCitySuggestRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripFlightCitySuggestRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripFlightCitySuggestRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripFlightCitySuggestRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripFlightCitySuggestRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.flight.city.suggest"
}
func (this *OapiAlitripBtripFlightCitySuggestRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripFlightCitySuggestRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripFlightCitySuggestRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripFlightCitySuggestRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripFlightCitySuggestRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripFlightCitySuggestRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripFlightCitySuggestRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripFlightCitySuggestRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SuggestRq struct {
	Corpid  string `json:"corpid,omitempty"`
	Keyword string `json:"keyword,omitempty"`
	Type    int64  `json:"type,omitempty"`
	Userid  string `json:"userid,omitempty"`
}
type OapiAlitripBtripFlightCitySuggestResponse struct {
	taobao.TaobaoResponse
	Result  SuggestRs `json:"result,omitempty"`
	Success bool      `json:"success,omitempty"`
}
type CityVo struct {
	Code       string `json:taobao.ERROR_COD,omitemptyE`
	Distance   int64  `json:"distance,omitempty"`
	Name       string `json:"name,omitempty"`
	TravelName string `json:"travel_name,omitempty"`
}
type SuggestRs struct {
	Cities []CityVo `json:"cities,omitempty"`
	Nearby bool     `json:"nearby,omitempty"`
}
