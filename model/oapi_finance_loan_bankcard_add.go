package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFinanceLoanBankcardAddRequest() *OapiFinanceLoanBankcardAddRequest {
	return &OapiFinanceLoanBankcardAddRequest{}
}

type OapiFinanceLoanBankcardAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFinanceLoanBankcardAddResponse
	BankCode        string
	BankName        string
	BankcardMobile  string
	BankcardNo      string
	IdCardNo        string
	TopHttpMethod   string
	TopResponseType string
	UserMobile      string
}

func (this *OapiFinanceLoanBankcardAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceLoanBankcardAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceLoanBankcardAddRequest) SetBankCode(bankCode2 string) {
	this.BankCode = bankCode2
}
func (this *OapiFinanceLoanBankcardAddRequest) GetBankCode() string {
	return this.BankCode
}
func (this *OapiFinanceLoanBankcardAddRequest) SetBankName(bankName2 string) {
	this.BankName = bankName2
}
func (this *OapiFinanceLoanBankcardAddRequest) GetBankName() string {
	return this.BankName
}
func (this *OapiFinanceLoanBankcardAddRequest) SetBankcardMobile(bankcardMobile2 string) {
	this.BankcardMobile = bankcardMobile2
}
func (this *OapiFinanceLoanBankcardAddRequest) GetBankcardMobile() string {
	return this.BankcardMobile
}
func (this *OapiFinanceLoanBankcardAddRequest) SetBankcardNo(bankcardNo2 string) {
	this.BankcardNo = bankcardNo2
}
func (this *OapiFinanceLoanBankcardAddRequest) GetBankcardNo() string {
	return this.BankcardNo
}
func (this *OapiFinanceLoanBankcardAddRequest) SetIdCardNo(idCardNo2 string) {
	this.IdCardNo = idCardNo2
}
func (this *OapiFinanceLoanBankcardAddRequest) GetIdCardNo() string {
	return this.IdCardNo
}
func (this *OapiFinanceLoanBankcardAddRequest) SetUserMobile(userMobile2 string) {
	this.UserMobile = userMobile2
}
func (this *OapiFinanceLoanBankcardAddRequest) GetUserMobile() string {
	return this.UserMobile
}
func (this *OapiFinanceLoanBankcardAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.loan.bankcard.add"
}
func (this *OapiFinanceLoanBankcardAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceLoanBankcardAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceLoanBankcardAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceLoanBankcardAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceLoanBankcardAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["bank_code"] = this.BankCode
	txtParams["bank_name"] = this.BankName
	txtParams["bankcard_mobile"] = this.BankcardMobile
	txtParams["bankcard_no"] = this.BankcardNo
	txtParams["id_card_no"] = this.IdCardNo
	txtParams["user_mobile"] = this.UserMobile
	return txtParams
}
func (this *OapiFinanceLoanBankcardAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BankCode, "bankCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFinanceLoanBankcardAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceLoanBankcardAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFinanceLoanBankcardAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64            `json:"errcode,omitempty"`
	Errmsg  string           `json:"errmsg,omitempty"`
	Result  OpenCommonResult `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
type OpenCommonResult struct {
	Result string `json:"result,omitempty"`
}
