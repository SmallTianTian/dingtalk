package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMpdevBuildCreateRequest() *OapiMpdevBuildCreateRequest {
	return &OapiMpdevBuildCreateRequest{}
}

type OapiMpdevBuildCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMpdevBuildCreateResponse
	EnableTabbar    string
	MainPage        string
	MiniappId       string
	PackageKey      string
	PackageMd5      string
	PackageVersion  string
	PluginRefs      string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMpdevBuildCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMpdevBuildCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMpdevBuildCreateRequest) SetEnableTabbar(enableTabbar2 string) {
	this.EnableTabbar = enableTabbar2
}
func (this *OapiMpdevBuildCreateRequest) GetEnableTabbar() string {
	return this.EnableTabbar
}
func (this *OapiMpdevBuildCreateRequest) SetMainPage(mainPage2 string) {
	this.MainPage = mainPage2
}
func (this *OapiMpdevBuildCreateRequest) GetMainPage() string {
	return this.MainPage
}
func (this *OapiMpdevBuildCreateRequest) SetMiniappId(miniappId2 string) {
	this.MiniappId = miniappId2
}
func (this *OapiMpdevBuildCreateRequest) GetMiniappId() string {
	return this.MiniappId
}
func (this *OapiMpdevBuildCreateRequest) SetPackageKey(packageKey2 string) {
	this.PackageKey = packageKey2
}
func (this *OapiMpdevBuildCreateRequest) GetPackageKey() string {
	return this.PackageKey
}
func (this *OapiMpdevBuildCreateRequest) SetPackageMd5(packageMd52 string) {
	this.PackageMd5 = packageMd52
}
func (this *OapiMpdevBuildCreateRequest) GetPackageMd5() string {
	return this.PackageMd5
}
func (this *OapiMpdevBuildCreateRequest) SetPackageVersion(packageVersion2 string) {
	this.PackageVersion = packageVersion2
}
func (this *OapiMpdevBuildCreateRequest) GetPackageVersion() string {
	return this.PackageVersion
}
func (this *OapiMpdevBuildCreateRequest) SetPluginRefs(pluginRefs2 string) {
	this.PluginRefs = pluginRefs2
}
func (this *OapiMpdevBuildCreateRequest) GetPluginRefs() string {
	return this.PluginRefs
}
func (this *OapiMpdevBuildCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.mpdev.build.create"
}
func (this *OapiMpdevBuildCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMpdevBuildCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMpdevBuildCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMpdevBuildCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMpdevBuildCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["enable_tabbar"] = this.EnableTabbar
	txtParams["main_page"] = this.MainPage
	txtParams["miniapp_id"] = this.MiniappId
	txtParams["package_key"] = this.PackageKey
	txtParams["package_md5"] = this.PackageMd5
	txtParams["package_version"] = this.PackageVersion
	txtParams["plugin_refs"] = this.PluginRefs
	return txtParams
}
func (this *OapiMpdevBuildCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.EnableTabbar, "enableTabbar"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMpdevBuildCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMpdevBuildCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type PluginReference struct {
	PluginId      string `json:"plugin_id,omitempty"`
	PluginVersion string `json:"plugin_version,omitempty"`
}
type OapiMpdevBuildCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64         `json:"errcode,omitempty"`
	Errmsg  string        `json:"errmsg,omitempty"`
	Result  BuildResultVo `json:"result,omitempty"`
	Success bool          `json:"success,omitempty"`
}
type BuildResultVo struct {
	BuildId        int64  `json:"build_id,omitempty"`
	BuildInfo      string `json:"build_info,omitempty"`
	Finished       bool   `json:"finished,omitempty"`
	LogUrl         string `json:"log_url,omitempty"`
	ResultUrl      string `json:"result_url,omitempty"`
	Status         string `json:"status,omitempty"`
	TaskId         string `json:"task_id,omitempty"`
	Version        string `json:"version,omitempty"`
	VersionCreated bool   `json:"version_created,omitempty"`
}
