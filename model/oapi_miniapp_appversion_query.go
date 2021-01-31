package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMiniappAppversionQueryRequest() *OapiMiniappAppversionQueryRequest {
	return &OapiMiniappAppversionQueryRequest{}
}

type OapiMiniappAppversionQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMiniappAppversionQueryResponse
	ModelKey        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMiniappAppversionQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMiniappAppversionQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMiniappAppversionQueryRequest) SetModelKey(modelKey2 string) {
	this.ModelKey = modelKey2
}
func (this *OapiMiniappAppversionQueryRequest) GetModelKey() string {
	return this.ModelKey
}
func (this *OapiMiniappAppversionQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.miniapp.appversion.query"
}
func (this *OapiMiniappAppversionQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMiniappAppversionQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMiniappAppversionQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMiniappAppversionQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMiniappAppversionQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["model_key"] = this.ModelKey
	return txtParams
}
func (this *OapiMiniappAppversionQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ModelKey, "modelKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMiniappAppversionQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMiniappAppversionQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMiniappAppversionQueryResponse struct {
	taobao.TaobaoResponse
	Data    AppVersionDoModel `json:"data,omitempty"`
	Errcode int64             `json:"errcode,omitempty"`
	Errmsg  string            `json:"errmsg,omitempty"`
}
type AppVersionDoModel struct {
	AppId            string `json:"app_id,omitempty"`
	BuildTaskNo      string `json:"build_task_no,omitempty"`
	BuildType        int64  `json:"build_type,omitempty"`
	ClientId         int64  `json:"client_id,omitempty"`
	ExtJsonConfig    string `json:"ext_json_config,omitempty"`
	ExtJsonConfigUrl string `json:"ext_json_config_url,omitempty"`
	ExtendInfo       string `json:"extend_info,omitempty"`
	Extra            string `json:"extra,omitempty"`
	FallbackBaseUrl  string `json:"fallback_base_url,omitempty"`
	GmtCreate        int64  `json:"gmt_create,omitempty"`
	GmtModified      int64  `json:"gmt_modified,omitempty"`
	Id               int64  `json:"id,omitempty"`
	InheritConfig    string `json:"inherit_config,omitempty"`
	InstId           int64  `json:"inst_id,omitempty"`
	IsDeleted        int64  `json:"is_deleted,omitempty"`
	MainUrl          string `json:"main_url,omitempty"`
	Md5              string `json:"md5,omitempty"`
	MosecJobId       string `json:"mosec_job_id,omitempty"`
	MosecStatus      int64  `json:"mosec_status,omitempty"`
	PackageUrl       string `json:"package_url,omitempty"`
	PluginRefs       string `json:"plugin_refs,omitempty"`
	PluginSize       int64  `json:"plugin_size,omitempty"`
	PluginUrl        string `json:"plugin_url,omitempty"`
	QcloudJobId      string `json:"qcloud_job_id,omitempty"`
	QcloudStatus     int64  `json:"qcloud_status,omitempty"`
	Size             int64  `json:"size,omitempty"`
	SourceUrl        string `json:"source_url,omitempty"`
	TemplateAppId    string `json:"template_app_id,omitempty"`
	TemplateId       int64  `json:"template_id,omitempty"`
	TemplateVersion  string `json:"template_version,omitempty"`
	TinyCliVersion   string `json:"tiny_cli_version,omitempty"`
	Version          string `json:"version,omitempty"`
	VersionUniqueId  string `json:"version_unique_id,omitempty"`
}
