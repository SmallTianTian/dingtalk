package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWikiDocPublicListRequest() *OapiWikiDocPublicListRequest {
	return &OapiWikiDocPublicListRequest{}
}

type OapiWikiDocPublicListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWikiDocPublicListResponse
	Agentid         int64
	Cursor          int64
	GroupId         string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWikiDocPublicListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWikiDocPublicListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWikiDocPublicListRequest) SetAgentid(agentid2 int64) {
	this.Agentid = agentid2
}
func (this *OapiWikiDocPublicListRequest) GetAgentid() int64 {
	return this.Agentid
}
func (this *OapiWikiDocPublicListRequest) SetCursor(cursor2 int64) {
	this.Cursor = cursor2
}
func (this *OapiWikiDocPublicListRequest) GetCursor() int64 {
	return this.Cursor
}
func (this *OapiWikiDocPublicListRequest) SetGroupId(groupId2 string) {
	this.GroupId = groupId2
}
func (this *OapiWikiDocPublicListRequest) GetGroupId() string {
	return this.GroupId
}
func (this *OapiWikiDocPublicListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiWikiDocPublicListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiWikiDocPublicListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.wiki.doc.public.list"
}
func (this *OapiWikiDocPublicListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWikiDocPublicListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWikiDocPublicListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWikiDocPublicListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWikiDocPublicListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["agentid"] = this.Agentid
	txtParams["cursor"] = this.Cursor
	txtParams["group_id"] = this.GroupId
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiWikiDocPublicListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Agentid, "agentid"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWikiDocPublicListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWikiDocPublicListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWikiDocPublicListResponse struct {
	taobao.TaobaoResponse
	Result  OpenPageResult `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
