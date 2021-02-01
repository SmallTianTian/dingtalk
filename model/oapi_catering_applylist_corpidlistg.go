package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCateringApplylistCorpidlistGetRequest() *OapiCateringApplylistCorpidlistGetRequest {
	return &OapiCateringApplylistCorpidlistGetRequest{}
}

type OapiCateringApplylistCorpidlistGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCateringApplylistCorpidlistGetResponse
	ShopIdList      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCateringApplylistCorpidlistGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCateringApplylistCorpidlistGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCateringApplylistCorpidlistGetRequest) SetShopIdList(shopIdList2 string) {
	this.ShopIdList = shopIdList2
}
func (this *OapiCateringApplylistCorpidlistGetRequest) GetShopIdList() string {
	return this.ShopIdList
}
func (this *OapiCateringApplylistCorpidlistGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.catering.applylist.corpidlist.get"
}
func (this *OapiCateringApplylistCorpidlistGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCateringApplylistCorpidlistGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCateringApplylistCorpidlistGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCateringApplylistCorpidlistGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCateringApplylistCorpidlistGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["shop_id_list"] = this.ShopIdList
	return txtParams
}
func (this *OapiCateringApplylistCorpidlistGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ShopIdList, "shopIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCateringApplylistCorpidlistGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCateringApplylistCorpidlistGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCateringApplylistCorpidlistGetResponse struct {
	taobao.TaobaoResponse
	Result  []Result `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
