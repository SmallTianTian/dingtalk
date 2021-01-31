package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWikiDocPublicDetailRequest() *OapiWikiDocPublicDetailRequest {
	return &OapiWikiDocPublicDetailRequest{}
}

type OapiWikiDocPublicDetailRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWikiDocPublicDetailResponse
	Agentid         int64
	DocId           string
	GroupId         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWikiDocPublicDetailRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWikiDocPublicDetailRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWikiDocPublicDetailRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiWikiDocPublicDetailRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiWikiDocPublicDetailRequest) SetDocId(docId2 string) {
	this.DocId = docId2
}
func (this *OapiWikiDocPublicDetailRequest) GetDocId() string {
	return this.DocId
}
func (this *OapiWikiDocPublicDetailRequest) SetGroupId(groupId2 string) {
	this.GroupId = groupId2
}
func (this *OapiWikiDocPublicDetailRequest) GetGroupId() string {
	return this.GroupId
}
func (this *OapiWikiDocPublicDetailRequest) GetApiMethodName() string {
	return "dingtalk.oapi.wiki.doc.public.detail"
}
func (this *OapiWikiDocPublicDetailRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWikiDocPublicDetailRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWikiDocPublicDetailRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWikiDocPublicDetailRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWikiDocPublicDetailRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["doc_id"] = this.DocId
	txtParams["group_id"] = this.GroupId
	return txtParams
}
func (this *OapiWikiDocPublicDetailRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWikiDocPublicDetailRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWikiDocPublicDetailRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWikiDocPublicDetailResponse struct {
	taobao.TaobaoResponse
	Result  OpenDocVo `json:"result,omitempty"`
	Success bool      `json:"success,omitempty"`
}
