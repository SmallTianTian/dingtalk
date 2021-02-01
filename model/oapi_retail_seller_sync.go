package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiRetailSellerSyncRequest() *OapiRetailSellerSyncRequest {
	return &OapiRetailSellerSyncRequest{}
}

type OapiRetailSellerSyncRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRetailSellerSyncResponse
	SellerParam     string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiRetailSellerSyncRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRetailSellerSyncRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRetailSellerSyncRequest) SetSellerParam(sellerParam2 string) {
	this.SellerParam = sellerParam2
}
func (this *OapiRetailSellerSyncRequest) GetSellerParam() string {
	return this.SellerParam
}
func (this *OapiRetailSellerSyncRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiRetailSellerSyncRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiRetailSellerSyncRequest) GetApiMethodName() string {
	return "dingtalk.oapi.retail.seller.sync"
}
func (this *OapiRetailSellerSyncRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRetailSellerSyncRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRetailSellerSyncRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRetailSellerSyncRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRetailSellerSyncRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["seller_param"] = this.SellerParam
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiRetailSellerSyncRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiRetailSellerSyncRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRetailSellerSyncRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SellerParam struct {
	ActionType string `json:"action_type,omitempty"`
	Mobile     string `json:"mobile,omitempty"`
	Name       string `json:"name,omitempty"`
	Nick       string `json:"nick,omitempty"`
	SellerId   int64  `json:"seller_id,omitempty"`
	Source     string `json:"source,omitempty"`
	SrcMobile  string `json:"src_mobile,omitempty"`
}
type OapiRetailSellerSyncResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
