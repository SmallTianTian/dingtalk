package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripAddressGetRequest() *OapiAlitripBtripAddressGetRequest {
	return &OapiAlitripBtripAddressGetRequest{}
}

type OapiAlitripBtripAddressGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripAddressGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripAddressGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripAddressGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripAddressGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiAlitripBtripAddressGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiAlitripBtripAddressGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.address.get"
}
func (this *OapiAlitripBtripAddressGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripAddressGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripAddressGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripAddressGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripAddressGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiAlitripBtripAddressGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripAddressGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripAddressGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenApiJumpInfoRq struct {
	ActionType  int64  `json:"action_type,omitempty"`
	Corpid      string `json:"corpid,omitempty"`
	ItineraryId string `json:"itinerary_id,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Type        int64  `json:"type,omitempty"`
	Userid      string `json:"userid,omitempty"`
}
type OapiAlitripBtripAddressGetResponse struct {
	taobao.TaobaoResponse
	Result  OpenApiJumpInfoRs `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type OpenApiJumpInfoRs struct {
	Url string `json:"url,omitempty"`
}
