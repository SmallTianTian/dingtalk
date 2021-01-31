package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringCooplistGetRequest() *OapiCateringCooplistGetRequest {
	return &OapiCateringCooplistGetRequest{}
}

type OapiCateringCooplistGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringCooplistGetResponse
	CoopStatus      int64
	OffSet          int64
	PgSize          int64
	ShopId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCateringCooplistGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringCooplistGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringCooplistGetRequest) SetCoopStatus(coopStatus2 int64) {
	this.CoopStatus = coopStatus2
}
func (this *OapiCateringCooplistGetRequest) GetCoopStatus() int64 {
	return this.CoopStatus
}
func (this *OapiCateringCooplistGetRequest) SetOffSet(offSet2 int64) {
	this.OffSet = offSet2
}
func (this *OapiCateringCooplistGetRequest) GetOffSet() int64 {
	return this.OffSet
}
func (this *OapiCateringCooplistGetRequest) SetPgSize(pgSize2 int64) {
	this.PgSize = pgSize2
}
func (this *OapiCateringCooplistGetRequest) GetPgSize() int64 {
	return this.PgSize
}
func (this *OapiCateringCooplistGetRequest) SetShopId(shopId2 string) {
	this.ShopId = shopId2
}
func (this *OapiCateringCooplistGetRequest) GetShopId() string {
	return this.ShopId
}
func (this *OapiCateringCooplistGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.cooplist.get"
}
func (this *OapiCateringCooplistGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringCooplistGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringCooplistGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringCooplistGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringCooplistGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["coop_status"] = this.CoopStatus
	txtParams["off_set"] = this.OffSet
	txtParams["pg_size"] = this.PgSize
	txtParams["shop_id"] = this.ShopId
	return txtParams
}
func (this *OapiCateringCooplistGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CoopStatus, "coopStatus"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringCooplistGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringCooplistGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringCooplistGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
