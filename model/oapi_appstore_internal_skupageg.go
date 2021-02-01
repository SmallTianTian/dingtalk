package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAppstoreInternalSkupageGetRequest() *OapiAppstoreInternalSkupageGetRequest {
	return &OapiAppstoreInternalSkupageGetRequest{}
}

type OapiAppstoreInternalSkupageGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAppstoreInternalSkupageGetResponse
	CallbackPage    string
	ExtendParam     string
	GoodsCode       string
	ItemCode        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAppstoreInternalSkupageGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAppstoreInternalSkupageGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAppstoreInternalSkupageGetRequest) SetCallbackPage(callbackPage2 string) {
	this.CallbackPage = callbackPage2
}
func (this *OapiAppstoreInternalSkupageGetRequest) GetCallbackPage() string {
	return this.CallbackPage
}
func (this *OapiAppstoreInternalSkupageGetRequest) SetExtendParam(extendParam2 string) {
	this.ExtendParam = extendParam2
}
func (this *OapiAppstoreInternalSkupageGetRequest) GetExtendParam() string {
	return this.ExtendParam
}
func (this *OapiAppstoreInternalSkupageGetRequest) SetGoodsCode(goodsCode2 string) {
	this.GoodsCode = goodsCode2
}
func (this *OapiAppstoreInternalSkupageGetRequest) GetGoodsCode() string {
	return this.GoodsCode
}
func (this *OapiAppstoreInternalSkupageGetRequest) SetItemCode(itemCode2 string) {
	this.ItemCode = itemCode2
}
func (this *OapiAppstoreInternalSkupageGetRequest) GetItemCode() string {
	return this.ItemCode
}
func (this *OapiAppstoreInternalSkupageGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.appstore.internal.skupage.get"
}
func (this *OapiAppstoreInternalSkupageGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAppstoreInternalSkupageGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAppstoreInternalSkupageGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAppstoreInternalSkupageGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAppstoreInternalSkupageGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["callback_page"] = this.CallbackPage
	txtParams["extend_param"] = this.ExtendParam
	txtParams["goods_code"] = this.GoodsCode
	txtParams["item_code"] = this.ItemCode
	return txtParams
}
func (this *OapiAppstoreInternalSkupageGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GoodsCode, "goodsCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAppstoreInternalSkupageGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAppstoreInternalSkupageGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAppstoreInternalSkupageGetResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
