package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWikiDocDetailRequest() *OapiWikiDocDetailRequest {
	return &OapiWikiDocDetailRequest{}
}

type OapiWikiDocDetailRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWikiDocDetailResponse
	Agentid         int64
	DocId           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWikiDocDetailRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWikiDocDetailRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWikiDocDetailRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiWikiDocDetailRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiWikiDocDetailRequest) SetDocId(docId2 string) {
	this.DocId = docId2
}
func (this *OapiWikiDocDetailRequest) GetDocId() string {
	return this.DocId
}
func (this *OapiWikiDocDetailRequest) GetApiMethodName() string {
	return "dingtalk.oapi.wiki.doc.detail"
}
func (this *OapiWikiDocDetailRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWikiDocDetailRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWikiDocDetailRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWikiDocDetailRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWikiDocDetailRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["doc_id"] = this.DocId
	return txtParams
}
func (this *OapiWikiDocDetailRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWikiDocDetailRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWikiDocDetailRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWikiDocDetailResponse struct {
	taobao.TaobaoResponse
	Errcode int64     `json:"errcode,omitempty"`
	Errmsg  string    `json:"errmsg,omitempty"`
	Result  OpenDocVo `json:"result,omitempty"`
	Success bool      `json:"success,omitempty"`
}
type OpenDocVo struct {
	GroupId     string `json:"group_id,omitempty"`
	Id          string `json:"id,omitempty"`
	LatestCpUrl string `json:"latest_cp_url,omitempty"`
	Name        string `json:"name,omitempty"`
	RepoId      string `json:"repo_id,omitempty"`
	ShareUrl    string `json:"share_url,omitempty"`
}
