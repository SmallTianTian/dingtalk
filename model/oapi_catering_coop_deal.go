package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringCoopDealRequest() *OapiCateringCoopDealRequest {
	return &OapiCateringCoopDealRequest{}
}

type OapiCateringCoopDealRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringCoopDealResponse
	CoopStatus      int64
	CorpEndCorpId   string
	ShopId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCateringCoopDealRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringCoopDealRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringCoopDealRequest) SetCoopStatus(coopStatus2 int64) {
	this.CoopStatus = coopStatus2
}
func (this *OapiCateringCoopDealRequest) GetCoopStatus() int64 {
	return this.CoopStatus
}
func (this *OapiCateringCoopDealRequest) SetCorpEndCorpId(corpEndCorpId2 string) {
	this.CorpEndCorpId = corpEndCorpId2
}
func (this *OapiCateringCoopDealRequest) GetCorpEndCorpId() string {
	return this.CorpEndCorpId
}
func (this *OapiCateringCoopDealRequest) SetShopId(shopId2 string) {
	this.ShopId = shopId2
}
func (this *OapiCateringCoopDealRequest) GetShopId() string {
	return this.ShopId
}
func (this *OapiCateringCoopDealRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.coop.deal"
}
func (this *OapiCateringCoopDealRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringCoopDealRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringCoopDealRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringCoopDealRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringCoopDealRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["coop_status"] = this.CoopStatus
	txtParams["corp_end_corp_id"] = this.CorpEndCorpId
	txtParams["shop_id"] = this.ShopId
	return txtParams
}
func (this *OapiCateringCoopDealRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CoopStatus, "coopStatus"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringCoopDealRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringCoopDealRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringCoopDealResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
