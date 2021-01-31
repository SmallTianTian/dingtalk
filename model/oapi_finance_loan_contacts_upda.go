package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiFinanceLoanContactsUpdateRequest() *OapiFinanceLoanContactsUpdateRequest {
	return &OapiFinanceLoanContactsUpdateRequest{}
}

type OapiFinanceLoanContactsUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFinanceLoanContactsUpdateResponse
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

func (this *OapiFinanceLoanContactsUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFinanceLoanContactsUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFinanceLoanContactsUpdateRequest) SetContactId(contactId2 int64) {
	this.ContactId = contactId2
}
func (this *OapiFinanceLoanContactsUpdateRequest) GetContactId() int64 {
	return this.ContactId
}
func (this *OapiFinanceLoanContactsUpdateRequest) SetContactMobile(contactMobile2 string) {
	this.ContactMobile = contactMobile2
}
func (this *OapiFinanceLoanContactsUpdateRequest) GetContactMobile() string {
	return this.ContactMobile
}
func (this *OapiFinanceLoanContactsUpdateRequest) SetIdCardNo(idCardNo2 string) {
	this.IdCardNo = idCardNo2
}
func (this *OapiFinanceLoanContactsUpdateRequest) GetIdCardNo() string {
	return this.IdCardNo
}
func (this *OapiFinanceLoanContactsUpdateRequest) SetUserCategory(userCategory2 string) {
	this.UserCategory = userCategory2
}
func (this *OapiFinanceLoanContactsUpdateRequest) GetUserCategory() string {
	return this.UserCategory
}
func (this *OapiFinanceLoanContactsUpdateRequest) SetUserMobile(userMobile2 string) {
	this.UserMobile = userMobile2
}
func (this *OapiFinanceLoanContactsUpdateRequest) GetUserMobile() string {
	return this.UserMobile
}
func (this *OapiFinanceLoanContactsUpdateRequest) SetUserName(userName2 string) {
	this.UserName = userName2
}
func (this *OapiFinanceLoanContactsUpdateRequest) GetUserName() string {
	return this.UserName
}
func (this *OapiFinanceLoanContactsUpdateRequest) SetUserRelation(userRelation2 string) {
	this.UserRelation = userRelation2
}
func (this *OapiFinanceLoanContactsUpdateRequest) GetUserRelation() string {
	return this.UserRelation
}
func (this *OapiFinanceLoanContactsUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.finance.loan.contacts.update"
}
func (this *OapiFinanceLoanContactsUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFinanceLoanContactsUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFinanceLoanContactsUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFinanceLoanContactsUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFinanceLoanContactsUpdateRequest) GetTextParams() map[string]interface{} {
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
func (this *OapiFinanceLoanContactsUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ContactId, "contactId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiFinanceLoanContactsUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFinanceLoanContactsUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFinanceLoanContactsUpdateResponse struct {
	taobao.TaobaoResponse
	Result  OpenCommonResult `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
