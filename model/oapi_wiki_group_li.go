package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWikiGroupListRequest() *OapiWikiGroupListRequest {
	return &OapiWikiGroupListRequest{}
}

type OapiWikiGroupListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWikiGroupListResponse
	AuthCode        string
	Cursor          int64
	GroupType       int64
	RoleType        int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWikiGroupListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWikiGroupListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWikiGroupListRequest) SetAuthCode(authCode2 string) {
	this.AuthCode = authCode2
}
func (this *OapiWikiGroupListRequest) GetAuthCode() string {
	return this.AuthCode
}
func (this *OapiWikiGroupListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiWikiGroupListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiWikiGroupListRequest) SetGroupType(groupType2 int64) {
	this.GroupType = groupType2
}
func (this *OapiWikiGroupListRequest) GetGroupType() int64 {
	return this.GroupType
}
func (this *OapiWikiGroupListRequest) SetRoleType(roleType2 int64) {
	this.RoleType = roleType2
}
func (this *OapiWikiGroupListRequest) GetRoleType() int64 {
	return this.RoleType
}
func (this *OapiWikiGroupListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiWikiGroupListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiWikiGroupListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.wiki.group.list"
}
func (this *OapiWikiGroupListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWikiGroupListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWikiGroupListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWikiGroupListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWikiGroupListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["auth_code"] = this.AuthCode
	txtParams["cursor"] = this.Cursor
	txtParams["group_type"] = this.GroupType
	txtParams["role_type"] = this.RoleType
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiWikiGroupListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.AuthCode, "authCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWikiGroupListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWikiGroupListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWikiGroupListResponse struct {
	taobao.TaobaoResponse
	Errcode int64          `json:"errcode,omitempty"`
	Errmsg  string         `json:"errmsg,omitempty"`
	Result  OpenPageResult `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type OpenGroupVo struct {
	Desc string `json:"desc,omitempty"`
	Icon string `json:"icon,omitempty"`
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
