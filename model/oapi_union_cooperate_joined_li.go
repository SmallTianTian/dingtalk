package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiUnionCooperateJoinedListRequest() *OapiUnionCooperateJoinedListRequest {
	return &OapiUnionCooperateJoinedListRequest{}
}

type OapiUnionCooperateJoinedListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUnionCooperateJoinedListResponse
	Status          int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiUnionCooperateJoinedListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiUnionCooperateJoinedListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUnionCooperateJoinedListRequest) SetStatus(status2 int64) {
	this.Status = status2
}
func (this *OapiUnionCooperateJoinedListRequest) GetStatus() int64 {
	return this.Status
}
func (this *OapiUnionCooperateJoinedListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.union.cooperate.joined.list"
}
func (this *OapiUnionCooperateJoinedListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUnionCooperateJoinedListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUnionCooperateJoinedListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUnionCooperateJoinedListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUnionCooperateJoinedListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["status"] = this.Status
	return txtParams
}
func (this *OapiUnionCooperateJoinedListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Status, "status"); err != nil {
		return err
	}
	return nil
}
func (this *OapiUnionCooperateJoinedListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUnionCooperateJoinedListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUnionCooperateJoinedListResponse struct {
	taobao.TaobaoResponse
	Result  []OpenCooperateOrgVo `json:"result,omitempty"`
	Success bool                 `json:"success,omitempty"`
}
type OpenCooperateOrgVo struct {
	BelongCorpId  string `json:"belong_corp_id,omitempty"`
	BelongOrgName string `json:"belong_org_name,omitempty"`
	CorpId        string `json:"corp_id,omitempty"`
	OrgName       string `json:"org_name,omitempty"`
}
