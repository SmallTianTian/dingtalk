package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHireOnboardingGuideCompleteRequest() *OapiHireOnboardingGuideCompleteRequest {
	return &OapiHireOnboardingGuideCompleteRequest{}
}

type OapiHireOnboardingGuideCompleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHireOnboardingGuideCompleteResponse
	GuideName       string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiHireOnboardingGuideCompleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHireOnboardingGuideCompleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHireOnboardingGuideCompleteRequest) SetGuideName(guideName2 string) {
	this.GuideName = guideName2
}
func (this *OapiHireOnboardingGuideCompleteRequest) GetGuideName() string {
	return this.GuideName
}
func (this *OapiHireOnboardingGuideCompleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hire.onboarding.guide.complete"
}
func (this *OapiHireOnboardingGuideCompleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHireOnboardingGuideCompleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHireOnboardingGuideCompleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHireOnboardingGuideCompleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHireOnboardingGuideCompleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["guide_name"] = this.GuideName
	return txtParams
}
func (this *OapiHireOnboardingGuideCompleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.GuideName, "guideName"); err != nil {
		return err
	}
	return nil
}
func (this *OapiHireOnboardingGuideCompleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHireOnboardingGuideCompleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHireOnboardingGuideCompleteResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
