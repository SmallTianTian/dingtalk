package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripCostCenterEntitySetRequest() *OapiAlitripBtripCostCenterEntitySetRequest {
	return &OapiAlitripBtripCostCenterEntitySetRequest{}
}

type OapiAlitripBtripCostCenterEntitySetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripCostCenterEntitySetResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripCostCenterEntitySetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripCostCenterEntitySetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripCostCenterEntitySetRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripCostCenterEntitySetRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripCostCenterEntitySetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.cost.center.entity.set"
}
func (this *OapiAlitripBtripCostCenterEntitySetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripCostCenterEntitySetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripCostCenterEntitySetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripCostCenterEntitySetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripCostCenterEntitySetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripCostCenterEntitySetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripCostCenterEntitySetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripCostCenterEntitySetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenCostCenterSetEntityRq struct {
	Corpid      string            `json:"corpid,omitempty"`
	EntityList  []OpenOrgEntityDo `json:"entity_list,omitempty"`
	ThirdpartId string            `json:"thirdpart_id,omitempty"`
}
type OapiAlitripBtripCostCenterEntitySetResponse struct {
	taobao.TaobaoResponse
	Result  OpenCostCenterSetEntityRs `json:"result,omitempty"`
	Success bool                      `json:"success,omitempty"`
}
type OpenCostCenterSetEntityRs struct {
	AddNum          int64 `json:"add_num,omitempty"`
	RemoveNum       int64 `json:"remove_num,omitempty"`
	SelectedUserNum int64 `json:"selected_user_num,omitempty"`
}
