package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCspaceFilePresignedurlGetRequest() *OapiCspaceFilePresignedurlGetRequest {
	return &OapiCspaceFilePresignedurlGetRequest{}
}

type OapiCspaceFilePresignedurlGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCspaceFilePresignedurlGetResponse
	Dentryid        string
	ExpireSeconds   int64
	InnerInvoke     bool
	Spaceid         int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCspaceFilePresignedurlGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCspaceFilePresignedurlGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCspaceFilePresignedurlGetRequest) SetDentryid(dentryid2 string) {
	this.Dentryid = dentryid2
}
func (this *OapiCspaceFilePresignedurlGetRequest) GetDentryid() string {
	return this.Dentryid
}
func (this *OapiCspaceFilePresignedurlGetRequest) SetExpireSeconds(expireSeconds2 int64) {
	this.ExpireSeconds = expireSeconds2
}
func (this *OapiCspaceFilePresignedurlGetRequest) GetExpireSeconds() int64 {
	return this.ExpireSeconds
}
func (this *OapiCspaceFilePresignedurlGetRequest) SetInnerInvoke(innerInvoke2 bool) {
	this.InnerInvoke = innerInvoke2
}
func (this *OapiCspaceFilePresignedurlGetRequest) GetInnerInvoke() bool {
	return this.InnerInvoke
}
func (this *OapiCspaceFilePresignedurlGetRequest) SetSpaceid(spaceid2 int64) {
	this.Spaceid = spaceid2
}
func (this *OapiCspaceFilePresignedurlGetRequest) GetSpaceid() int64 {
	return this.Spaceid
}
func (this *OapiCspaceFilePresignedurlGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.cspace.file.presignedurl.get"
}
func (this *OapiCspaceFilePresignedurlGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCspaceFilePresignedurlGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCspaceFilePresignedurlGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCspaceFilePresignedurlGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCspaceFilePresignedurlGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dentryid"] = this.Dentryid
	txtParams["expire_seconds"] = this.ExpireSeconds
	txtParams["inner_invoke"] = this.InnerInvoke
	txtParams["spaceid"] = this.Spaceid
	return txtParams
}
func (this *OapiCspaceFilePresignedurlGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Dentryid, "dentryid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCspaceFilePresignedurlGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCspaceFilePresignedurlGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCspaceFilePresignedurlGetResponse struct {
	taobao.TaobaoResponse
	Errcode            int64             `json:"errcode,omitempty"`
	Errmsg             string            `json:"errmsg,omitempty"`
	PresignedUrlResult GenerateUrlResult `json:"presigned_url_result,omitempty"`
	Success            bool              `json:"success,omitempty"`
}
type GenerateUrlResult struct {
	Code    string `json:"code,omitempty"`
	Mediaid string `json:"mediaid,omitempty"`
	Url     string `json:"url,omitempty"`
}
