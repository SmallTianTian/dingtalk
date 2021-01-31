package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWikiRepoListRequest() *OapiWikiRepoListRequest {
	return &OapiWikiRepoListRequest{}
}

type OapiWikiRepoListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWikiRepoListResponse
	Agentid         int64
	Cursor          int64
	GroupId         string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWikiRepoListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWikiRepoListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWikiRepoListRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiWikiRepoListRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiWikiRepoListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiWikiRepoListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiWikiRepoListRequest) SetGroupId(groupId2 string) {
	this.GroupId = groupId2
}
func (this *OapiWikiRepoListRequest) GetGroupId() string {
	return this.GroupId
}
func (this *OapiWikiRepoListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiWikiRepoListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiWikiRepoListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.wiki.repo.list"
}
func (this *OapiWikiRepoListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWikiRepoListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWikiRepoListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWikiRepoListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWikiRepoListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["cursor"] = this.Cursor
	txtParams["group_id"] = this.GroupId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiWikiRepoListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWikiRepoListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWikiRepoListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWikiRepoListResponse struct {
	taobao.TaobaoResponse
	Result  OpenPageResult `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type OpenRepoVo struct {
	Desc    string `json:"desc,omitempty"`
	GroupId string `json:"group_id,omitempty"`
	Icon    string `json:"icon,omitempty"`
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
}
