package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWikiGroupPublicListRequest() *OapiWikiGroupPublicListRequest {
	return &OapiWikiGroupPublicListRequest{}
}

type OapiWikiGroupPublicListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWikiGroupPublicListResponse
	AuthCode        string
	Cursor          int64
	RoleType        int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWikiGroupPublicListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWikiGroupPublicListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWikiGroupPublicListRequest) SetAuthCode(authCode2 string) {
	this.AuthCode = authCode2
}
func (this *OapiWikiGroupPublicListRequest) GetAuthCode() string {
	return this.AuthCode
}
func (this *OapiWikiGroupPublicListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiWikiGroupPublicListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiWikiGroupPublicListRequest) SetRoleType(roleType2 int64) {
	this.RoleType = roleType2
}
func (this *OapiWikiGroupPublicListRequest) GetRoleType() int64 {
	return this.RoleType
}
func (this *OapiWikiGroupPublicListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiWikiGroupPublicListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiWikiGroupPublicListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.wiki.group.public.list"
}
func (this *OapiWikiGroupPublicListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWikiGroupPublicListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWikiGroupPublicListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWikiGroupPublicListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWikiGroupPublicListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["auth_code"] = this.AuthCode
	txtParams["cursor"] = this.Cursor
	txtParams["role_type"] = this.RoleType
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiWikiGroupPublicListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AuthCode, "authCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWikiGroupPublicListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWikiGroupPublicListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWikiGroupPublicListResponse struct {
	taobao.TaobaoResponse
	Result  OpenPageResult `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type OpenGroupVO struct {
	Desc string `json:"desc,omitempty"`
	Icon string `json:"icon,omitempty"`
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
