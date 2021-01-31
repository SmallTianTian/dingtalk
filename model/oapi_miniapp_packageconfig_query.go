package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMiniappPackageconfigQueryRequest() *OapiMiniappPackageconfigQueryRequest {
	return &OapiMiniappPackageconfigQueryRequest{}
}

type OapiMiniappPackageconfigQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMiniappPackageconfigQueryResponse
	ModelKey        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMiniappPackageconfigQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMiniappPackageconfigQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMiniappPackageconfigQueryRequest) SetModelKey(modelKey2 string) {
	this.ModelKey = modelKey2
}
func (this *OapiMiniappPackageconfigQueryRequest) GetModelKey() string {
	return this.ModelKey
}
func (this *OapiMiniappPackageconfigQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.miniapp.packageconfig.query"
}
func (this *OapiMiniappPackageconfigQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMiniappPackageconfigQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMiniappPackageconfigQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMiniappPackageconfigQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMiniappPackageconfigQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["model_key"] = this.ModelKey
	return txtParams
}
func (this *OapiMiniappPackageconfigQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ModelKey, "modelKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMiniappPackageconfigQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMiniappPackageconfigQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMiniappPackageconfigQueryResponse struct {
	taobao.TaobaoResponse
	Data    PackageConfigDOModel `json:"data,omitempty"`
	Errcode int64                `json:"errcode,omitempty"`
	Errmsg  string               `json:"errmsg,omitempty"`
}
type PackageConfigDOModel struct {
	AppId           string `json:"app_id,omitempty"`
	BuildResultUrl  string `json:"build_result_url,omitempty"`
	FallbackUrl     string `json:"fallback_url,omitempty"`
	IsDeleted       int64  `json:"is_deleted,omitempty"`
	PackagePath     string `json:"package_path,omitempty"`
	PackageType     string `json:"package_type,omitempty"`
	PackageUrl      string `json:"package_url,omitempty"`
	Size            int64  `json:"size,omitempty"`
	VersionUniqueId string `json:"version_unique_id,omitempty"`
}
