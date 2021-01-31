package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFinanceLoanContactsAddRequest() *OapiFinanceLoanContactsAddRequest {
	return &OapiFinanceLoanContactsAddRequest{}
}

type OapiFinanceLoanContactsAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFinanceLoanContactsAddResponse
	ContactMobile   string
	IdCardNo        string
	TopHttpMethod   string
	TopResponseType string
	UserCategory    string
	UserMobile      string
	UserName        string
	UserRelation    string
}

func (this *OapiFinanceLoanContactsAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceLoanContactsAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceLoanContactsAddRequest) SetContactMobile(contactMobile2 string) {
	this.ContactMobile = contactMobile2
}
func (this *OapiFinanceLoanContactsAddRequest) GetContactMobile() string {
	return this.ContactMobile
}
func (this *OapiFinanceLoanContactsAddRequest) SetIdCardNo(idCardNo2 string) {
	this.IdCardNo = idCardNo2
}
func (this *OapiFinanceLoanContactsAddRequest) GetIdCardNo() string {
	return this.IdCardNo
}
func (this *OapiFinanceLoanContactsAddRequest) SetUserCategory(userCategory2 string) {
	this.UserCategory = userCategory2
}
func (this *OapiFinanceLoanContactsAddRequest) GetUserCategory() string {
	return this.UserCategory
}
func (this *OapiFinanceLoanContactsAddRequest) SetUserMobile(userMobile2 string) {
	this.UserMobile = userMobile2
}
func (this *OapiFinanceLoanContactsAddRequest) GetUserMobile() string {
	return this.UserMobile
}
func (this *OapiFinanceLoanContactsAddRequest) SetUserName(userName2 string) {
	this.UserName = userName2
}
func (this *OapiFinanceLoanContactsAddRequest) GetUserName() string {
	return this.UserName
}
func (this *OapiFinanceLoanContactsAddRequest) SetUserRelation(userRelation2 string) {
	this.UserRelation = userRelation2
}
func (this *OapiFinanceLoanContactsAddRequest) GetUserRelation() string {
	return this.UserRelation
}
func (this *OapiFinanceLoanContactsAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.loan.contacts.add"
}
func (this *OapiFinanceLoanContactsAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceLoanContactsAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceLoanContactsAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceLoanContactsAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceLoanContactsAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["contact_mobile"] = this.ContactMobile
	txtParams["id_card_no"] = this.IdCardNo
	txtParams["user_category"] = this.UserCategory
	txtParams["user_mobile"] = this.UserMobile
	txtParams["user_name"] = this.UserName
	txtParams["user_relation"] = this.UserRelation
	return txtParams
}
func (this *OapiFinanceLoanContactsAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ContactMobile, "contactMobile"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFinanceLoanContactsAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceLoanContactsAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFinanceLoanContactsAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64            `json:"errcode,omitempty"`
	Errmsg  string           `json:"errmsg,omitempty"`
	Result  OpenCommonResult `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
