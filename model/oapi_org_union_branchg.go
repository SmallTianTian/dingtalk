package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiOrgUnionBranchGetRequest() *OapiOrgUnionBranchGetRequest {
	return &OapiOrgUnionBranchGetRequest{}
}

type OapiOrgUnionBranchGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiOrgUnionBranchGetResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiOrgUnionBranchGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiOrgUnionBranchGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiOrgUnionBranchGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.org.union.branch.get"
}
func (this *OapiOrgUnionBranchGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiOrgUnionBranchGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiOrgUnionBranchGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiOrgUnionBranchGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiOrgUnionBranchGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiOrgUnionBranchGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiOrgUnionBranchGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiOrgUnionBranchGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiOrgUnionBranchGetResponse struct {
	taobao.TaobaoResponse
	Result  []OpenOrgUnion `json:"result,omitempty"`
	Success bool           `json:"success,omitempty"`
}
type OpenOrgUnion struct {
	UnionCorpid  string `json:"union_corpid,omitempty"`
	UnionOrgName string `json:"union_org_name,omitempty"`
}
