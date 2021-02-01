package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripCostCenterQueryRequest() *OapiAlitripBtripCostCenterQueryRequest {
	return &OapiAlitripBtripCostCenterQueryRequest{}
}

type OapiAlitripBtripCostCenterQueryRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripCostCenterQueryResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripCostCenterQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripCostCenterQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripCostCenterQueryRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripCostCenterQueryRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripCostCenterQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.cost.center.query"
}
func (this *OapiAlitripBtripCostCenterQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripCostCenterQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripCostCenterQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripCostCenterQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripCostCenterQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripCostCenterQueryRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripCostCenterQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripCostCenterQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenCostCenterQueryRq struct {
	Corpid        string `json:"corpid,omitempty"`
	NeedOrgEntity bool   `json:"need_org_entity,omitempty"`
	ThirdpartId   string `json:"thirdpart_id,omitempty"`
	Title         string `json:"title,omitempty"`
	Userid        string `json:"userid,omitempty"`
}
type OapiAlitripBtripCostCenterQueryResponse struct {
	taobao.TaobaoResponse
	CostCenterList []OpenCostCenterQueryRs `json:"cost_center_list,omitempty"`

	Success bool `json:"success,omitempty"`
}
type OpenCostCenterQueryRs struct {
	AlipayNo    string            `json:"alipay_no,omitempty"`
	Corpid      string            `json:"corpid,omitempty"`
	EntityList  []OpenOrgEntityDo `json:"entity_list,omitempty"`
	Id          int64             `json:"id,omitempty"`
	Number      string            `json:"number,omitempty"`
	Scope       int64             `json:"scope,omitempty"`
	ThirdpartId string            `json:"thirdpart_id,omitempty"`
	Title       string            `json:"title,omitempty"`
}
