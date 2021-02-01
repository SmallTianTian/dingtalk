package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessDentryAuthRequest() *OapiProcessDentryAuthRequest {
	return &OapiProcessDentryAuthRequest{}
}

type OapiProcessDentryAuthRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessDentryAuthResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessDentryAuthRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessDentryAuthRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessDentryAuthRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessDentryAuthRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessDentryAuthRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.dentry.auth"
}
func (this *OapiProcessDentryAuthRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessDentryAuthRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessDentryAuthRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessDentryAuthRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessDentryAuthRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessDentryAuthRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessDentryAuthRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessDentryAuthRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type FileInfo struct {
	FileId  string `json:"file_id,omitempty"`
	SpaceId int64  `json:"space_id,omitempty"`
}
type GrantCspaceRequestV2 struct {
	FileInfos []FileInfo `json:"file_infos,omitempty"`
	Userid    string     `json:"userid,omitempty"`
}
type OapiProcessDentryAuthResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
