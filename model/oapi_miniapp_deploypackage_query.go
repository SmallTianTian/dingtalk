package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMiniappDeploypackageQueryRequest() *OapiMiniappDeploypackageQueryRequest {
	return &OapiMiniappDeploypackageQueryRequest{}
}

type OapiMiniappDeploypackageQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMiniappDeploypackageQueryResponse
	ModelKey        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMiniappDeploypackageQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMiniappDeploypackageQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMiniappDeploypackageQueryRequest) SetModelKey(modelKey2 string) {
	this.ModelKey = modelKey2
}
func (this *OapiMiniappDeploypackageQueryRequest) GetModelKey() string {
	return this.ModelKey
}
func (this *OapiMiniappDeploypackageQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.miniapp.deploypackage.query"
}
func (this *OapiMiniappDeploypackageQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMiniappDeploypackageQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMiniappDeploypackageQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMiniappDeploypackageQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMiniappDeploypackageQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["model_key"] = this.ModelKey
	return txtParams
}
func (this *OapiMiniappDeploypackageQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ModelKey, "modelKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMiniappDeploypackageQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMiniappDeploypackageQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMiniappDeploypackageQueryResponse struct {
	taobao.TaobaoResponse
	Data    DeployPackageDoModel `json:"data,omitempty"`
	Errcode int64                `json:"errcode,omitempty"`
	Errmsg  string               `json:"errmsg,omitempty"`
}
type DeployPackageDoModel struct {
	AppId           string `json:"app_id,omitempty"`
	AutoInstall     int64  `json:"auto_install,omitempty"`
	ClientId        int64  `json:"client_id,omitempty"`
	DeployTime      int64  `json:"deploy_time,omitempty"`
	Desc            string `json:"desc,omitempty"`
	EnglishName     string `json:"english_name,omitempty"`
	ExtendInfo      string `json:"extend_info,omitempty"`
	FallbackBaseUrl string `json:"fallback_base_url,omitempty"`
	GmtCreate       int64  `json:"gmt_create,omitempty"`
	GmtModified     int64  `json:"gmt_modified,omitempty"`
	Gray            int64  `json:"gray,omitempty"`
	GrayCode        string `json:"gray_code,omitempty"`
	GrayTime        int64  `json:"gray_time,omitempty"`
	IconUrl         string `json:"icon_url,omitempty"`
	Id              int64  `json:"id,omitempty"`
	InstId          int64  `json:"inst_id,omitempty"`
	IsDeleted       int64  `json:"is_deleted,omitempty"`
	MainUrl         string `json:"main_url,omitempty"`
	Name            string `json:"name,omitempty"`
	Online          int64  `json:"online,omitempty"`
	PackageId       int64  `json:"package_id,omitempty"`
	PackageUrl      string `json:"package_url,omitempty"`
	PluginRefs      string `json:"plugin_refs,omitempty"`
	PluginSize      int64  `json:"plugin_size,omitempty"`
	PluginUrl       string `json:"plugin_url,omitempty"`
	Pre             int64  `json:"pre,omitempty"`
	PreTime         int64  `json:"pre_time,omitempty"`
	Preset          int64  `json:"preset,omitempty"`
	Prod            int64  `json:"prod,omitempty"`
	RollbackFrom    int64  `json:"rollback_from,omitempty"`
	Size            int64  `json:"size,omitempty"`
	Slogan          string `json:"slogan,omitempty"`
	Version         string `json:"version,omitempty"`
	Vhost           string `json:"vhost,omitempty"`
	WindowId        int64  `json:"window_id,omitempty"`
}
