package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiUnionCooperateInfoListRequest() *OapiUnionCooperateInfoListRequest {
	return &OapiUnionCooperateInfoListRequest{}
}

type OapiUnionCooperateInfoListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUnionCooperateInfoListResponse
	Status          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUnionCooperateInfoListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUnionCooperateInfoListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUnionCooperateInfoListRequest) SetStatus(status2 int64) {
	this.Status = status2
}
func (this *OapiUnionCooperateInfoListRequest) GetStatus() int64 {
	return this.Status
}
func (this *OapiUnionCooperateInfoListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.union.cooperate.info.list"
}
func (this *OapiUnionCooperateInfoListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUnionCooperateInfoListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUnionCooperateInfoListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUnionCooperateInfoListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUnionCooperateInfoListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["status"] = this.Status
	return txtParams
}
func (this *OapiUnionCooperateInfoListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Status, "status"); err != nil {
		return err
	}
	return nil
}
func (this *OapiUnionCooperateInfoListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUnionCooperateInfoListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUnionCooperateInfoListResponse struct {
	taobao.TaobaoResponse
	Result  []OpenCooperateUnionVo `json:"result,omitempty"`
	Success bool                   `json:"success,omitempty"`
}
type OpenCooperateUnionVo struct {
	AuthLevel    int64    `json:"auth_level,omitempty"`
	DeptId       int64    `json:"dept_id,omitempty"`
	DeptIds      []int64  `json:"dept_ids,omitempty"`
	DeptName     string   `json:"dept_name,omitempty"`
	UnionCorpId  string   `json:"union_corp_id,omitempty"`
	UnionOrgName string   `json:"union_org_name,omitempty"`
	UnionType    int64    `json:"union_type,omitempty"`
	Userids      []string `json:"userids,omitempty"`
}
