package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiAlitripBtripCostCenterEntityDeleteRequest() *OapiAlitripBtripCostCenterEntityDeleteRequest {
	return &OapiAlitripBtripCostCenterEntityDeleteRequest{}
}

type OapiAlitripBtripCostCenterEntityDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAlitripBtripCostCenterEntityDeleteResponse
	Rq              string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAlitripBtripCostCenterEntityDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAlitripBtripCostCenterEntityDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAlitripBtripCostCenterEntityDeleteRequest) SetRq(rq2 string) {
	this.Rq = rq2
}
func (this *OapiAlitripBtripCostCenterEntityDeleteRequest) GetRq() string {
	return this.Rq
}
func (this *OapiAlitripBtripCostCenterEntityDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.alitrip.btrip.cost.center.entity.delete"
}
func (this *OapiAlitripBtripCostCenterEntityDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAlitripBtripCostCenterEntityDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAlitripBtripCostCenterEntityDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAlitripBtripCostCenterEntityDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAlitripBtripCostCenterEntityDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["rq"] = this.Rq
	return txtParams
}
func (this *OapiAlitripBtripCostCenterEntityDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiAlitripBtripCostCenterEntityDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAlitripBtripCostCenterEntityDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenCostCenterDeleteEntityRq struct {
	Corpid      string            `json:"corpid,omitempty"`
	DelAll      bool              `json:"del_all,omitempty"`
	EntityList  []OpenOrgEntityDo `json:"entity_list,omitempty"`
	ThirdpartId string            `json:"thirdpart_id,omitempty"`
}
type OapiAlitripBtripCostCenterEntityDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64                        `json:"errcode,omitempty"`
	Errmsg  string                       `json:"errmsg,omitempty"`
	Result  OpenCostCenterDeleteEntityRs `json:"result,omitempty"`
	Success bool                         `json:"success,omitempty"`
}
type OpenCostCenterDeleteEntityRs struct {
	RemoveNum       int64 `json:"remove_num,omitempty"`
	SelectedUserNum int64 `json:"selected_user_num,omitempty"`
}
