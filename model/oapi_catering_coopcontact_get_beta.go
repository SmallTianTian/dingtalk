package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringCoopcontactGetBetaRequest() *OapiCateringCoopcontactGetBetaRequest {
	return &OapiCateringCoopcontactGetBetaRequest{}
}

type OapiCateringCoopcontactGetBetaRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringCoopcontactGetBetaResponse
	CallerUserid    string
	CorpEndCorpId   string
	ShopId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCateringCoopcontactGetBetaRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringCoopcontactGetBetaRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringCoopcontactGetBetaRequest) SetCallerUserid(callerUserid2 string) {
	this.CallerUserid = callerUserid2
}
func (this *OapiCateringCoopcontactGetBetaRequest) GetCallerUserid() string {
	return this.CallerUserid
}
func (this *OapiCateringCoopcontactGetBetaRequest) SetCorpEndCorpId(corpEndCorpId2 string) {
	this.CorpEndCorpId = corpEndCorpId2
}
func (this *OapiCateringCoopcontactGetBetaRequest) GetCorpEndCorpId() string {
	return this.CorpEndCorpId
}
func (this *OapiCateringCoopcontactGetBetaRequest) SetShopId(shopId2 string) {
	this.ShopId = shopId2
}
func (this *OapiCateringCoopcontactGetBetaRequest) GetShopId() string {
	return this.ShopId
}
func (this *OapiCateringCoopcontactGetBetaRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.coopcontact.get.beta"
}
func (this *OapiCateringCoopcontactGetBetaRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringCoopcontactGetBetaRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringCoopcontactGetBetaRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringCoopcontactGetBetaRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringCoopcontactGetBetaRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["caller_userid"] = this.CallerUserid
	txtParams["corp_end_corp_id"] = this.CorpEndCorpId
	txtParams["shop_id"] = this.ShopId
	return txtParams
}
func (this *OapiCateringCoopcontactGetBetaRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CallerUserid, "callerUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringCoopcontactGetBetaRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringCoopcontactGetBetaRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringCoopcontactGetBetaResponse struct {
	taobao.TaobaoResponse

	Expiration int64  `json:"expiration,omitempty"`
	PriMobile  string `json:"pri_mobile,omitempty"`
	Success    bool   `json:"success,omitempty"`
}
