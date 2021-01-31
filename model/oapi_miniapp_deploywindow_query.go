package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMiniappDeploywindowQueryRequest() *OapiMiniappDeploywindowQueryRequest {
	return &OapiMiniappDeploywindowQueryRequest{}
}

type OapiMiniappDeploywindowQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMiniappDeploywindowQueryResponse
	ModelKey        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiMiniappDeploywindowQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMiniappDeploywindowQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMiniappDeploywindowQueryRequest) SetModelKey(modelKey2 string) {
	this.ModelKey = modelKey2
}
func (this *OapiMiniappDeploywindowQueryRequest) GetModelKey() string {
	return this.ModelKey
}
func (this *OapiMiniappDeploywindowQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.miniapp.deploywindow.query"
}
func (this *OapiMiniappDeploywindowQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMiniappDeploywindowQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMiniappDeploywindowQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMiniappDeploywindowQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMiniappDeploywindowQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["model_key"] = this.ModelKey
	return txtParams
}
func (this *OapiMiniappDeploywindowQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ModelKey, "modelKey"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMiniappDeploywindowQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMiniappDeploywindowQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMiniappDeploywindowQueryResponse struct {
	taobao.TaobaoResponse
	Data    DeployWindowDoModel `json:"data,omitempty"`
	Errcode int64               `json:"errcode,omitempty"`
	Errmsg  string              `json:"errmsg,omitempty"`
}
type DeployWindowDoModel struct {
	AndroidClientMax  string `json:"android_client_max,omitempty"`
	AndroidClientMin  string `json:"android_client_min,omitempty"`
	AndroidInstanceId int64  `json:"android_instance_id,omitempty"`
	AppId             string `json:"app_id,omitempty"`
	ClientId          int64  `json:"client_id,omitempty"`
	DeployPackageId   int64  `json:"deploy_package_id,omitempty"`
	GmtCreate         int64  `json:"gmt_create,omitempty"`
	GmtModified       int64  `json:"gmt_modified,omitempty"`
	Id                int64  `json:"id,omitempty"`
	InstId            int64  `json:"inst_id,omitempty"`
	IosClientMax      string `json:"ios_client_max,omitempty"`
	IosClientMin      string `json:"ios_client_min,omitempty"`
	IosInstanceId     int64  `json:"ios_instance_id,omitempty"`
	Name              string `json:"name,omitempty"`
	PlatformAndroid   int64  `json:"platform_android,omitempty"`
	PlatformIos       int64  `json:"platform_ios,omitempty"`
	Version           string `json:"version,omitempty"`
}
