package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMicroappCustomUpdateRequest() *OapiMicroappCustomUpdateRequest {
	return &OapiMicroappCustomUpdateRequest{}
}

type OapiMicroappCustomUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp             OapiMicroappCustomUpdateResponse
	AgentId          int64
	AppCorpId        string
	Desc             string
	HomepageLink     string
	Icon             string
	IpWhiteList      string
	Name             string
	OmpLink          string
	PcHomepageLink   string
	TopHttpMethod    string
	TopRelatedCorpId string
	TopResponseType  string
}

func (this *OapiMicroappCustomUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappCustomUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMicroappCustomUpdateRequest) SetAgentId(agentId2 int64) {
	this.AgentId = agentId2
}
func (this *OapiMicroappCustomUpdateRequest) GetAgentId() int64 {
	return this.AgentId
}
func (this *OapiMicroappCustomUpdateRequest) SetAppCorpId(appCorpId2 string) {
	this.AppCorpId = appCorpId2
}
func (this *OapiMicroappCustomUpdateRequest) GetAppCorpId() string {
	return this.AppCorpId
}
func (this *OapiMicroappCustomUpdateRequest) SetDesc(desc2 string) {
	this.Desc = desc2
}
func (this *OapiMicroappCustomUpdateRequest) GetDesc() string {
	return this.Desc
}
func (this *OapiMicroappCustomUpdateRequest) SetHomepageLink(homepageLink2 string) {
	this.HomepageLink = homepageLink2
}
func (this *OapiMicroappCustomUpdateRequest) GetHomepageLink() string {
	return this.HomepageLink
}
func (this *OapiMicroappCustomUpdateRequest) SetIcon(icon2 string) {
	this.Icon = icon2
}
func (this *OapiMicroappCustomUpdateRequest) GetIcon() string {
	return this.Icon
}
func (this *OapiMicroappCustomUpdateRequest) SetIpWhiteList(ipWhiteList2 string) {
	this.IpWhiteList = ipWhiteList2
}
func (this *OapiMicroappCustomUpdateRequest) GetIpWhiteList() string {
	return this.IpWhiteList
}
func (this *OapiMicroappCustomUpdateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiMicroappCustomUpdateRequest) GetName() string {
	return this.Name
}
func (this *OapiMicroappCustomUpdateRequest) SetOmpLink(ompLink2 string) {
	this.OmpLink = ompLink2
}
func (this *OapiMicroappCustomUpdateRequest) GetOmpLink() string {
	return this.OmpLink
}
func (this *OapiMicroappCustomUpdateRequest) SetPcHomepageLink(pcHomepageLink2 string) {
	this.PcHomepageLink = pcHomepageLink2
}
func (this *OapiMicroappCustomUpdateRequest) GetPcHomepageLink() string {
	return this.PcHomepageLink
}
func (this *OapiMicroappCustomUpdateRequest) SetTopRelatedCorpId(topRelatedCorpId2 string) {
	this.TopRelatedCorpId = topRelatedCorpId2
}
func (this *OapiMicroappCustomUpdateRequest) GetTopRelatedCorpId() string {
	return this.TopRelatedCorpId
}
func (this *OapiMicroappCustomUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.custom.update"
}
func (this *OapiMicroappCustomUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappCustomUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappCustomUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappCustomUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappCustomUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agent_id"] = this.AgentId
	txtParams["app_corp_id"] = this.AppCorpId
	txtParams["desc"] = this.Desc
	txtParams["homepage_link"] = this.HomepageLink
	txtParams["icon"] = this.Icon
	txtParams["ip_white_list"] = this.IpWhiteList
	txtParams["name"] = this.Name
	txtParams["omp_link"] = this.OmpLink
	txtParams["pc_homepage_link"] = this.PcHomepageLink
	txtParams["top_related_corp_id"] = this.TopRelatedCorpId
	return txtParams
}
func (this *OapiMicroappCustomUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AgentId, "agentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMicroappCustomUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappCustomUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappCustomUpdateResponse struct {
	taobao.TaobaoResponse
}
