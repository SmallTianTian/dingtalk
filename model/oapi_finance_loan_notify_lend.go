package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFinanceLoanNotifyLendRequest() *OapiFinanceLoanNotifyLendRequest {
	return &OapiFinanceLoanNotifyLendRequest{}
}

type OapiFinanceLoanNotifyLendRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                 OapiFinanceLoanNotifyLendResponse
	AvailableLimit       int64
	BankName             string
	BankcardNo           string
	BillDate             int64
	BillInfoList         string
	CreditAmount         int64
	DailyInterestRate    string
	DiscountsId          string
	FailReason           string
	FirstBillDate        string
	FirstRepayDate       string
	IdCardNo             string
	LastRepayDate        string
	LoanAmount           int64
	LoanEffectiveTime    string
	LoanEndTime          string
	LoanOrderFlowNo      string
	LoanOrderNo          string
	LoanSubmitTime       string
	LoanTxnTime          string
	LoanUsage            string
	OpenChannelName      string
	OpenProductCode      string
	OpenProductName      string
	OpenProductType      string
	PayableInterest      int64
	PayablePrincipal     int64
	PayableTotalAmount   int64
	ReductionTotalAmount int64
	RepayDate            int64
	RepayType            string
	Status               string
	TopHttpMethod        string
	TopResponseType      string
	TotalTerm            int64
	UserMobile           string
	YearLoanInterestRate string
}

func (this *OapiFinanceLoanNotifyLendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceLoanNotifyLendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceLoanNotifyLendRequest) SetAvailableLimit(availableLimit2 int64) {
	this.AvailableLimit = availableLimit2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetAvailableLimit() int64 {
	return this.AvailableLimit
}
func (this *OapiFinanceLoanNotifyLendRequest) SetBankName(bankName2 string) {
	this.BankName = bankName2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetBankName() string {
	return this.BankName
}
func (this *OapiFinanceLoanNotifyLendRequest) SetBankcardNo(bankcardNo2 string) {
	this.BankcardNo = bankcardNo2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetBankcardNo() string {
	return this.BankcardNo
}
func (this *OapiFinanceLoanNotifyLendRequest) SetBillDate(billDate2 int64) {
	this.BillDate = billDate2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetBillDate() int64 {
	return this.BillDate
}
func (this *OapiFinanceLoanNotifyLendRequest) SetBillInfoList(billInfoList2 string) {
	this.BillInfoList = billInfoList2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetBillInfoList() string {
	return this.BillInfoList
}
func (this *OapiFinanceLoanNotifyLendRequest) SetCreditAmount(creditAmount2 int64) {
	this.CreditAmount = creditAmount2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetCreditAmount() int64 {
	return this.CreditAmount
}
func (this *OapiFinanceLoanNotifyLendRequest) SetDailyInterestRate(dailyInterestRate2 string) {
	this.DailyInterestRate = dailyInterestRate2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetDailyInterestRate() string {
	return this.DailyInterestRate
}
func (this *OapiFinanceLoanNotifyLendRequest) SetDiscountsId(discountsId2 string) {
	this.DiscountsId = discountsId2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetDiscountsId() string {
	return this.DiscountsId
}
func (this *OapiFinanceLoanNotifyLendRequest) SetFailReason(failReason2 string) {
	this.FailReason = failReason2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetFailReason() string {
	return this.FailReason
}
func (this *OapiFinanceLoanNotifyLendRequest) SetFirstBillDate(firstBillDate2 string) {
	this.FirstBillDate = firstBillDate2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetFirstBillDate() string {
	return this.FirstBillDate
}
func (this *OapiFinanceLoanNotifyLendRequest) SetFirstRepayDate(firstRepayDate2 string) {
	this.FirstRepayDate = firstRepayDate2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetFirstRepayDate() string {
	return this.FirstRepayDate
}
func (this *OapiFinanceLoanNotifyLendRequest) SetIdCardNo(idCardNo2 string) {
	this.IdCardNo = idCardNo2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetIdCardNo() string {
	return this.IdCardNo
}
func (this *OapiFinanceLoanNotifyLendRequest) SetLastRepayDate(lastRepayDate2 string) {
	this.LastRepayDate = lastRepayDate2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetLastRepayDate() string {
	return this.LastRepayDate
}
func (this *OapiFinanceLoanNotifyLendRequest) SetLoanAmount(loanAmount2 int64) {
	this.LoanAmount = loanAmount2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetLoanAmount() int64 {
	return this.LoanAmount
}
func (this *OapiFinanceLoanNotifyLendRequest) SetLoanEffectiveTime(loanEffectiveTime2 string) {
	this.LoanEffectiveTime = loanEffectiveTime2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetLoanEffectiveTime() string {
	return this.LoanEffectiveTime
}
func (this *OapiFinanceLoanNotifyLendRequest) SetLoanEndTime(loanEndTime2 string) {
	this.LoanEndTime = loanEndTime2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetLoanEndTime() string {
	return this.LoanEndTime
}
func (this *OapiFinanceLoanNotifyLendRequest) SetLoanOrderFlowNo(loanOrderFlowNo2 string) {
	this.LoanOrderFlowNo = loanOrderFlowNo2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetLoanOrderFlowNo() string {
	return this.LoanOrderFlowNo
}
func (this *OapiFinanceLoanNotifyLendRequest) SetLoanOrderNo(loanOrderNo2 string) {
	this.LoanOrderNo = loanOrderNo2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetLoanOrderNo() string {
	return this.LoanOrderNo
}
func (this *OapiFinanceLoanNotifyLendRequest) SetLoanSubmitTime(loanSubmitTime2 string) {
	this.LoanSubmitTime = loanSubmitTime2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetLoanSubmitTime() string {
	return this.LoanSubmitTime
}
func (this *OapiFinanceLoanNotifyLendRequest) SetLoanTxnTime(loanTxnTime2 string) {
	this.LoanTxnTime = loanTxnTime2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetLoanTxnTime() string {
	return this.LoanTxnTime
}
func (this *OapiFinanceLoanNotifyLendRequest) SetLoanUsage(loanUsage2 string) {
	this.LoanUsage = loanUsage2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetLoanUsage() string {
	return this.LoanUsage
}
func (this *OapiFinanceLoanNotifyLendRequest) SetOpenChannelName(openChannelName2 string) {
	this.OpenChannelName = openChannelName2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetOpenChannelName() string {
	return this.OpenChannelName
}
func (this *OapiFinanceLoanNotifyLendRequest) SetOpenProductCode(openProductCode2 string) {
	this.OpenProductCode = openProductCode2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetOpenProductCode() string {
	return this.OpenProductCode
}
func (this *OapiFinanceLoanNotifyLendRequest) SetOpenProductName(openProductName2 string) {
	this.OpenProductName = openProductName2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetOpenProductName() string {
	return this.OpenProductName
}
func (this *OapiFinanceLoanNotifyLendRequest) SetOpenProductType(openProductType2 string) {
	this.OpenProductType = openProductType2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetOpenProductType() string {
	return this.OpenProductType
}
func (this *OapiFinanceLoanNotifyLendRequest) SetPayableInterest(payableInterest2 int64) {
	this.PayableInterest = payableInterest2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetPayableInterest() int64 {
	return this.PayableInterest
}
func (this *OapiFinanceLoanNotifyLendRequest) SetPayablePrincipal(payablePrincipal2 int64) {
	this.PayablePrincipal = payablePrincipal2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetPayablePrincipal() int64 {
	return this.PayablePrincipal
}
func (this *OapiFinanceLoanNotifyLendRequest) SetPayableTotalAmount(payableTotalAmount2 int64) {
	this.PayableTotalAmount = payableTotalAmount2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetPayableTotalAmount() int64 {
	return this.PayableTotalAmount
}
func (this *OapiFinanceLoanNotifyLendRequest) SetReductionTotalAmount(reductionTotalAmount2 int64) {
	this.ReductionTotalAmount = reductionTotalAmount2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetReductionTotalAmount() int64 {
	return this.ReductionTotalAmount
}
func (this *OapiFinanceLoanNotifyLendRequest) SetRepayDate(repayDate2 int64) {
	this.RepayDate = repayDate2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetRepayDate() int64 {
	return this.RepayDate
}
func (this *OapiFinanceLoanNotifyLendRequest) SetRepayType(repayType2 string) {
	this.RepayType = repayType2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetRepayType() string {
	return this.RepayType
}
func (this *OapiFinanceLoanNotifyLendRequest) SetStatus(status2 string) {
	this.Status = status2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetStatus() string {
	return this.Status
}
func (this *OapiFinanceLoanNotifyLendRequest) SetTotalTerm(totalTerm2 int64) {
	this.TotalTerm = totalTerm2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetTotalTerm() int64 {
	return this.TotalTerm
}
func (this *OapiFinanceLoanNotifyLendRequest) SetUserMobile(userMobile2 string) {
	this.UserMobile = userMobile2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetUserMobile() string {
	return this.UserMobile
}
func (this *OapiFinanceLoanNotifyLendRequest) SetYearLoanInterestRate(yearLoanInterestRate2 string) {
	this.YearLoanInterestRate = yearLoanInterestRate2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetYearLoanInterestRate() string {
	return this.YearLoanInterestRate
}
func (this *OapiFinanceLoanNotifyLendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.loan.notify.lend"
}
func (this *OapiFinanceLoanNotifyLendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceLoanNotifyLendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceLoanNotifyLendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceLoanNotifyLendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceLoanNotifyLendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["available_limit"] = this.AvailableLimit
	txtParams["bank_name"] = this.BankName
	txtParams["bankcard_no"] = this.BankcardNo
	txtParams["bill_date"] = this.BillDate
	txtParams["bill_info_list"] = this.BillInfoList
	txtParams["credit_amount"] = this.CreditAmount
	txtParams["daily_interest_rate"] = this.DailyInterestRate
	txtParams["discounts_id"] = this.DiscountsId
	txtParams["fail_reason"] = this.FailReason
	txtParams["first_bill_date"] = this.FirstBillDate
	txtParams["first_repay_date"] = this.FirstRepayDate
	txtParams["id_card_no"] = this.IdCardNo
	txtParams["last_repay_date"] = this.LastRepayDate
	txtParams["loan_amount"] = this.LoanAmount
	txtParams["loan_effective_time"] = this.LoanEffectiveTime
	txtParams["loan_end_time"] = this.LoanEndTime
	txtParams["loan_order_flow_no"] = this.LoanOrderFlowNo
	txtParams["loan_order_no"] = this.LoanOrderNo
	txtParams["loan_submit_time"] = this.LoanSubmitTime
	txtParams["loan_txn_time"] = this.LoanTxnTime
	txtParams["loan_usage"] = this.LoanUsage
	txtParams["open_channel_name"] = this.OpenChannelName
	txtParams["open_product_code"] = this.OpenProductCode
	txtParams["open_product_name"] = this.OpenProductName
	txtParams["open_product_type"] = this.OpenProductType
	txtParams["payable_interest"] = this.PayableInterest
	txtParams["payable_principal"] = this.PayablePrincipal
	txtParams["payable_total_amount"] = this.PayableTotalAmount
	txtParams["reduction_total_amount"] = this.ReductionTotalAmount
	txtParams["repay_date"] = this.RepayDate
	txtParams["repay_type"] = this.RepayType
	txtParams["status"] = this.Status
	txtParams["total_term"] = this.TotalTerm
	txtParams["user_mobile"] = this.UserMobile
	txtParams["year_loan_interest_rate"] = this.YearLoanInterestRate
	return txtParams
}
func (this *OapiFinanceLoanNotifyLendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AvailableLimit, "availableLimit"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFinanceLoanNotifyLendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceLoanNotifyLendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type BillInfo struct {
	BillDate           string `json:"bill_date,omitempty"`
	PayableInterest    int64  `json:"payable_interest,omitempty"`
	PayablePrincipal   int64  `json:"payable_principal,omitempty"`
	PayableTotalAmount int64  `json:"payable_total_amount,omitempty"`
	RepayDate          string `json:"repay_date,omitempty"`
	Status             string `json:"status,omitempty"`
	Terms              string `json:"terms,omitempty"`
}
type OapiFinanceLoanNotifyLendResponse struct {
	taobao.TaobaoResponse
	Result  OpenCommonResult `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
