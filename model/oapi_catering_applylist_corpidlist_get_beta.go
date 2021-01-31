package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringApplylistCorpidlistGetBetaRequest() *OapiCateringApplylistCorpidlistGetBetaRequest {
	return &OapiCateringApplylistCorpidlistGetBetaRequest{}
}

type OapiCateringApplylistCorpidlistGetBetaRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringApplylistCorpidlistGetBetaResponse
	ShopIdList      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCateringApplylistCorpidlistGetBetaRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringApplylistCorpidlistGetBetaRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringApplylistCorpidlistGetBetaRequest) SetShopIdList(shopIdList2 string) {
	this.ShopIdList = shopIdList2
}
func (this *OapiCateringApplylistCorpidlistGetBetaRequest) GetShopIdList() string {
	return this.ShopIdList
}
func (this *OapiCateringApplylistCorpidlistGetBetaRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.applylist.corpidlist.get.beta"
}
func (this *OapiCateringApplylistCorpidlistGetBetaRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringApplylistCorpidlistGetBetaRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringApplylistCorpidlistGetBetaRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringApplylistCorpidlistGetBetaRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringApplylistCorpidlistGetBetaRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["shop_id_list"] = this.ShopIdList
	return txtParams
}
func (this *OapiCateringApplylistCorpidlistGetBetaRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ShopIdList, "shopIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringApplylistCorpidlistGetBetaRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringApplylistCorpidlistGetBetaRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringApplylistCorpidlistGetBetaResponse struct {
	taobao.TaobaoResponse
	Result  []Result `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
