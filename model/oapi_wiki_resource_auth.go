package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWikiResourceAuthRequest() *OapiWikiResourceAuthRequest {
	return &OapiWikiResourceAuthRequest{}
}

type OapiWikiResourceAuthRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWikiResourceAuthResponse
	Agentid         int64
	AuthCode        string
	IsPublic        bool
	ResourceList    string
	ResourceType    int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWikiResourceAuthRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWikiResourceAuthRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWikiResourceAuthRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiWikiResourceAuthRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiWikiResourceAuthRequest) SetAuthCode(authCode2 string) {
	this.AuthCode = authCode2
}
func (this *OapiWikiResourceAuthRequest) GetAuthCode() string {
	return this.AuthCode
}
func (this *OapiWikiResourceAuthRequest) SetIsPublic(isPublic2 bool) {
	this.IsPublic = isPublic2
}
func (this *OapiWikiResourceAuthRequest) GetIsPublic() bool {
	return this.IsPublic
}
func (this *OapiWikiResourceAuthRequest) SetResourceList(resourceList2 string) {
	this.ResourceList = resourceList2
}
func (this *OapiWikiResourceAuthRequest) GetResourceList() string {
	return this.ResourceList
}
func (this *OapiWikiResourceAuthRequest) SetResourceType(resourceType2 int64) {
	this.ResourceType = resourceType2
}
func (this *OapiWikiResourceAuthRequest) GetResourceType() int64 {
	return this.ResourceType
}
func (this *OapiWikiResourceAuthRequest) GetApiMethodName() string {
	return "dingtalk.oapi.wiki.resource.auth"
}
func (this *OapiWikiResourceAuthRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWikiResourceAuthRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWikiResourceAuthRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWikiResourceAuthRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWikiResourceAuthRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["auth_code"] = this.AuthCode
	txtParams["is_public"] = this.IsPublic
	txtParams["resource_list"] = this.ResourceList
	txtParams["resource_type"] = this.ResourceType
	return txtParams
}
func (this *OapiWikiResourceAuthRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWikiResourceAuthRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWikiResourceAuthRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWikiResourceAuthResponse struct {
	taobao.TaobaoResponse
	Result  []OpenResourceVo `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
type OpenResourceVo struct {
	ResourceId   string `json:"resource_id,omitempty"`
	ResourceType int64  `json:"resource_type,omitempty"`
	Status       int64  `json:"status,omitempty"`
}
