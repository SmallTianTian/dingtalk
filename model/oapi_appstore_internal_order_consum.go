package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAppstoreInternalOrderConsumeRequest() *OapiAppstoreInternalOrderConsumeRequest {
	return &OapiAppstoreInternalOrderConsumeRequest{}
}

type OapiAppstoreInternalOrderConsumeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAppstoreInternalOrderConsumeResponse
	BizOrderId      int64
	Quantity        int64
	RequestId       string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiAppstoreInternalOrderConsumeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAppstoreInternalOrderConsumeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAppstoreInternalOrderConsumeRequest) SetBizOrderId(bizOrderId2 int64) {
	this.BizOrderId = bizOrderId2
}
func (this *OapiAppstoreInternalOrderConsumeRequest) GetBizOrderId() int64 {
	return this.BizOrderId
}
func (this *OapiAppstoreInternalOrderConsumeRequest) SetQuantity(quantity2 int64) {
	this.Quantity = quantity2
}
func (this *OapiAppstoreInternalOrderConsumeRequest) GetQuantity() int64 {
	return this.Quantity
}
func (this *OapiAppstoreInternalOrderConsumeRequest) SetRequestId(requestId2 string) {
	this.RequestId = requestId2
}
func (this *OapiAppstoreInternalOrderConsumeRequest) GetRequestId() string {
	return this.RequestId
}
func (this *OapiAppstoreInternalOrderConsumeRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAppstoreInternalOrderConsumeRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAppstoreInternalOrderConsumeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.appstore.internal.order.consume"
}
func (this *OapiAppstoreInternalOrderConsumeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAppstoreInternalOrderConsumeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAppstoreInternalOrderConsumeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAppstoreInternalOrderConsumeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAppstoreInternalOrderConsumeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_order_id"] = this.BizOrderId
	txtParams["quantity"] = this.Quantity
	txtParams["request_id"] = this.RequestId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAppstoreInternalOrderConsumeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizOrderId, "bizOrderId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAppstoreInternalOrderConsumeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAppstoreInternalOrderConsumeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAppstoreInternalOrderConsumeResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}
