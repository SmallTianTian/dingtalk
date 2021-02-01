package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFinanceLoanContactsDeleteRequest() *OapiFinanceLoanContactsDeleteRequest {
	return &OapiFinanceLoanContactsDeleteRequest{}
}

type OapiFinanceLoanContactsDeleteRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFinanceLoanContactsDeleteResponse
	ContactId       int64
	ContactMobile   string
	IdCardNo        string
	TopHttpMethod   string
	TopResponseType string
	UserCategory    string
	UserMobile      string
	UserName        string
	UserRelation    string
}

func (this *OapiFinanceLoanContactsDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceLoanContactsDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceLoanContactsDeleteRequest) SetContactId(contactId2 int64) {
	this.ContactId = contactId2
}
func (this *OapiFinanceLoanContactsDeleteRequest) GetContactId() int64 {
	return this.ContactId
}
func (this *OapiFinanceLoanContactsDeleteRequest) SetContactMobile(contactMobile2 string) {
	this.ContactMobile = contactMobile2
}
func (this *OapiFinanceLoanContactsDeleteRequest) GetContactMobile() string {
	return this.ContactMobile
}
func (this *OapiFinanceLoanContactsDeleteRequest) SetIdCardNo(idCardNo2 string) {
	this.IdCardNo = idCardNo2
}
func (this *OapiFinanceLoanContactsDeleteRequest) GetIdCardNo() string {
	return this.IdCardNo
}
func (this *OapiFinanceLoanContactsDeleteRequest) SetUserCategory(userCategory2 string) {
	this.UserCategory = userCategory2
}
func (this *OapiFinanceLoanContactsDeleteRequest) GetUserCategory() string {
	return this.UserCategory
}
func (this *OapiFinanceLoanContactsDeleteRequest) SetUserMobile(userMobile2 string) {
	this.UserMobile = userMobile2
}
func (this *OapiFinanceLoanContactsDeleteRequest) GetUserMobile() string {
	return this.UserMobile
}
func (this *OapiFinanceLoanContactsDeleteRequest) SetUserName(userName2 string) {
	this.UserName = userName2
}
func (this *OapiFinanceLoanContactsDeleteRequest) GetUserName() string {
	return this.UserName
}
func (this *OapiFinanceLoanContactsDeleteRequest) SetUserRelation(userRelation2 string) {
	this.UserRelation = userRelation2
}
func (this *OapiFinanceLoanContactsDeleteRequest) GetUserRelation() string {
	return this.UserRelation
}
func (this *OapiFinanceLoanContactsDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.loan.contacts.delete"
}
func (this *OapiFinanceLoanContactsDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceLoanContactsDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceLoanContactsDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceLoanContactsDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceLoanContactsDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["contact_id"] = this.ContactId
	txtParams["contact_mobile"] = this.ContactMobile
	txtParams["id_card_no"] = this.IdCardNo
	txtParams["user_category"] = this.UserCategory
	txtParams["user_mobile"] = this.UserMobile
	txtParams["user_name"] = this.UserName
	txtParams["user_relation"] = this.UserRelation
	return txtParams
}
func (this *OapiFinanceLoanContactsDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ContactId, "contactId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFinanceLoanContactsDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceLoanContactsDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFinanceLoanContactsDeleteResponse struct {
	taobao.TaobaoResponse
	Result  OpenCommonResult `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
