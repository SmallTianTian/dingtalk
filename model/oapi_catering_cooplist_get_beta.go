package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringCooplistGetBetaRequest() *OapiCateringCooplistGetBetaRequest {
	return &OapiCateringCooplistGetBetaRequest{}
}

type OapiCateringCooplistGetBetaRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringCooplistGetBetaResponse
	CoopStatus      int64
	OffSet          int64
	PgSize          int64
	ShopId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCateringCooplistGetBetaRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringCooplistGetBetaRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringCooplistGetBetaRequest) SetCoopStatus(coopStatus2 int64) {
	this.CoopStatus = coopStatus2
}
func (this *OapiCateringCooplistGetBetaRequest) GetCoopStatus() int64 {
	return this.CoopStatus
}
func (this *OapiCateringCooplistGetBetaRequest) SetOffSet(offSet2 int64) {
	this.OffSet = offSet2
}
func (this *OapiCateringCooplistGetBetaRequest) GetOffSet() int64 {
	return this.OffSet
}
func (this *OapiCateringCooplistGetBetaRequest) SetPgSize(pgSize2 int64) {
	this.PgSize = pgSize2
}
func (this *OapiCateringCooplistGetBetaRequest) GetPgSize() int64 {
	return this.PgSize
}
func (this *OapiCateringCooplistGetBetaRequest) SetShopId(shopId2 string) {
	this.ShopId = shopId2
}
func (this *OapiCateringCooplistGetBetaRequest) GetShopId() string {
	return this.ShopId
}
func (this *OapiCateringCooplistGetBetaRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.cooplist.get.beta"
}
func (this *OapiCateringCooplistGetBetaRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringCooplistGetBetaRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringCooplistGetBetaRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringCooplistGetBetaRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringCooplistGetBetaRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["coop_status"] = this.CoopStatus
	txtParams["off_set"] = this.OffSet
	txtParams["pg_size"] = this.PgSize
	txtParams["shop_id"] = this.ShopId
	return txtParams
}
func (this *OapiCateringCooplistGetBetaRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CoopStatus, "coopStatus"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringCooplistGetBetaRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringCooplistGetBetaRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringCooplistGetBetaResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
type Mealitemlist struct {
	DeliveryTime string `json:"delivery_time,omitempty"`
	Id           int64  `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
}
type MealSettingList struct {
	Address       string         `json:"address,omitempty"`
	AddressDetail string         `json:"address_detail,omitempty"`
	CorpId        string         `json:"corp_id,omitempty"`
	CorpName      string         `json:"corp_name,omitempty"`
	Logo          string         `json:"logo,omitempty"`
	MealItemList  []Mealitemlist `json:"meal_item_list,omitempty"`
	MealTime      int64          `json:"meal_time,omitempty"`
}
