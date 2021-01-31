package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMpdevPreviewbuildCreateRequest() *OapiMpdevPreviewbuildCreateRequest {
	return &OapiMpdevPreviewbuildCreateRequest{}
}

type OapiMpdevPreviewbuildCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                     OapiMpdevPreviewbuildCreateResponse
	BuildScriptVersion       string
	Corpid                   string
	EnableTabbar             string
	IgnoreHttpReqPermission  bool
	IgnoreWebviewDomainCheck bool
	MainPage                 string
	MiniappId                string
	PackageKey               string
	Page                     string
	PluginPackageKey         string
	PluginRefs               string
	Query                    string
	SubPackages              string
	TopHttpMethod            string
	TopResponseType          string
}

func (this *OapiMpdevPreviewbuildCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetBuildScriptVersion(buildScriptVersion2 string) {
	this.BuildScriptVersion = buildScriptVersion2
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetBuildScriptVersion() string {
	return this.BuildScriptVersion
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetCorpid(corpid2 string) {
	this.Corpid = corpid2
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetCorpid() string {
	return this.Corpid
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetEnableTabbar(enableTabbar2 string) {
	this.EnableTabbar = enableTabbar2
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetEnableTabbar() string {
	return this.EnableTabbar
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetIgnoreHttpReqPermission(ignoreHttpReqPermission2 bool) {
	this.IgnoreHttpReqPermission = ignoreHttpReqPermission2
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetIgnoreHttpReqPermission() bool {
	return this.IgnoreHttpReqPermission
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetIgnoreWebviewDomainCheck(ignoreWebviewDomainCheck2 bool) {
	this.IgnoreWebviewDomainCheck = ignoreWebviewDomainCheck2
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetIgnoreWebviewDomainCheck() bool {
	return this.IgnoreWebviewDomainCheck
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetMainPage(mainPage2 string) {
	this.MainPage = mainPage2
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetMainPage() string {
	return this.MainPage
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetMiniappId(miniappId2 string) {
	this.MiniappId = miniappId2
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetMiniappId() string {
	return this.MiniappId
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetPackageKey(packageKey2 string) {
	this.PackageKey = packageKey2
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetPackageKey() string {
	return this.PackageKey
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetPage(page2 string) {
	this.Page = page2
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetPage() string {
	return this.Page
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetPluginPackageKey(pluginPackageKey2 string) {
	this.PluginPackageKey = pluginPackageKey2
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetPluginPackageKey() string {
	return this.PluginPackageKey
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetPluginRefs(pluginRefs2 string) {
	this.PluginRefs = pluginRefs2
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetPluginRefs() string {
	return this.PluginRefs
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetQuery(query2 string) {
	this.Query = query2
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetQuery() string {
	return this.Query
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetSubPackages(subPackages2 string) {
	this.SubPackages = subPackages2
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetSubPackages() string {
	return this.SubPackages
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.mpdev.previewbuild.create"
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMpdevPreviewbuildCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["build_script_version"] = this.BuildScriptVersion
	txtParams["corpid"] = this.Corpid
	txtParams["enable_tabbar"] = this.EnableTabbar
	txtParams["ignore_http_req_permission"] = this.IgnoreHttpReqPermission
	txtParams["ignore_webview_domain_check"] = this.IgnoreWebviewDomainCheck
	txtParams["main_page"] = this.MainPage
	txtParams["miniapp_id"] = this.MiniappId
	txtParams["package_key"] = this.PackageKey
	txtParams["page"] = this.Page
	txtParams["plugin_package_key"] = this.PluginPackageKey
	txtParams["plugin_refs"] = this.PluginRefs
	txtParams["query"] = this.Query
	txtParams["sub_packages"] = this.SubPackages
	return txtParams
}
func (this *OapiMpdevPreviewbuildCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.EnableTabbar, "enableTabbar"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMpdevPreviewbuildCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenSubPackageVo struct {
	PackageKey string `json:"package_key,omitempty"`
	Path       string `json:"path,omitempty"`
	Type       string `json:"type,omitempty"`
}
type OapiMpdevPreviewbuildCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64         `json:"errcode,omitempty"`
	Errmsg  string        `json:"errmsg,omitempty"`
	Result  BuildResultVo `json:"result,omitempty"`
	Success bool          `json:"success,omitempty"`
}
