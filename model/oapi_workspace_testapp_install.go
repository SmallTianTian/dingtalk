package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiWorkspaceTestappInstallRequest() *OapiWorkspaceTestappInstallRequest {
	return &OapiWorkspaceTestappInstallRequest{}
}

type OapiWorkspaceTestappInstallRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceTestappInstallResponse
	InstallTestapp  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceTestappInstallRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceTestappInstallRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceTestappInstallRequest) SetInstallTestapp(installTestapp2 string) {
	this.InstallTestapp = installTestapp2
}
func (this *OapiWorkspaceTestappInstallRequest) GetInstallTestapp() string {
	return this.InstallTestapp
}
func (this *OapiWorkspaceTestappInstallRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.testapp.install"
}
func (this *OapiWorkspaceTestappInstallRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceTestappInstallRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceTestappInstallRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceTestappInstallRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceTestappInstallRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["install_testapp"] = this.InstallTestapp
	return txtParams
}
func (this *OapiWorkspaceTestappInstallRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiWorkspaceTestappInstallRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceTestappInstallRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type AuthOpenIsvMicroAppDto struct {
	CircleCorpId string `json:"circle_corp_id,omitempty"`
}
type OapiWorkspaceTestappInstallResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
