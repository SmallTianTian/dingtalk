package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDingpayRedenvelopeSendRequest() *OapiDingpayRedenvelopeSendRequest {
	return &OapiDingpayRedenvelopeSendRequest{}
}

type OapiDingpayRedenvelopeSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingpayRedenvelopeSendResponse
	CorpBizNo       string
	ExtParams       string
	Greetings       string
	PayMethod       string
	PaySign         string
	ReceiverId      string
	SenderId        string
	ThemeId         string
	TopHttpMethod   string
	TopResponseType string
	TotalAmount     string
	Type            string
}

func (this *OapiDingpayRedenvelopeSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingpayRedenvelopeSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingpayRedenvelopeSendRequest) SetCorpBizNo(corpBizNo2 string) {
	this.CorpBizNo = corpBizNo2
}
func (this *OapiDingpayRedenvelopeSendRequest) GetCorpBizNo() string {
	return this.CorpBizNo
}
func (this *OapiDingpayRedenvelopeSendRequest) SetExtParams(extParams2 string) {
	this.ExtParams = extParams2
}
func (this *OapiDingpayRedenvelopeSendRequest) SetExtParamsString(extParams2 string) {
	this.ExtParams = extParams2
}
func (this *OapiDingpayRedenvelopeSendRequest) GetExtParams() string {
	return this.ExtParams
}
func (this *OapiDingpayRedenvelopeSendRequest) SetGreetings(greetings2 string) {
	this.Greetings = greetings2
}
func (this *OapiDingpayRedenvelopeSendRequest) GetGreetings() string {
	return this.Greetings
}
func (this *OapiDingpayRedenvelopeSendRequest) SetPayMethod(payMethod2 string) {
	this.PayMethod = payMethod2
}
func (this *OapiDingpayRedenvelopeSendRequest) GetPayMethod() string {
	return this.PayMethod
}
func (this *OapiDingpayRedenvelopeSendRequest) SetPaySign(paySign2 string) {
	this.PaySign = paySign2
}
func (this *OapiDingpayRedenvelopeSendRequest) GetPaySign() string {
	return this.PaySign
}
func (this *OapiDingpayRedenvelopeSendRequest) SetReceiverId(receiverId2 string) {
	this.ReceiverId = receiverId2
}
func (this *OapiDingpayRedenvelopeSendRequest) GetReceiverId() string {
	return this.ReceiverId
}
func (this *OapiDingpayRedenvelopeSendRequest) SetSenderId(senderId2 string) {
	this.SenderId = senderId2
}
func (this *OapiDingpayRedenvelopeSendRequest) GetSenderId() string {
	return this.SenderId
}
func (this *OapiDingpayRedenvelopeSendRequest) SetThemeId(themeId2 string) {
	this.ThemeId = themeId2
}
func (this *OapiDingpayRedenvelopeSendRequest) GetThemeId() string {
	return this.ThemeId
}
func (this *OapiDingpayRedenvelopeSendRequest) SetTotalAmount(totalAmount2 string) {
	this.TotalAmount = totalAmount2
}
func (this *OapiDingpayRedenvelopeSendRequest) GetTotalAmount() string {
	return this.TotalAmount
}
func (this *OapiDingpayRedenvelopeSendRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *OapiDingpayRedenvelopeSendRequest) GetType() string {
	return this.Type
}
func (this *OapiDingpayRedenvelopeSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingpay.redenvelope.send"
}
func (this *OapiDingpayRedenvelopeSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingpayRedenvelopeSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingpayRedenvelopeSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingpayRedenvelopeSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingpayRedenvelopeSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corp_biz_no"] = this.CorpBizNo
	txtParams["ext_params"] = this.ExtParams
	txtParams["greetings"] = this.Greetings
	txtParams["pay_method"] = this.PayMethod
	txtParams["pay_sign"] = this.PaySign
	txtParams["receiver_id"] = this.ReceiverId
	txtParams["sender_id"] = this.SenderId
	txtParams["theme_id"] = this.ThemeId
	txtParams["total_amount"] = this.TotalAmount
	txtParams["type"] = this.Type
	return txtParams
}
func (this *OapiDingpayRedenvelopeSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.CorpBizNo, "corpBizNo"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingpayRedenvelopeSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingpayRedenvelopeSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingpayRedenvelopeSendResponse struct {
	taobao.TaobaoResponse
	Errcode int64                `json:"errcode,omitempty"`
	Errmsg  string               `json:"errmsg,omitempty"`
	Result  RedEnvelopeGetResult `json:"result,omitempty"`
}
