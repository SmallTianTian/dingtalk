package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmContactCreateRequest() *OapiCrmContactCreateRequest {
	return &OapiCrmContactCreateRequest{}
}

type OapiCrmContactCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                OapiCrmContactCreateResponse
	ContactName         string
	ContactPhone        string
	ContactPositionList string
	CreatorUserid       string
	CustomerInstanceId  string
	ProviderCorpid      string
	TopHttpMethod       string
	TopResponseType     string
}

func (this *OapiCrmContactCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmContactCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmContactCreateRequest) SetContactName(contactName2 string) {
	this.ContactName = contactName2
}
func (this *OapiCrmContactCreateRequest) GetContactName() string {
	return this.ContactName
}
func (this *OapiCrmContactCreateRequest) SetContactPhone(contactPhone2 string) {
	this.ContactPhone = contactPhone2
}
func (this *OapiCrmContactCreateRequest) GetContactPhone() string {
	return this.ContactPhone
}
func (this *OapiCrmContactCreateRequest) SetContactPositionList(contactPositionList2 string) {
	this.ContactPositionList = contactPositionList2
}
func (this *OapiCrmContactCreateRequest) GetContactPositionList() string {
	return this.ContactPositionList
}
func (this *OapiCrmContactCreateRequest) SetCreatorUserid(creatorUserid2 string) {
	this.CreatorUserid = creatorUserid2
}
func (this *OapiCrmContactCreateRequest) GetCreatorUserid() string {
	return this.CreatorUserid
}
func (this *OapiCrmContactCreateRequest) SetCustomerInstanceId(customerInstanceId2 string) {
	this.CustomerInstanceId = customerInstanceId2
}
func (this *OapiCrmContactCreateRequest) GetCustomerInstanceId() string {
	return this.CustomerInstanceId
}
func (this *OapiCrmContactCreateRequest) SetProviderCorpid(providerCorpid2 string) {
	this.ProviderCorpid = providerCorpid2
}
func (this *OapiCrmContactCreateRequest) GetProviderCorpid() string {
	return this.ProviderCorpid
}
func (this *OapiCrmContactCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.contact.create"
}
func (this *OapiCrmContactCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmContactCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmContactCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmContactCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmContactCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["contact_name"] = this.ContactName
	txtParams["contact_phone"] = this.ContactPhone
	txtParams["contact_position_list"] = this.ContactPositionList
	txtParams["creator_userid"] = this.CreatorUserid
	txtParams["customer_instance_id"] = this.CustomerInstanceId
	txtParams["provider_corpid"] = this.ProviderCorpid
	return txtParams
}
func (this *OapiCrmContactCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ContactName, "contactName"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmContactCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmContactCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmContactCreateResponse struct {
	taobao.TaobaoResponse
	Result CreateContactResponse `json:"result,omitempty"`
}
type CreateContactResponse struct {
	ContactInstanceId string `json:"contact_instance_id,omitempty"`
	ContactUnionid    string `json:"contact_unionid,omitempty"`
	ContactUserid     string `json:"contact_userid,omitempty"`
}
