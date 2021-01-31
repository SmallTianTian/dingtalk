package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMedalCorpmedalWearRequest() *OapiMedalCorpmedalWearRequest {
	return &OapiMedalCorpmedalWearRequest{}
}

type OapiMedalCorpmedalWearRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMedalCorpmedalWearResponse
	Operation       int64
	TemplateId      int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiMedalCorpmedalWearRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMedalCorpmedalWearRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMedalCorpmedalWearRequest) SetOperation(operation2 int64) {
	this.Operation = operation2
}
func (this *OapiMedalCorpmedalWearRequest) GetOperation() int64 {
	return this.Operation
}
func (this *OapiMedalCorpmedalWearRequest) SetTemplateId(templateId2 int64) {
	this.TemplateId = templateId2
}
func (this *OapiMedalCorpmedalWearRequest) GetTemplateId() int64 {
	return this.TemplateId
}
func (this *OapiMedalCorpmedalWearRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiMedalCorpmedalWearRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiMedalCorpmedalWearRequest) GetApiMethodName() string {
	return "dingtalk.oapi.medal.corpmedal.wear"
}
func (this *OapiMedalCorpmedalWearRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMedalCorpmedalWearRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMedalCorpmedalWearRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMedalCorpmedalWearRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMedalCorpmedalWearRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["operation"] = this.Operation
	txtParams["template_id"] = this.TemplateId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiMedalCorpmedalWearRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Operation, "operation"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMedalCorpmedalWearRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMedalCorpmedalWearRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMedalCorpmedalWearResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
