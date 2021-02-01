package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMedalCorpmedalRemoveRequest() *OapiMedalCorpmedalRemoveRequest {
	return &OapiMedalCorpmedalRemoveRequest{}
}

type OapiMedalCorpmedalRemoveRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMedalCorpmedalRemoveResponse
	TemplateId      int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiMedalCorpmedalRemoveRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMedalCorpmedalRemoveRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMedalCorpmedalRemoveRequest) SetTemplateId(templateId2 int64) {
	this.TemplateId = templateId2
}
func (this *OapiMedalCorpmedalRemoveRequest) GetTemplateId() int64 {
	return this.TemplateId
}
func (this *OapiMedalCorpmedalRemoveRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiMedalCorpmedalRemoveRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiMedalCorpmedalRemoveRequest) GetApiMethodName() string {
	return "dingtalk.oapi.medal.corpmedal.remove"
}
func (this *OapiMedalCorpmedalRemoveRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMedalCorpmedalRemoveRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMedalCorpmedalRemoveRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMedalCorpmedalRemoveRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMedalCorpmedalRemoveRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["template_id"] = this.TemplateId
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiMedalCorpmedalRemoveRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.TemplateId, "templateId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMedalCorpmedalRemoveRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMedalCorpmedalRemoveRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiMedalCorpmedalRemoveResponse struct {
	taobao.TaobaoResponse
	Result  bool `json:"result,omitempty"`
	Success bool `json:"success,omitempty"`
}
