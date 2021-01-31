package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiHireOnboardingGetRequest() *OapiHireOnboardingGetRequest {
	return &OapiHireOnboardingGetRequest{}
}

type OapiHireOnboardingGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHireOnboardingGetResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiHireOnboardingGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHireOnboardingGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHireOnboardingGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hire.onboarding.get"
}
func (this *OapiHireOnboardingGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHireOnboardingGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHireOnboardingGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHireOnboardingGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHireOnboardingGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiHireOnboardingGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiHireOnboardingGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHireOnboardingGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHireOnboardingGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64        `json:"errcode,omitempty"`
	Errmsg  string       `json:"errmsg,omitempty"`
	Result  OnboardingVo `json:"result,omitempty"`
}
type FeatureGuideVo struct {
	GuideStatus string `json:"guide_status,omitempty"`
	GuideTime   int64  `json:"guide_time,omitempty"`
}
type OnboardingVo struct {
	CandidateRecommendGuide FeatureGuideVo `json:"candidate_recommend_guide,omitempty"`
	CreateTime              int64          `json:"create_time,omitempty"`
	JobGuide                FeatureGuideVo `json:"job_guide,omitempty"`
	ResumeCollectMailGuide  FeatureGuideVo `json:"resume_collect_mail_guide,omitempty"`
}
