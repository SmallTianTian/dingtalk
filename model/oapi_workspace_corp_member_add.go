package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiWorkspaceCorpMemberAddRequest() *OapiWorkspaceCorpMemberAddRequest {
	return &OapiWorkspaceCorpMemberAddRequest{}
}

type OapiWorkspaceCorpMemberAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp             OapiWorkspaceCorpMemberAddResponse
	MemberAddDtoList string
	TargetCorpId     string
	TopHttpMethod    string
	TopResponseType  string
}

func (this *OapiWorkspaceCorpMemberAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiWorkspaceCorpMemberAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiWorkspaceCorpMemberAddRequest) SetMemberAddDtoList(memberAddDtoList2 string) {
	this.MemberAddDtoList = memberAddDtoList2
}
func (this *OapiWorkspaceCorpMemberAddRequest) GetMemberAddDtoList() string {
	return this.MemberAddDtoList
}
func (this *OapiWorkspaceCorpMemberAddRequest) SetTargetCorpId(targetCorpId2 string) {
	this.TargetCorpId = targetCorpId2
}
func (this *OapiWorkspaceCorpMemberAddRequest) GetTargetCorpId() string {
	return this.TargetCorpId
}
func (this *OapiWorkspaceCorpMemberAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.workspace.corp.member.add"
}
func (this *OapiWorkspaceCorpMemberAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiWorkspaceCorpMemberAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiWorkspaceCorpMemberAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiWorkspaceCorpMemberAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiWorkspaceCorpMemberAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["member_add_dto_list"] = this.MemberAddDtoList
	txtParams["target_corp_id"] = this.TargetCorpId
	return txtParams
}
func (this *OapiWorkspaceCorpMemberAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.MemberAddDtoList, 20, "memberAddDtoList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiWorkspaceCorpMemberAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiWorkspaceCorpMemberAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenMemberAddDto struct {
	FromCorpId string `json:"from_corp_id,omitempty"`
	FromUserid string `json:"from_userid,omitempty"`
	Userid     string `json:"userid,omitempty"`
}
type OapiWorkspaceCorpMemberAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64           `json:"errcode,omitempty"`
	Errmsg  string          `json:"errmsg,omitempty"`
	Result  []OpenMemberDto `json:"result,omitempty"`
	Success bool            `json:"success,omitempty"`
}
type OpenTagDto struct {
	Code string `json:taobao.ERROR_COD,omitemptyE`
	Name string `json:"name,omitempty"`
}
type OpenMemberDto struct {
	CorpId     string       `json:"corp_id,omitempty"`
	FromCorpId string       `json:"from_corp_id,omitempty"`
	FromUserid string       `json:"from_userid,omitempty"`
	Tags       []OpenTagDto `json:"tags,omitempty"`
	Userid     string       `json:"userid,omitempty"`
}
