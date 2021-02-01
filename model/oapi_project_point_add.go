package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProjectPointAddRequest() *OapiProjectPointAddRequest {
	return &OapiProjectPointAddRequest{}
}

type OapiProjectPointAddRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProjectPointAddResponse
	ActionTime      int64
	RuleCode        string
	RuleName        string
	Score           int64
	TenantId        int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
	Uuid            string
}

func (this *OapiProjectPointAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProjectPointAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProjectPointAddRequest) SetActionTime(actionTime2 int64) {
	this.ActionTime = actionTime2
}
func (this *OapiProjectPointAddRequest) GetActionTime() int64 {
	return this.ActionTime
}
func (this *OapiProjectPointAddRequest) SetRuleCode(ruleCode2 string) {
	this.RuleCode = ruleCode2
}
func (this *OapiProjectPointAddRequest) GetRuleCode() string {
	return this.RuleCode
}
func (this *OapiProjectPointAddRequest) SetRuleName(ruleName2 string) {
	this.RuleName = ruleName2
}
func (this *OapiProjectPointAddRequest) GetRuleName() string {
	return this.RuleName
}
func (this *OapiProjectPointAddRequest) SetScore(score2 int64) {
	this.Score = score2
}
func (this *OapiProjectPointAddRequest) GetScore() int64 {
	return this.Score
}
func (this *OapiProjectPointAddRequest) SetTenantId(tenantId2 int64) {
	this.TenantId = tenantId2
}
func (this *OapiProjectPointAddRequest) GetTenantId() int64 {
	return this.TenantId
}
func (this *OapiProjectPointAddRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiProjectPointAddRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiProjectPointAddRequest) SetUuid(uuid2 string) {
	this.Uuid = uuid2
}
func (this *OapiProjectPointAddRequest) GetUuid() string {
	return this.Uuid
}
func (this *OapiProjectPointAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.project.point.add"
}
func (this *OapiProjectPointAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProjectPointAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProjectPointAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProjectPointAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProjectPointAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["action_time"] = this.ActionTime
	txtParams["rule_code"] = this.RuleCode
	txtParams["rule_name"] = this.RuleName
	txtParams["score"] = this.Score
	txtParams["tenant_id"] = this.TenantId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	txtParams["uuid"] = this.Uuid
	return txtParams
}
func (this *OapiProjectPointAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ActionTime, "actionTime"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProjectPointAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProjectPointAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProjectPointAddResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
