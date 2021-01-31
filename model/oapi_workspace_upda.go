package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiWorkspaceUpdateRequest() *OapiWorkspaceUpdateRequest {
	return &OapiWorkspaceUpdateRequest{}
}

type OapiWorkspaceUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceUpdateResponse
	TopHttpMethod   string
	TopResponseType string
	UpdateInfo      string
}

func (this *OapiWorkspaceUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceUpdateRequest) SetUpdateInfo(updateInfo2 string) {
	this.UpdateInfo = updateInfo2
}
func (this *OapiWorkspaceUpdateRequest) GetUpdateInfo() string {
	return this.UpdateInfo
}
func (this *OapiWorkspaceUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.update"
}
func (this *OapiWorkspaceUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["update_info"] = this.UpdateInfo
	return txtParams
}
func (this *OapiWorkspaceUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiWorkspaceUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenUpdateWorkspaceDto struct {
	Desc        string `json:"desc,omitempty"`
	LogoMediaId string `json:"logo_media_id,omitempty"`
	Name        string `json:"name,omitempty"`
	OwnerUserid string `json:"owner_userid,omitempty"`
}
type OapiWorkspaceUpdateResponse struct {
	taobao.TaobaoResponse
	Errcode int64            `json:"errcode,omitempty"`
	Errmsg  string           `json:"errmsg,omitempty"`
	Result  OpenWorkspaceDto `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
type OpenWorkspaceDto struct {
	CreateTime int64         `json:"create_time,omitempty"`
	Creator    OpenMemberDto `json:"creator,omitempty"`
	Desc       string        `json:"desc,omitempty"`
	Name       string        `json:"name,omitempty"`
	OuterId    string        `json:"outer_id,omitempty"`
	Owner      OpenMemberDto `json:"owner,omitempty"`
	Type       int64         `json:"type,omitempty"`
}
