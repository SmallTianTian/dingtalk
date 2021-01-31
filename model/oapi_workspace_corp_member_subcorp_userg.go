package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceCorpMemberSubcorpUserGetRequest() *OapiWorkspaceCorpMemberSubcorpUserGetRequest {
	return &OapiWorkspaceCorpMemberSubcorpUserGetRequest{}
}

type OapiWorkspaceCorpMemberSubcorpUserGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp             OapiWorkspaceCorpMemberSubcorpUserGetResponse
	BelongOrgUserids string
	TargetCorpId     string
	TopHttpMethod    string
	TopResponseType  string
}

func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) SetBelongOrgUserids(belongOrgUserids2 string) {
	this.BelongOrgUserids = belongOrgUserids2
}
func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) GetBelongOrgUserids() string {
	return this.BelongOrgUserids
}
func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) SetTargetCorpId(targetCorpId2 string) {
	this.TargetCorpId = targetCorpId2
}
func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) GetTargetCorpId() string {
	return this.TargetCorpId
}
func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.corp.member.subcorp.user.get"
}
func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["belong_org_userids"] = this.BelongOrgUserids
	txtParams["target_corp_id"] = this.TargetCorpId
	return txtParams
}
func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BelongOrgUserids, "belongOrgUserids"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceCorpMemberSubcorpUserGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceCorpMemberSubcorpUserGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  string `json:"result,omitempty"`
}
