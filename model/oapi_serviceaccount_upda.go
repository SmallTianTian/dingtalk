package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiServiceaccountUpdateRequest() *OapiServiceaccountUpdateRequest {
	return &OapiServiceaccountUpdateRequest{}
}

type OapiServiceaccountUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiServiceaccountUpdateResponse
	AvatarMediaId   string
	Brief           string
	Desc            string
	Name            string
	PreviewMediaId  string
	Status          string
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
}

func (this *OapiServiceaccountUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiServiceaccountUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiServiceaccountUpdateRequest) SetAvatarMediaId(avatarMediaId2 string) {
	this.AvatarMediaId = avatarMediaId2
}
func (this *OapiServiceaccountUpdateRequest) GetAvatarMediaId() string {
	return this.AvatarMediaId
}
func (this *OapiServiceaccountUpdateRequest) SetBrief(brief2 string) {
	this.Brief = brief2
}
func (this *OapiServiceaccountUpdateRequest) GetBrief() string {
	return this.Brief
}
func (this *OapiServiceaccountUpdateRequest) SetDesc(desc2 string) {
	this.Desc = desc2
}
func (this *OapiServiceaccountUpdateRequest) GetDesc() string {
	return this.Desc
}
func (this *OapiServiceaccountUpdateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiServiceaccountUpdateRequest) GetName() string {
	return this.Name
}
func (this *OapiServiceaccountUpdateRequest) SetPreviewMediaId(previewMediaId2 string) {
	this.PreviewMediaId = previewMediaId2
}
func (this *OapiServiceaccountUpdateRequest) GetPreviewMediaId() string {
	return this.PreviewMediaId
}
func (this *OapiServiceaccountUpdateRequest) SetStatus(status2 string) {
	this.Status = status2
}
func (this *OapiServiceaccountUpdateRequest) GetStatus() string {
	return this.Status
}
func (this *OapiServiceaccountUpdateRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiServiceaccountUpdateRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiServiceaccountUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.serviceaccount.update"
}
func (this *OapiServiceaccountUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiServiceaccountUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiServiceaccountUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiServiceaccountUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiServiceaccountUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["avatar_media_id"] = this.AvatarMediaId
	txtParams["brief"] = this.Brief
	txtParams["desc"] = this.Desc
	txtParams["name"] = this.Name
	txtParams["preview_media_id"] = this.PreviewMediaId
	txtParams["status"] = this.Status
	txtParams["unionid"] = this.Unionid
	return txtParams
}
func (this *OapiServiceaccountUpdateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckMaxLength(this.Brief, 60, "brief"); err != nil {
		return err
	}
	return nil
}
func (this *OapiServiceaccountUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiServiceaccountUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiServiceaccountUpdateResponse struct {
	taobao.TaobaoResponse
}
