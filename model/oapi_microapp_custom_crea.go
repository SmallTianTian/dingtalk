package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMicroappCustomCreateRequest() *OapiMicroappCustomCreateRequest {
	return &OapiMicroappCustomCreateRequest{}
}

type OapiMicroappCustomCreateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp             OapiMicroappCustomCreateResponse
	AppCorpId        string
	Desc             string
	DevelopType      int64
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

func (this *OapiMicroappCustomCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMicroappCustomCreateRequest) SetAppCorpId(appCorpId2 string) {
	this.AppCorpId = appCorpId2
}
func (this *OapiMicroappCustomCreateRequest) GetAppCorpId() string {
	return this.AppCorpId
}
func (this *OapiMicroappCustomCreateRequest) SetDesc(desc2 string) {
	this.Desc = desc2
}
func (this *OapiMicroappCustomCreateRequest) GetDesc() string {
	return this.Desc
}
func (this *OapiMicroappCustomCreateRequest) SetDevelopType(developType2 int64) {
	this.DevelopType = developType2
}
func (this *OapiMicroappCustomCreateRequest) GetDevelopType() int64 {
	return this.DevelopType
}
func (this *OapiMicroappCustomCreateRequest) SetHomepageLink(homepageLink2 string) {
	this.HomepageLink = homepageLink2
}
func (this *OapiMicroappCustomCreateRequest) GetHomepageLink() string {
	return this.HomepageLink
}
func (this *OapiMicroappCustomCreateRequest) SetIcon(icon2 string) {
	this.Icon = icon2
}
func (this *OapiMicroappCustomCreateRequest) GetIcon() string {
	return this.Icon
}
func (this *OapiMicroappCustomCreateRequest) SetIpWhiteList(ipWhiteList2 string) {
	this.IpWhiteList = ipWhiteList2
}
func (this *OapiMicroappCustomCreateRequest) GetIpWhiteList() string {
	return this.IpWhiteList
}
func (this *OapiMicroappCustomCreateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiMicroappCustomCreateRequest) GetName() string {
	return this.Name
}
func (this *OapiMicroappCustomCreateRequest) SetOmpLink(ompLink2 string) {
	this.OmpLink = ompLink2
}
func (this *OapiMicroappCustomCreateRequest) GetOmpLink() string {
	return this.OmpLink
}
func (this *OapiMicroappCustomCreateRequest) SetPcHomepageLink(pcHomepageLink2 string) {
	this.PcHomepageLink = pcHomepageLink2
}
func (this *OapiMicroappCustomCreateRequest) GetPcHomepageLink() string {
	return this.PcHomepageLink
}
func (this *OapiMicroappCustomCreateRequest) SetTopRelatedCorpId(topRelatedCorpId2 string) {
	this.TopRelatedCorpId = topRelatedCorpId2
}
func (this *OapiMicroappCustomCreateRequest) GetTopRelatedCorpId() string {
	return this.TopRelatedCorpId
}
func (this *OapiMicroappCustomCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.microapp.custom.create"
}
func (this *OapiMicroappCustomCreateRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *OapiMicroappCustomCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMicroappCustomCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMicroappCustomCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMicroappCustomCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMicroappCustomCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["app_corp_id"] = this.AppCorpId
	txtParams["desc"] = this.Desc
	txtParams["develop_type"] = this.DevelopType
	txtParams["homepage_link"] = this.HomepageLink
	txtParams["icon"] = this.Icon
	txtParams["ip_white_list"] = this.IpWhiteList
	txtParams["name"] = this.Name
	txtParams["omp_link"] = this.OmpLink
	txtParams["pc_homepage_link"] = this.PcHomepageLink
	txtParams["top_related_corp_id"] = this.TopRelatedCorpId
	return txtParams
}
func (this *OapiMicroappCustomCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AppCorpId, "appCorpId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMicroappCustomCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMicroappCustomCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMicroappCustomCreateResponse struct {
	taobao.TaobaoResponse
	Result CustomAppCreateResponseVo `json:"result,omitempty"`
}
type CustomAppCreateResponseVo struct {
	AgentId      int64  `json:"agent_id,omitempty"`
	CustomKey    string `json:"custom_key,omitempty"`
	CustomSecret string `json:"custom_secret,omitempty"`
}
