package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceCorpMemberGetuseridsRequest() *OapiWorkspaceCorpMemberGetuseridsRequest {
	return &OapiWorkspaceCorpMemberGetuseridsRequest{}
}

type OapiWorkspaceCorpMemberGetuseridsRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceCorpMemberGetuseridsResponse
	SourceCorpId    string
	TopHttpMethod   string
	TopResponseType string
	Userids         string
}

func (this *OapiWorkspaceCorpMemberGetuseridsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceCorpMemberGetuseridsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceCorpMemberGetuseridsRequest) SetSourceCorpId(sourceCorpId2 string) {
	this.SourceCorpId = sourceCorpId2
}
func (this *OapiWorkspaceCorpMemberGetuseridsRequest) GetSourceCorpId() string {
	return this.SourceCorpId
}
func (this *OapiWorkspaceCorpMemberGetuseridsRequest) SetUserids(userids2 string) {
	this.Userids = userids2
}
func (this *OapiWorkspaceCorpMemberGetuseridsRequest) GetUserids() string {
	return this.Userids
}
func (this *OapiWorkspaceCorpMemberGetuseridsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.corp.member.getuserids"
}
func (this *OapiWorkspaceCorpMemberGetuseridsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceCorpMemberGetuseridsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceCorpMemberGetuseridsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceCorpMemberGetuseridsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceCorpMemberGetuseridsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["source_corp_id"] = this.SourceCorpId
	txtParams["userids"] = this.Userids
	return txtParams
}
func (this *OapiWorkspaceCorpMemberGetuseridsRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.SourceCorpId, "sourceCorpId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceCorpMemberGetuseridsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceCorpMemberGetuseridsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceCorpMemberGetuseridsResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
}
