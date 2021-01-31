package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFinanceLoanBankcardUpdateRequest() *OapiFinanceLoanBankcardUpdateRequest {
	return &OapiFinanceLoanBankcardUpdateRequest{}
}

type OapiFinanceLoanBankcardUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFinanceLoanBankcardUpdateResponse
	BankCode        string
	BankName        string
	BankcardId      int64
	BankcardMobile  string
	BankcardNo      string
	IdCardNo        string
	TopHttpMethod   string
	TopResponseType string
	UserMobile      string
}

func (this *OapiFinanceLoanBankcardUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceLoanBankcardUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceLoanBankcardUpdateRequest) SetBankCode(bankCode2 string) {
	this.BankCode = bankCode2
}
func (this *OapiFinanceLoanBankcardUpdateRequest) GetBankCode() string {
	return this.BankCode
}
func (this *OapiFinanceLoanBankcardUpdateRequest) SetBankName(bankName2 string) {
	this.BankName = bankName2
}
func (this *OapiFinanceLoanBankcardUpdateRequest) GetBankName() string {
	return this.BankName
}
func (this *OapiFinanceLoanBankcardUpdateRequest) SetBankcardId(bankcardId2 int64) {
	this.BankcardId = bankcardId2
}
func (this *OapiFinanceLoanBankcardUpdateRequest) GetBankcardId() int64 {
	return this.BankcardId
}
func (this *OapiFinanceLoanBankcardUpdateRequest) SetBankcardMobile(bankcardMobile2 string) {
	this.BankcardMobile = bankcardMobile2
}
func (this *OapiFinanceLoanBankcardUpdateRequest) GetBankcardMobile() string {
	return this.BankcardMobile
}
func (this *OapiFinanceLoanBankcardUpdateRequest) SetBankcardNo(bankcardNo2 string) {
	this.BankcardNo = bankcardNo2
}
func (this *OapiFinanceLoanBankcardUpdateRequest) GetBankcardNo() string {
	return this.BankcardNo
}
func (this *OapiFinanceLoanBankcardUpdateRequest) SetIdCardNo(idCardNo2 string) {
	this.IdCardNo = idCardNo2
}
func (this *OapiFinanceLoanBankcardUpdateRequest) GetIdCardNo() string {
	return this.IdCardNo
}
func (this *OapiFinanceLoanBankcardUpdateRequest) SetUserMobile(userMobile2 string) {
	this.UserMobile = userMobile2
}
func (this *OapiFinanceLoanBankcardUpdateRequest) GetUserMobile() string {
	return this.UserMobile
}
func (this *OapiFinanceLoanBankcardUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.loan.bankcard.update"
}
func (this *OapiFinanceLoanBankcardUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceLoanBankcardUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceLoanBankcardUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceLoanBankcardUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceLoanBankcardUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["bank_code"] = this.BankCode
	txtParams["bank_name"] = this.BankName
	txtParams["bankcard_id"] = this.BankcardId
	txtParams["bankcard_mobile"] = this.BankcardMobile
	txtParams["bankcard_no"] = this.BankcardNo
	txtParams["id_card_no"] = this.IdCardNo
	txtParams["user_mobile"] = this.UserMobile
	return txtParams
}
func (this *OapiFinanceLoanBankcardUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BankCode, "bankCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFinanceLoanBankcardUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceLoanBankcardUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFinanceLoanBankcardUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64            `json:"errcode,omitempty"`
	Errmsg  string           `json:"errmsg,omitempty"`
	Result  OpenCommonResult `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
