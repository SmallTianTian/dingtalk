package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiServiceaccountAddRequest() *OapiServiceaccountAddRequest {
	return &OapiServiceaccountAddRequest{}
}

type OapiServiceaccountAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceaccountAddResponse
	AvatarMediaId   string
	Brief           string
	Desc            string
	Name            string
	PreviewMediaId  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiServiceaccountAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceaccountAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceaccountAddRequest) SetAvatarMediaId(avatarMediaId2 string) {
	this.AvatarMediaId = avatarMediaId2
}
func (this *OapiServiceaccountAddRequest) GetAvatarMediaId() string {
	return this.AvatarMediaId
}
func (this *OapiServiceaccountAddRequest) SetBrief(brief2 string) {
	this.Brief = brief2
}
func (this *OapiServiceaccountAddRequest) GetBrief() string {
	return this.Brief
}
func (this *OapiServiceaccountAddRequest) SetDesc(desc2 string) {
	this.Desc = desc2
}
func (this *OapiServiceaccountAddRequest) GetDesc() string {
	return this.Desc
}
func (this *OapiServiceaccountAddRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiServiceaccountAddRequest) GetName() string {
	return this.Name
}
func (this *OapiServiceaccountAddRequest) SetPreviewMediaId(previewMediaId2 string) {
	this.PreviewMediaId = previewMediaId2
}
func (this *OapiServiceaccountAddRequest) GetPreviewMediaId() string {
	return this.PreviewMediaId
}
func (this *OapiServiceaccountAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.serviceaccount.add"
}
func (this *OapiServiceaccountAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceaccountAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceaccountAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceaccountAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceaccountAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["avatar_media_id"] = this.AvatarMediaId
	txtParams["brief"] = this.Brief
	txtParams["desc"] = this.Desc
	txtParams["name"] = this.Name
	txtParams["preview_media_id"] = this.PreviewMediaId
	return txtParams
}
func (this *OapiServiceaccountAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AvatarMediaId, "avatarMediaId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiServiceaccountAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceaccountAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceaccountAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Unionid string `json:"unionid,omitempty"`
}
