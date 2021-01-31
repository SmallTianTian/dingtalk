package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsEvaluateJobmatchStartRequest() *OapiAtsEvaluateJobmatchStartRequest {
	return &OapiAtsEvaluateJobmatchStartRequest{}
}

type OapiAtsEvaluateJobmatchStartRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsEvaluateJobmatchStartResponse
	BizCode         string
	CandidateId     string
	Category        string
	ExtData         string
	InviteUrl       string
	JobId           string
	OuterEvaluateId string
	ResultUrl       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsEvaluateJobmatchStartRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsEvaluateJobmatchStartRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsEvaluateJobmatchStartRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsEvaluateJobmatchStartRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsEvaluateJobmatchStartRequest) SetCandidateId(candidateId2 string) {
	this.CandidateId = candidateId2
}
func (this *OapiAtsEvaluateJobmatchStartRequest) GetCandidateId() string {
	return this.CandidateId
}
func (this *OapiAtsEvaluateJobmatchStartRequest) SetCategory(category2 string) {
	this.Category = category2
}
func (this *OapiAtsEvaluateJobmatchStartRequest) GetCategory() string {
	return this.Category
}
func (this *OapiAtsEvaluateJobmatchStartRequest) SetExtData(extData2 string) {
	this.ExtData = extData2
}
func (this *OapiAtsEvaluateJobmatchStartRequest) GetExtData() string {
	return this.ExtData
}
func (this *OapiAtsEvaluateJobmatchStartRequest) SetInviteUrl(inviteUrl2 string) {
	this.InviteUrl = inviteUrl2
}
func (this *OapiAtsEvaluateJobmatchStartRequest) GetInviteUrl() string {
	return this.InviteUrl
}
func (this *OapiAtsEvaluateJobmatchStartRequest) SetJobId(jobId2 string) {
	this.JobId = jobId2
}
func (this *OapiAtsEvaluateJobmatchStartRequest) GetJobId() string {
	return this.JobId
}
func (this *OapiAtsEvaluateJobmatchStartRequest) SetOuterEvaluateId(outerEvaluateId2 string) {
	this.OuterEvaluateId = outerEvaluateId2
}
func (this *OapiAtsEvaluateJobmatchStartRequest) GetOuterEvaluateId() string {
	return this.OuterEvaluateId
}
func (this *OapiAtsEvaluateJobmatchStartRequest) SetResultUrl(resultUrl2 string) {
	this.ResultUrl = resultUrl2
}
func (this *OapiAtsEvaluateJobmatchStartRequest) GetResultUrl() string {
	return this.ResultUrl
}
func (this *OapiAtsEvaluateJobmatchStartRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.evaluate.jobmatch.start"
}
func (this *OapiAtsEvaluateJobmatchStartRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsEvaluateJobmatchStartRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsEvaluateJobmatchStartRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsEvaluateJobmatchStartRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsEvaluateJobmatchStartRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["candidate_id"] = this.CandidateId
	txtParams["category"] = this.Category
	txtParams["ext_data"] = this.ExtData
	txtParams["invite_url"] = this.InviteUrl
	txtParams["job_id"] = this.JobId
	txtParams["outer_evaluate_id"] = this.OuterEvaluateId
	txtParams["result_url"] = this.ResultUrl
	return txtParams
}
func (this *OapiAtsEvaluateJobmatchStartRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsEvaluateJobmatchStartRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsEvaluateJobmatchStartRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsEvaluateJobmatchStartResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
