package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceProjectCreateRequest() *OapiWorkspaceProjectCreateRequest {
	return &OapiWorkspaceProjectCreateRequest{}
}

type OapiWorkspaceProjectCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp               OapiWorkspaceProjectCreateResponse
	BelongCorpUserid   string
	CreateGroup        bool
	Desc               string
	Name               string
	OpenConversationId string
	OuterId            string
	TemplateId         int64
	TopHttpMethod      string
	TopResponseType    string
	Type               int64
	Userid             string
}

func (this *OapiWorkspaceProjectCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceProjectCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceProjectCreateRequest) SetBelongCorpUserid(belongCorpUserid2 string) {
	this.BelongCorpUserid = belongCorpUserid2
}
func (this *OapiWorkspaceProjectCreateRequest) GetBelongCorpUserid() string {
	return this.BelongCorpUserid
}
func (this *OapiWorkspaceProjectCreateRequest) SetCreateGroup(createGroup2 bool) {
	this.CreateGroup = createGroup2
}
func (this *OapiWorkspaceProjectCreateRequest) GetCreateGroup() bool {
	return this.CreateGroup
}
func (this *OapiWorkspaceProjectCreateRequest) SetDesc(desc2 string) {
	this.Desc = desc2
}
func (this *OapiWorkspaceProjectCreateRequest) GetDesc() string {
	return this.Desc
}
func (this *OapiWorkspaceProjectCreateRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiWorkspaceProjectCreateRequest) GetName() string {
	return this.Name
}
func (this *OapiWorkspaceProjectCreateRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiWorkspaceProjectCreateRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiWorkspaceProjectCreateRequest) SetOuterId(outerId2 string) {
	this.OuterId = outerId2
}
func (this *OapiWorkspaceProjectCreateRequest) GetOuterId() string {
	return this.OuterId
}
func (this *OapiWorkspaceProjectCreateRequest) SetTemplateId(templateId2 int64) {
	this.TemplateId = templateId2
}
func (this *OapiWorkspaceProjectCreateRequest) GetTemplateId() int64 {
	return this.TemplateId
}
func (this *OapiWorkspaceProjectCreateRequest) SetType(type2 int64) {
	this.Type = type2
}
func (this *OapiWorkspaceProjectCreateRequest) GetType() int64 {
	return this.Type
}
func (this *OapiWorkspaceProjectCreateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiWorkspaceProjectCreateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiWorkspaceProjectCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.project.create"
}
func (this *OapiWorkspaceProjectCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceProjectCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceProjectCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceProjectCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceProjectCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["belong_corp_userid"] = this.BelongCorpUserid
	txtParams["create_group"] = this.CreateGroup
	txtParams["desc"] = this.Desc
	txtParams["name"] = this.Name
	txtParams["open_conversation_id"] = this.OpenConversationId
	txtParams["outer_id"] = this.OuterId
	txtParams["template_id"] = this.TemplateId
	txtParams["type"] = this.Type
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiWorkspaceProjectCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BelongCorpUserid, "belongCorpUserid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceProjectCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceProjectCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceProjectCreateResponse struct {
	taobao.TaobaoResponse
	Result  OpenProjectDto `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type OpenProjectMemberDto struct {
	CorpId string       `json:"corp_id,omitempty"`
	Name   string       `json:"name,omitempty"`
	Tags   []OpenTagDto `json:"tags,omitempty"`
	Userid string       `json:"userid,omitempty"`
}
type OpenProjectDto struct {
	CorpId     string               `json:"corp_id,omitempty"`
	CorpSecret string               `json:"corp_secret,omitempty"`
	CreateTime int64                `json:"create_time,omitempty"`
	Creator    OpenProjectMemberDto `json:"creator,omitempty"`
	Desc       string               `json:"desc,omitempty"`
	Name       string               `json:"name,omitempty"`
	OuterId    string               `json:"outer_id,omitempty"`
	Owner      OpenProjectMemberDto `json:"owner,omitempty"`
	Type       int64                `json:"type,omitempty"`
}
