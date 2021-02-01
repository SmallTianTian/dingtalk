package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmGroupCreateRequest() *OapiCrmGroupCreateRequest {
	return &OapiCrmGroupCreateRequest{}
}

type OapiCrmGroupCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp                OapiCrmGroupCreateResponse
	ColleagueUseridList string
	ContactIdList       string
	CustomerCorpid      string
	CustomerId          string
	GroupOwner          string
	TopHttpMethod       string
	TopResponseType     string
}

func (this *OapiCrmGroupCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmGroupCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmGroupCreateRequest) SetColleagueUseridList(colleagueUseridList2 string) {
	this.ColleagueUseridList = colleagueUseridList2
}
func (this *OapiCrmGroupCreateRequest) GetColleagueUseridList() string {
	return this.ColleagueUseridList
}
func (this *OapiCrmGroupCreateRequest) SetContactIdList(contactIdList2 string) {
	this.ContactIdList = contactIdList2
}
func (this *OapiCrmGroupCreateRequest) GetContactIdList() string {
	return this.ContactIdList
}
func (this *OapiCrmGroupCreateRequest) SetCustomerCorpid(customerCorpid2 string) {
	this.CustomerCorpid = customerCorpid2
}
func (this *OapiCrmGroupCreateRequest) GetCustomerCorpid() string {
	return this.CustomerCorpid
}
func (this *OapiCrmGroupCreateRequest) SetCustomerId(customerId2 string) {
	this.CustomerId = customerId2
}
func (this *OapiCrmGroupCreateRequest) GetCustomerId() string {
	return this.CustomerId
}
func (this *OapiCrmGroupCreateRequest) SetGroupOwner(groupOwner2 string) {
	this.GroupOwner = groupOwner2
}
func (this *OapiCrmGroupCreateRequest) GetGroupOwner() string {
	return this.GroupOwner
}
func (this *OapiCrmGroupCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.group.create"
}
func (this *OapiCrmGroupCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmGroupCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmGroupCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmGroupCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmGroupCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["colleague_userid_list"] = this.ColleagueUseridList
	txtParams["contact_id_list"] = this.ContactIdList
	txtParams["customer_corpid"] = this.CustomerCorpid
	txtParams["customer_id"] = this.CustomerId
	txtParams["group_owner"] = this.GroupOwner
	return txtParams
}
func (this *OapiCrmGroupCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.ColleagueUseridList, 20, "colleagueUseridList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmGroupCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmGroupCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmGroupCreateResponse struct {
	taobao.TaobaoResponse
	Result CreateGroupResponse `json:"result,omitempty"`
}
type CreateGroupResponse struct {
	Cid string `json:"cid,omitempty"`
}
