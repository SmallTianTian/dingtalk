package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripCostCenterEntityAddRequest() *OapiAlitripBtripCostCenterEntityAddRequest {
	return &OapiAlitripBtripCostCenterEntityAddRequest{}
}

type OapiAlitripBtripCostCenterEntityAddRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripCostCenterEntityAddResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripCostCenterEntityAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripCostCenterEntityAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripCostCenterEntityAddRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripCostCenterEntityAddRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripCostCenterEntityAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.cost.center.entity.add"
}
func (this *OapiAlitripBtripCostCenterEntityAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripCostCenterEntityAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripCostCenterEntityAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripCostCenterEntityAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripCostCenterEntityAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripCostCenterEntityAddRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripCostCenterEntityAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripCostCenterEntityAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenOrgEntityDo struct {
	EntityId   string `json:"entity_id,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
}
type OpenCostCenterAddEntityRq struct {
	Corpid      string            `json:"corpid,omitempty"`
	EntityList  []OpenOrgEntityDo `json:"entity_list,omitempty"`
	ThirdpartId string            `json:"thirdpart_id,omitempty"`
}
type OapiAlitripBtripCostCenterEntityAddResponse struct {
	taobao.TaobaoResponse
	Result  OpenCostCenterAddEntityRs `json:"result,omitempty"`
	Success bool                      `json:"success,omitempty"`
}
type OpenCostCenterAddEntityRs struct {
	AddNum          int64 `json:"add_num,omitempty"`
	SelectedUserNum int64 `json:"selected_user_num,omitempty"`
}
