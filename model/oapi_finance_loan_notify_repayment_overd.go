package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFinanceLoanNotifyRepaymentOverdueRequest() *OapiFinanceLoanNotifyRepaymentOverdueRequest {
	return &OapiFinanceLoanNotifyRepaymentOverdueRequest{}
}

type OapiFinanceLoanNotifyRepaymentOverdueRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFinanceLoanNotifyRepaymentOverdueResponse
	IdCardNo        string
	IntOvdDays      int64
	LoanOrderNo     string
	OpenChannelName string
	OpenProductName string
	OvdTerms        string
	PaidInterest    int64
	PaidPenalty     int64
	PaidPrincipal   int64
	PaidTotalAmount int64
	PrinOvdDays     int64
	RepayRealDate   string
	RepaymentTerms  string
	TopHttpMethod   string
	TopResponseType string
	UserMobile      string
}

func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetIdCardNo(idCardNo2 string) {
	this.IdCardNo = idCardNo2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetIdCardNo() string {
	return this.IdCardNo
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetIntOvdDays(intOvdDays2 int64) {
	this.IntOvdDays = intOvdDays2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetIntOvdDays() int64 {
	return this.IntOvdDays
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetLoanOrderNo(loanOrderNo2 string) {
	this.LoanOrderNo = loanOrderNo2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetLoanOrderNo() string {
	return this.LoanOrderNo
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetOpenChannelName(openChannelName2 string) {
	this.OpenChannelName = openChannelName2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetOpenChannelName() string {
	return this.OpenChannelName
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetOpenProductName(openProductName2 string) {
	this.OpenProductName = openProductName2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetOpenProductName() string {
	return this.OpenProductName
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetOvdTerms(ovdTerms2 string) {
	this.OvdTerms = ovdTerms2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetOvdTerms() string {
	return this.OvdTerms
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetPaidInterest(paidInterest2 int64) {
	this.PaidInterest = paidInterest2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetPaidInterest() int64 {
	return this.PaidInterest
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetPaidPenalty(paidPenalty2 int64) {
	this.PaidPenalty = paidPenalty2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetPaidPenalty() int64 {
	return this.PaidPenalty
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetPaidPrincipal(paidPrincipal2 int64) {
	this.PaidPrincipal = paidPrincipal2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetPaidPrincipal() int64 {
	return this.PaidPrincipal
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetPaidTotalAmount(paidTotalAmount2 int64) {
	this.PaidTotalAmount = paidTotalAmount2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetPaidTotalAmount() int64 {
	return this.PaidTotalAmount
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetPrinOvdDays(prinOvdDays2 int64) {
	this.PrinOvdDays = prinOvdDays2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetPrinOvdDays() int64 {
	return this.PrinOvdDays
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetRepayRealDate(repayRealDate2 string) {
	this.RepayRealDate = repayRealDate2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetRepayRealDate() string {
	return this.RepayRealDate
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetRepaymentTerms(repaymentTerms2 string) {
	this.RepaymentTerms = repaymentTerms2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetRepaymentTerms() string {
	return this.RepaymentTerms
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetUserMobile(userMobile2 string) {
	this.UserMobile = userMobile2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetUserMobile() string {
	return this.UserMobile
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.loan.notify.repayment.overdue"
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["id_card_no"] = this.IdCardNo
	txtParams["int_ovd_days"] = this.IntOvdDays
	txtParams["loan_order_no"] = this.LoanOrderNo
	txtParams["open_channel_name"] = this.OpenChannelName
	txtParams["open_product_name"] = this.OpenProductName
	txtParams["ovd_terms"] = this.OvdTerms
	txtParams["paid_interest"] = this.PaidInterest
	txtParams["paid_penalty"] = this.PaidPenalty
	txtParams["paid_principal"] = this.PaidPrincipal
	txtParams["paid_total_amount"] = this.PaidTotalAmount
	txtParams["prin_ovd_days"] = this.PrinOvdDays
	txtParams["repay_real_date"] = this.RepayRealDate
	txtParams["repayment_terms"] = this.RepaymentTerms
	txtParams["user_mobile"] = this.UserMobile
	return txtParams
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.IdCardNo, "idCardNo"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceLoanNotifyRepaymentOverdueRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFinanceLoanNotifyRepaymentOverdueResponse struct {
	taobao.TaobaoResponse
	Result  OpenCommonResult `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
