package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCrmObjectdataContactCreateRequest() *OapiCrmObjectdataContactCreateRequest {
	return &OapiCrmObjectdataContactCreateRequest{}
}

type OapiCrmObjectdataContactCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmObjectdataContactCreateResponse
	Instance        string
	ProviderCorpid  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmObjectdataContactCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectdataContactCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectdataContactCreateRequest) SetInstance(instance2 string) {
	this.Instance = instance2
}
func (this *OapiCrmObjectdataContactCreateRequest) GetInstance() string {
	return this.Instance
}
func (this *OapiCrmObjectdataContactCreateRequest) SetProviderCorpid(providerCorpid2 string) {
	this.ProviderCorpid = providerCorpid2
}
func (this *OapiCrmObjectdataContactCreateRequest) GetProviderCorpid() string {
	return this.ProviderCorpid
}
func (this *OapiCrmObjectdataContactCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectdata.contact.create"
}
func (this *OapiCrmObjectdataContactCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectdataContactCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectdataContactCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectdataContactCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectdataContactCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["instance"] = this.Instance
	txtParams["provider_corpid"] = this.ProviderCorpid
	return txtParams
}
func (this *OapiCrmObjectdataContactCreateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCrmObjectdataContactCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectdataContactCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type DataPermissionVo struct {
	OwnerUserids       []string `json:"owner_userids,omitempty"`
	ParticipantUserids []string `json:"participant_userids,omitempty"`
}
type ObjectDataInstanceVo struct {
	CreatorNick   string           `json:"creator_nick,omitempty"`
	CreatorUserid string           `json:"creator_userid,omitempty"`
	Data          string           `json:"data,omitempty"`
	ExtendData    string           `json:"extend_data,omitempty"`
	Permission    DataPermissionVo `json:"permission,omitempty"`
}
type OapiCrmObjectdataContactCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64               `json:"errcode,omitempty"`
	Errmsg  string              `json:"errmsg,omitempty"`
	Result  ObjectDataCreateDto `json:"result,omitempty"`
	Success bool                `json:"success,omitempty"`
}
type ObjectDataCreateDto struct {
	ContactUnionid string `json:"contact_unionid,omitempty"`
	ContactUserid  string `json:"contact_userid,omitempty"`
	InstanceId     string `json:"instance_id,omitempty"`
}
