package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripCategoryAddressGetRequest() *OapiAlitripBtripCategoryAddressGetRequest {
	return &OapiAlitripBtripCategoryAddressGetRequest{}
}

type OapiAlitripBtripCategoryAddressGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripCategoryAddressGetResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripCategoryAddressGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripCategoryAddressGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripCategoryAddressGetRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripCategoryAddressGetRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripCategoryAddressGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.category.address.get"
}
func (this *OapiAlitripBtripCategoryAddressGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripCategoryAddressGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripCategoryAddressGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripCategoryAddressGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripCategoryAddressGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripCategoryAddressGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripCategoryAddressGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripCategoryAddressGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenJumpInfoRq struct {
	Corpid      string `json:"corpid,omitempty"`
	ItineraryId string `json:"itinerary_id,omitempty"`
	Type        int64  `json:"type,omitempty"`
	Userid      string `json:"userid,omitempty"`
}
type OapiAlitripBtripCategoryAddressGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64          `json:"errcode,omitempty"`
	Errmsg  string         `json:"errmsg,omitempty"`
	Result  OpenJumpInfoRs `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type OpenJumpInfoRs struct {
	Url string `json:"url,omitempty"`
}
