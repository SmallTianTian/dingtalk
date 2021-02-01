package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceCorpMemberRemoveRequest() *OapiWorkspaceCorpMemberRemoveRequest {
	return &OapiWorkspaceCorpMemberRemoveRequest{}
}

type OapiWorkspaceCorpMemberRemoveRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceCorpMemberRemoveResponse
	TargetCorpId    string
	TopHttpMethod   string
	TopResponseType string
	Userids         string
}

func (this *OapiWorkspaceCorpMemberRemoveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceCorpMemberRemoveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceCorpMemberRemoveRequest) SetTargetCorpId(targetCorpId2 string) {
	this.TargetCorpId = targetCorpId2
}
func (this *OapiWorkspaceCorpMemberRemoveRequest) GetTargetCorpId() string {
	return this.TargetCorpId
}
func (this *OapiWorkspaceCorpMemberRemoveRequest) SetUserids(userids2 string) {
	this.Userids = userids2
}
func (this *OapiWorkspaceCorpMemberRemoveRequest) GetUserids() string {
	return this.Userids
}
func (this *OapiWorkspaceCorpMemberRemoveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.corp.member.remove"
}
func (this *OapiWorkspaceCorpMemberRemoveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceCorpMemberRemoveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceCorpMemberRemoveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceCorpMemberRemoveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceCorpMemberRemoveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["target_corp_id"] = this.TargetCorpId
	txtParams["userids"] = this.Userids
	return txtParams
}
func (this *OapiWorkspaceCorpMemberRemoveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.TargetCorpId, "targetCorpId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceCorpMemberRemoveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceCorpMemberRemoveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceCorpMemberRemoveResponse struct {
	taobao.TaobaoResponse
	Result []string `json:"result,omitempty"`
}
