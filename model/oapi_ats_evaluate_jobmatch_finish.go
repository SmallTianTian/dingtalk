package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsEvaluateJobmatchFinishRequest() *OapiAtsEvaluateJobmatchFinishRequest {
	return &OapiAtsEvaluateJobmatchFinishRequest{}
}

type OapiAtsEvaluateJobmatchFinishRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp              OapiAtsEvaluateJobmatchFinishResponse
	BizCode           string
	Conclusion        string
	ExtData           string
	OuterEvaluateId   string
	ReportDownloadUrl string
	Result            string
	Score             string
	TopHttpMethod     string
	TopResponseType   string
}

func (this *OapiAtsEvaluateJobmatchFinishRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) SetConclusion(conclusion2 string) {
	this.Conclusion = conclusion2
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) GetConclusion() string {
	return this.Conclusion
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) SetExtData(extData2 string) {
	this.ExtData = extData2
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) GetExtData() string {
	return this.ExtData
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) SetOuterEvaluateId(outerEvaluateId2 string) {
	this.OuterEvaluateId = outerEvaluateId2
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) GetOuterEvaluateId() string {
	return this.OuterEvaluateId
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) SetReportDownloadUrl(reportDownloadUrl2 string) {
	this.ReportDownloadUrl = reportDownloadUrl2
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) GetReportDownloadUrl() string {
	return this.ReportDownloadUrl
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) SetResult(result2 string) {
	this.Result = result2
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) GetResult() string {
	return this.Result
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) SetScore(score2 string) {
	this.Score = score2
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) GetScore() string {
	return this.Score
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.evaluate.jobmatch.finish"
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["conclusion"] = this.Conclusion
	txtParams["ext_data"] = this.ExtData
	txtParams["outer_evaluate_id"] = this.OuterEvaluateId
	txtParams["report_download_url"] = this.ReportDownloadUrl
	txtParams["result"] = this.Result
	txtParams["score"] = this.Score
	return txtParams
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsEvaluateJobmatchFinishRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsEvaluateJobmatchFinishResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
