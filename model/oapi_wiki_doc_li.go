package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWikiDocListRequest() *OapiWikiDocListRequest {
	return &OapiWikiDocListRequest{}
}

type OapiWikiDocListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWikiDocListResponse
	Agentid         int64
	Cursor          int64
	RepoId          string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWikiDocListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWikiDocListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWikiDocListRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiWikiDocListRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiWikiDocListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiWikiDocListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiWikiDocListRequest) SetRepoId(repoId2 string) {
	this.RepoId = repoId2
}
func (this *OapiWikiDocListRequest) GetRepoId() string {
	return this.RepoId
}
func (this *OapiWikiDocListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiWikiDocListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiWikiDocListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.wiki.doc.list"
}
func (this *OapiWikiDocListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWikiDocListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWikiDocListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWikiDocListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWikiDocListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["cursor"] = this.Cursor
	txtParams["repo_id"] = this.RepoId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiWikiDocListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWikiDocListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWikiDocListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWikiDocListResponse struct {
	taobao.TaobaoResponse
	Errcode int64          `json:"errcode,omitempty"`
	Errmsg  string         `json:"errmsg,omitempty"`
	Result  OpenPageResult `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type OpenDocVO struct {
	GroupId     string `json:"group_id,omitempty"`
	Id          string `json:"id,omitempty"`
	LatestCpUrl string `json:"latest_cp_url,omitempty"`
	Name        string `json:"name,omitempty"`
	RepoId      string `json:"repo_id,omitempty"`
	ShareUrl    string `json:"share_url,omitempty"`
}
