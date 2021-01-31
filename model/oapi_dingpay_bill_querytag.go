package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiDingpayBillQuerytagRequest() *OapiDingpayBillQuerytagRequest {
	return &OapiDingpayBillQuerytagRequest{}
}

type OapiDingpayBillQuerytagRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiDingpayBillQuerytagResponse
	BizCode         string
	DayRange        int64
	SourceAppId     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiDingpayBillQuerytagRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiDingpayBillQuerytagRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiDingpayBillQuerytagRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiDingpayBillQuerytagRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiDingpayBillQuerytagRequest) SetDayRange(dayRange2 int64) {
	this.DayRange = dayRange2
}
func (this *OapiDingpayBillQuerytagRequest) GetDayRange() int64 {
	return this.DayRange
}
func (this *OapiDingpayBillQuerytagRequest) SetSourceAppId(sourceAppId2 string) {
	this.SourceAppId = sourceAppId2
}
func (this *OapiDingpayBillQuerytagRequest) GetSourceAppId() string {
	return this.SourceAppId
}
func (this *OapiDingpayBillQuerytagRequest) GetApiMethodName() string {
	return "dingtalk.oapi.dingpay.bill.querytag"
}
func (this *OapiDingpayBillQuerytagRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiDingpayBillQuerytagRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiDingpayBillQuerytagRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiDingpayBillQuerytagRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiDingpayBillQuerytagRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["day_range"] = this.DayRange
	txtParams["source_app_id"] = this.SourceAppId
	return txtParams
}
func (this *OapiDingpayBillQuerytagRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiDingpayBillQuerytagRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiDingpayBillQuerytagRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiDingpayBillQuerytagResponse struct {
	taobao.TaobaoResponse
	Errcode int64                    `json:"errcode,omitempty"`
	Errmsg  string                   `json:"errmsg,omitempty"`
	Result  BillTagQueryOpenResponse `json:"result,omitempty"`
}
type BillTagQueryOpenResponse struct {
	Tags []string `json:"tags,omitempty"`
}
