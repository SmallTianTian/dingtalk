package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringCoopDealBetaRequest() *OapiCateringCoopDealBetaRequest {
	return &OapiCateringCoopDealBetaRequest{}
}

type OapiCateringCoopDealBetaRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringCoopDealBetaResponse
	CoopStatus      int64
	CorpEndCorpId   string
	ShopId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCateringCoopDealBetaRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringCoopDealBetaRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringCoopDealBetaRequest) SetCoopStatus(coopStatus2 int64) {
	this.CoopStatus = coopStatus2
}
func (this *OapiCateringCoopDealBetaRequest) GetCoopStatus() int64 {
	return this.CoopStatus
}
func (this *OapiCateringCoopDealBetaRequest) SetCorpEndCorpId(corpEndCorpId2 string) {
	this.CorpEndCorpId = corpEndCorpId2
}
func (this *OapiCateringCoopDealBetaRequest) GetCorpEndCorpId() string {
	return this.CorpEndCorpId
}
func (this *OapiCateringCoopDealBetaRequest) SetShopId(shopId2 string) {
	this.ShopId = shopId2
}
func (this *OapiCateringCoopDealBetaRequest) GetShopId() string {
	return this.ShopId
}
func (this *OapiCateringCoopDealBetaRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.coop.deal.beta"
}
func (this *OapiCateringCoopDealBetaRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringCoopDealBetaRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringCoopDealBetaRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringCoopDealBetaRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringCoopDealBetaRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["coop_status"] = this.CoopStatus
	txtParams["corp_end_corp_id"] = this.CorpEndCorpId
	txtParams["shop_id"] = this.ShopId
	return txtParams
}
func (this *OapiCateringCoopDealBetaRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CoopStatus, "coopStatus"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringCoopDealBetaRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringCoopDealBetaRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringCoopDealBetaResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
