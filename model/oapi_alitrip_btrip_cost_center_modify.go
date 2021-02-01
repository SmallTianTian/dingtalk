package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripCostCenterModifyRequest() *OapiAlitripBtripCostCenterModifyRequest {
	return &OapiAlitripBtripCostCenterModifyRequest{}
}

type OapiAlitripBtripCostCenterModifyRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripCostCenterModifyResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripCostCenterModifyRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripCostCenterModifyRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripCostCenterModifyRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripCostCenterModifyRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripCostCenterModifyRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.cost.center.modify"
}
func (this *OapiAlitripBtripCostCenterModifyRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripCostCenterModifyRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripCostCenterModifyRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripCostCenterModifyRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripCostCenterModifyRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripCostCenterModifyRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripCostCenterModifyRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripCostCenterModifyRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenCostCenterModifyRq struct {
	AlipayNo    string `json:"alipay_no,omitempty"`
	Corpid      string `json:"corpid,omitempty"`
	Number      string `json:"number,omitempty"`
	Scope       int64  `json:"scope,omitempty"`
	ThirdpartId string `json:"thirdpart_id,omitempty"`
	Title       string `json:"title,omitempty"`
}
type OapiAlitripBtripCostCenterModifyResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
