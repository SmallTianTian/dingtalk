package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripCostCenterNewRequest() *OapiAlitripBtripCostCenterNewRequest {
	return &OapiAlitripBtripCostCenterNewRequest{}
}

type OapiAlitripBtripCostCenterNewRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripCostCenterNewResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripCostCenterNewRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripCostCenterNewRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripCostCenterNewRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripCostCenterNewRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripCostCenterNewRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.cost.center.new"
}
func (this *OapiAlitripBtripCostCenterNewRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripCostCenterNewRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripCostCenterNewRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripCostCenterNewRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripCostCenterNewRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripCostCenterNewRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripCostCenterNewRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripCostCenterNewRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenCostCenterSaveRq struct {
	AlipayNo    string `json:"alipay_no,omitempty"`
	Corpid      string `json:"corpid,omitempty"`
	Number      string `json:"number,omitempty"`
	Scope       int64  `json:"scope,omitempty"`
	ThirdpartId string `json:"thirdpart_id,omitempty"`
	Title       string `json:"title,omitempty"`
}
type OapiAlitripBtripCostCenterNewResponse struct {
	taobao.TaobaoResponse
	Errcode int64                `json:"errcode,omitempty"`
	Errmsg  string               `json:"errmsg,omitempty"`
	Result  OpenCostCenterSaveRs `json:"result,omitempty"`
	Success bool                 `json:"success,omitempty"`
}
type OpenCostCenterSaveRs struct {
	Id int64 `json:"id,omitempty"`
}
