package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMedalCorpmedalGrantRequest() *OapiMedalCorpmedalGrantRequest {
	return &OapiMedalCorpmedalGrantRequest{}
}

type OapiMedalCorpmedalGrantRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMedalCorpmedalGrantResponse
	TemplateId      int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiMedalCorpmedalGrantRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMedalCorpmedalGrantRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMedalCorpmedalGrantRequest) SetTemplateId(templateId2 int64) {
	this.TemplateId = templateId2
}
func (this *OapiMedalCorpmedalGrantRequest) GetTemplateId() int64 {
	return this.TemplateId
}
func (this *OapiMedalCorpmedalGrantRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiMedalCorpmedalGrantRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiMedalCorpmedalGrantRequest) GetApiMethodName() string {
	return "dingtalk.oapi.medal.corpmedal.grant"
}
func (this *OapiMedalCorpmedalGrantRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMedalCorpmedalGrantRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMedalCorpmedalGrantRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMedalCorpmedalGrantRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMedalCorpmedalGrantRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["template_id"] = this.TemplateId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiMedalCorpmedalGrantRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.TemplateId, "templateId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMedalCorpmedalGrantRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMedalCorpmedalGrantRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMedalCorpmedalGrantResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
