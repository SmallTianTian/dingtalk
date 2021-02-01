package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripCostCenterDeleteRequest() *OapiAlitripBtripCostCenterDeleteRequest {
	return &OapiAlitripBtripCostCenterDeleteRequest{}
}

type OapiAlitripBtripCostCenterDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripCostCenterDeleteResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripCostCenterDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripCostCenterDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripCostCenterDeleteRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripCostCenterDeleteRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripCostCenterDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.cost.center.delete"
}
func (this *OapiAlitripBtripCostCenterDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripCostCenterDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripCostCenterDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripCostCenterDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripCostCenterDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripCostCenterDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripCostCenterDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripCostCenterDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenCostCenterDeleteRq struct {
	Corpid      string `json:"corpid,omitempty"`
	ThirdpartId string `json:"thirdpart_id,omitempty"`
}
type OapiAlitripBtripCostCenterDeleteResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
