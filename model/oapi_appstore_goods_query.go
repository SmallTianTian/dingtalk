package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAppstoreGoodsQueryRequest() *OapiAppstoreGoodsQueryRequest {
	return &OapiAppstoreGoodsQueryRequest{}
}

type OapiAppstoreGoodsQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAppstoreGoodsQueryResponse
	GoodsCode       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAppstoreGoodsQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAppstoreGoodsQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAppstoreGoodsQueryRequest) SetGoodsCode(goodsCode2 string) {
	this.GoodsCode = goodsCode2
}
func (this *OapiAppstoreGoodsQueryRequest) GetGoodsCode() string {
	return this.GoodsCode
}
func (this *OapiAppstoreGoodsQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.appstore.goods.query"
}
func (this *OapiAppstoreGoodsQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAppstoreGoodsQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAppstoreGoodsQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAppstoreGoodsQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAppstoreGoodsQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["goods_code"] = this.GoodsCode
	return txtParams
}
func (this *OapiAppstoreGoodsQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GoodsCode, "goodsCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAppstoreGoodsQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAppstoreGoodsQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAppstoreGoodsQueryResponse struct {
	taobao.TaobaoResponse
	Errcode   int64       `json:"errcode,omitempty"`
	Errmsg    string      `json:"errmsg,omitempty"`
	GoodsInfo OpenGoodsVo `json:"goods_info,omitempty"`
}
type OpenGoodsItemCycVo struct {
	AliasName string `json:"alias_name,omitempty"`
	CycNum    int64  `json:"cyc_num,omitempty"`
	CycUnit   int64  `json:"cyc_unit,omitempty"`
}
type OpenGoodsItemVo struct {
	IsTryOuts   bool                 `json:"is_try_outs,omitempty"`
	ItemCode    string               `json:"item_code,omitempty"`
	ItemCycList []OpenGoodsItemCycVo `json:"item_cyc_list,omitempty"`
	ItemName    string               `json:"item_name,omitempty"`
	MaxNum      int64                `json:"max_num,omitempty"`
	MinNum      int64                `json:"min_num,omitempty"`
}
type OpenGoodsVo struct {
	ItemList []OpenGoodsItemVo `json:"item_list,omitempty"`
	Name     string            `json:"name,omitempty"`
}
