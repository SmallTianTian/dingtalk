package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFinanceLoanNotifyCreditRequest() *OapiFinanceLoanNotifyCreditRequest {
	return &OapiFinanceLoanNotifyCreditRequest{}
}

type OapiFinanceLoanNotifyCreditRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFinanceLoanNotifyCreditResponse
	Amount          int64
	AvailableAmount int64
	CompleteTime    string
	CreditNo        string
	Extension       string
	IdCardNo        string
	NextApplyDay    string
	OpenChannelName string
	OpenProductCode string
	OpenProductName string
	OpenProductType string
	RefuseCode      string
	RefuseReason    string
	Status          int64
	SubmitTime      string
	TopHttpMethod   string
	TopResponseType string
	UserMobile      string
}

func (this *OapiFinanceLoanNotifyCreditRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetAmount(amount2 int64) {
	this.Amount = amount2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetAmount() int64 {
	return this.Amount
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetAvailableAmount(availableAmount2 int64) {
	this.AvailableAmount = availableAmount2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetAvailableAmount() int64 {
	return this.AvailableAmount
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetCompleteTime(completeTime2 string) {
	this.CompleteTime = completeTime2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetCompleteTime() string {
	return this.CompleteTime
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetCreditNo(creditNo2 string) {
	this.CreditNo = creditNo2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetCreditNo() string {
	return this.CreditNo
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetExtension(extension2 string) {
	this.Extension = extension2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetExtension() string {
	return this.Extension
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetIdCardNo(idCardNo2 string) {
	this.IdCardNo = idCardNo2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetIdCardNo() string {
	return this.IdCardNo
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetNextApplyDay(nextApplyDay2 string) {
	this.NextApplyDay = nextApplyDay2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetNextApplyDay() string {
	return this.NextApplyDay
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetOpenChannelName(openChannelName2 string) {
	this.OpenChannelName = openChannelName2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetOpenChannelName() string {
	return this.OpenChannelName
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetOpenProductCode(openProductCode2 string) {
	this.OpenProductCode = openProductCode2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetOpenProductCode() string {
	return this.OpenProductCode
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetOpenProductName(openProductName2 string) {
	this.OpenProductName = openProductName2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetOpenProductName() string {
	return this.OpenProductName
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetOpenProductType(openProductType2 string) {
	this.OpenProductType = openProductType2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetOpenProductType() string {
	return this.OpenProductType
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetRefuseCode(refuseCode2 string) {
	this.RefuseCode = refuseCode2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetRefuseCode() string {
	return this.RefuseCode
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetRefuseReason(refuseReason2 string) {
	this.RefuseReason = refuseReason2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetRefuseReason() string {
	return this.RefuseReason
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetStatus(status2 int64) {
	this.Status = status2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetStatus() int64 {
	return this.Status
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetSubmitTime(submitTime2 string) {
	this.SubmitTime = submitTime2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetSubmitTime() string {
	return this.SubmitTime
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetUserMobile(userMobile2 string) {
	this.UserMobile = userMobile2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetUserMobile() string {
	return this.UserMobile
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.loan.notify.credit"
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceLoanNotifyCreditRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["amount"] = this.Amount
	txtParams["available_amount"] = this.AvailableAmount
	txtParams["complete_time"] = this.CompleteTime
	txtParams["credit_no"] = this.CreditNo
	txtParams["extension"] = this.Extension
	txtParams["id_card_no"] = this.IdCardNo
	txtParams["next_apply_day"] = this.NextApplyDay
	txtParams["open_channel_name"] = this.OpenChannelName
	txtParams["open_product_code"] = this.OpenProductCode
	txtParams["open_product_name"] = this.OpenProductName
	txtParams["open_product_type"] = this.OpenProductType
	txtParams["refuse_code"] = this.RefuseCode
	txtParams["refuse_reason"] = this.RefuseReason
	txtParams["status"] = this.Status
	txtParams["submit_time"] = this.SubmitTime
	txtParams["user_mobile"] = this.UserMobile
	return txtParams
}
func (this *OapiFinanceLoanNotifyCreditRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Amount, "amount"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceLoanNotifyCreditRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFinanceLoanNotifyCreditResponse struct {
	taobao.TaobaoResponse
	Result  OpenCommonResult `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
