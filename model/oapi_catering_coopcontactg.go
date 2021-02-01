package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringCoopcontactGetRequest() *OapiCateringCoopcontactGetRequest {
	return &OapiCateringCoopcontactGetRequest{}
}

type OapiCateringCoopcontactGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringCoopcontactGetResponse
	CallerUserid    string
	CorpEndCorpId   string
	ShopId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCateringCoopcontactGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringCoopcontactGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringCoopcontactGetRequest) SetCallerUserid(callerUserid2 string) {
	this.CallerUserid = callerUserid2
}
func (this *OapiCateringCoopcontactGetRequest) GetCallerUserid() string {
	return this.CallerUserid
}
func (this *OapiCateringCoopcontactGetRequest) SetCorpEndCorpId(corpEndCorpId2 string) {
	this.CorpEndCorpId = corpEndCorpId2
}
func (this *OapiCateringCoopcontactGetRequest) GetCorpEndCorpId() string {
	return this.CorpEndCorpId
}
func (this *OapiCateringCoopcontactGetRequest) SetShopId(shopId2 string) {
	this.ShopId = shopId2
}
func (this *OapiCateringCoopcontactGetRequest) GetShopId() string {
	return this.ShopId
}
func (this *OapiCateringCoopcontactGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.coopcontact.get"
}
func (this *OapiCateringCoopcontactGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringCoopcontactGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringCoopcontactGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringCoopcontactGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringCoopcontactGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["caller_userid"] = this.CallerUserid
	txtParams["corp_end_corp_id"] = this.CorpEndCorpId
	txtParams["shop_id"] = this.ShopId
	return txtParams
}
func (this *OapiCateringCoopcontactGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CallerUserid, "callerUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringCoopcontactGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringCoopcontactGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringCoopcontactGetResponse struct {
	taobao.TaobaoResponse

	Expiration int64  `json:"expiration,omitempty"`
	PriMobile  string `json:"pri_mobile,omitempty"`
	Success    bool   `json:"success,omitempty"`
}
