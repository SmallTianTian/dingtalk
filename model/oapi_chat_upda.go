package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiChatUpdateRequest() *OapiChatUpdateRequest {
	return &OapiChatUpdateRequest{}
}

type OapiChatUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp                OapiChatUpdateResponse
	AddExtidlist        []string
	AddUseridlist       []string
	ChatBannedType      int64
	Chatid              string
	DelExtidlist        []string
	DelUseridlist       []string
	Icon                string
	IsBan               bool
	ManagementType      int64
	MentionAllAuthority int64
	Name                string
	Owner               string
	OwnerType           string
	Searchable          int64
	ShowHistoryType     int64
	TopHttpMethod       string
	TopResponseType     string
	ValidationType      int64
}

func (this *OapiChatUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiChatUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiChatUpdateRequest) SetAddExtidlist(addExtidlist2 []string) {
	this.AddExtidlist = addExtidlist2
}
func (this *OapiChatUpdateRequest) GetAddExtidlist() []string {
	return this.AddExtidlist
}
func (this *OapiChatUpdateRequest) SetAddUseridlist(addUseridlist2 []string) {
	this.AddUseridlist = addUseridlist2
}
func (this *OapiChatUpdateRequest) GetAddUseridlist() []string {
	return this.AddUseridlist
}
func (this *OapiChatUpdateRequest) SetChatBannedType(chatBannedType2 int64) {
	this.ChatBannedType = chatBannedType2
}
func (this *OapiChatUpdateRequest) GetChatBannedType() int64 {
	return this.ChatBannedType
}
func (this *OapiChatUpdateRequest) SetChatid(chatid2 string) {
	this.Chatid = chatid2
}
func (this *OapiChatUpdateRequest) GetChatid() string {
	return this.Chatid
}
func (this *OapiChatUpdateRequest) SetDelExtidlist(delExtidlist2 []string) {
	this.DelExtidlist = delExtidlist2
}
func (this *OapiChatUpdateRequest) GetDelExtidlist() []string {
	return this.DelExtidlist
}
func (this *OapiChatUpdateRequest) SetDelUseridlist(delUseridlist2 []string) {
	this.DelUseridlist = delUseridlist2
}
func (this *OapiChatUpdateRequest) GetDelUseridlist() []string {
	return this.DelUseridlist
}
func (this *OapiChatUpdateRequest) SetIcon(icon2 string) {
	this.Icon = icon2
}
func (this *OapiChatUpdateRequest) GetIcon() string {
	return this.Icon
}
func (this *OapiChatUpdateRequest) SetIsBan(isBan2 bool) {
	this.IsBan = isBan2
}
func (this *OapiChatUpdateRequest) GetIsBan() bool {
	return this.IsBan
}
func (this *OapiChatUpdateRequest) SetManagementType(managementType2 int64) {
	this.ManagementType = managementType2
}
func (this *OapiChatUpdateRequest) GetManagementType() int64 {
	return this.ManagementType
}
func (this *OapiChatUpdateRequest) SetMentionAllAuthority(mentionAllAuthority2 int64) {
	this.MentionAllAuthority = mentionAllAuthority2
}
func (this *OapiChatUpdateRequest) GetMentionAllAuthority() int64 {
	return this.MentionAllAuthority
}
func (this *OapiChatUpdateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiChatUpdateRequest) GetName() string {
	return this.Name
}
func (this *OapiChatUpdateRequest) SetOwner(owner2 string) {
	this.Owner = owner2
}
func (this *OapiChatUpdateRequest) GetOwner() string {
	return this.Owner
}
func (this *OapiChatUpdateRequest) SetOwnerType(ownerType2 string) {
	this.OwnerType = ownerType2
}
func (this *OapiChatUpdateRequest) GetOwnerType() string {
	return this.OwnerType
}
func (this *OapiChatUpdateRequest) SetSearchable(searchable2 int64) {
	this.Searchable = searchable2
}
func (this *OapiChatUpdateRequest) GetSearchable() int64 {
	return this.Searchable
}
func (this *OapiChatUpdateRequest) SetShowHistoryType(showHistoryType2 int64) {
	this.ShowHistoryType = showHistoryType2
}
func (this *OapiChatUpdateRequest) GetShowHistoryType() int64 {
	return this.ShowHistoryType
}
func (this *OapiChatUpdateRequest) SetValidationType(validationType2 int64) {
	this.ValidationType = validationType2
}
func (this *OapiChatUpdateRequest) GetValidationType() int64 {
	return this.ValidationType
}
func (this *OapiChatUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.chat.update"
}
func (this *OapiChatUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiChatUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiChatUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiChatUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiChatUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["add_extidlist"] = this.AddExtidlist
	txtParams["add_useridlist"] = this.AddUseridlist
	txtParams["chatBannedType"] = this.ChatBannedType
	txtParams["chatid"] = this.Chatid
	txtParams["del_extidlist"] = this.DelExtidlist
	txtParams["del_useridlist"] = this.DelUseridlist
	txtParams["icon"] = this.Icon
	txtParams["isBan"] = this.IsBan
	txtParams["managementType"] = this.ManagementType
	txtParams["mentionAllAuthority"] = this.MentionAllAuthority
	txtParams["name"] = this.Name
	txtParams["owner"] = this.Owner
	txtParams["ownerType"] = this.OwnerType
	txtParams["searchable"] = this.Searchable
	txtParams["showHistoryType"] = this.ShowHistoryType
	txtParams["validationType"] = this.ValidationType
	return txtParams
}
func (this *OapiChatUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiChatUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiChatUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiChatUpdateResponse struct {
	taobao.TaobaoResponse
}
