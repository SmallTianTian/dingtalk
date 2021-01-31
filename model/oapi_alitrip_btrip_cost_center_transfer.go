package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripCostCenterTransferRequest() *OapiAlitripBtripCostCenterTransferRequest {
	return &OapiAlitripBtripCostCenterTransferRequest{}
}

type OapiAlitripBtripCostCenterTransferRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripCostCenterTransferResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripCostCenterTransferRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripCostCenterTransferRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripCostCenterTransferRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripCostCenterTransferRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripCostCenterTransferRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.cost.center.transfer"
}
func (this *OapiAlitripBtripCostCenterTransferRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripCostCenterTransferRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripCostCenterTransferRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripCostCenterTransferRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripCostCenterTransferRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripCostCenterTransferRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripCostCenterTransferRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripCostCenterTransferRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenCostCenterTransferRq struct {
	Corpid       string `json:"corpid,omitempty"`
	CostCenterId int64  `json:"cost_center_id,omitempty"`
	ThirdpartId  string `json:"thirdpart_id,omitempty"`
}
type OapiAlitripBtripCostCenterTransferResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
