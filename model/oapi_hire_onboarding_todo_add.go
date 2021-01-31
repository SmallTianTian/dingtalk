package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHireOnboardingTodoAddRequest() *OapiHireOnboardingTodoAddRequest {
	return &OapiHireOnboardingTodoAddRequest{}
}

type OapiHireOnboardingTodoAddRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHireOnboardingTodoAddResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiHireOnboardingTodoAddRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHireOnboardingTodoAddRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHireOnboardingTodoAddRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiHireOnboardingTodoAddRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiHireOnboardingTodoAddRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hire.onboarding.todo.add"
}
func (this *OapiHireOnboardingTodoAddRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHireOnboardingTodoAddRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHireOnboardingTodoAddRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHireOnboardingTodoAddRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHireOnboardingTodoAddRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiHireOnboardingTodoAddRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiHireOnboardingTodoAddRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHireOnboardingTodoAddRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHireOnboardingTodoAddResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
}
