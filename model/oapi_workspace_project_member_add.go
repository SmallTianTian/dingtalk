package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceProjectMemberAddRequest() *OapiWorkspaceProjectMemberAddRequest {
	return &OapiWorkspaceProjectMemberAddRequest{}
}

type OapiWorkspaceProjectMemberAddRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiWorkspaceProjectMemberAddResponse
	Members         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiWorkspaceProjectMemberAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceProjectMemberAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceProjectMemberAddRequest) SetMembers(members2 string) {
	this.Members = members2
}
func (this *OapiWorkspaceProjectMemberAddRequest) GetMembers() string {
	return this.Members
}
func (this *OapiWorkspaceProjectMemberAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.project.member.add"
}
func (this *OapiWorkspaceProjectMemberAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceProjectMemberAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceProjectMemberAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceProjectMemberAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceProjectMemberAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["members"] = this.Members
	return txtParams
}
func (this *OapiWorkspaceProjectMemberAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.Members, 20, "members"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceProjectMemberAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceProjectMemberAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiWorkspaceProjectMemberAddResponse struct {
	taobao.TaobaoResponse
	Result  []OpenProjectMemberDto `json:"result,omitempty"`
	Success bool                   `json:"success,omitempty"`
}
