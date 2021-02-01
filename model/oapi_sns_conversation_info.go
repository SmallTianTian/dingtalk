package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiSnsConversationInfoRequest() *OapiSnsConversationInfoRequest {
	return &OapiSnsConversationInfoRequest{}
}

type OapiSnsConversationInfoRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp               OapiSnsConversationInfoResponse
	OpenConversationId string
	TopHttpMethod      string
	TopResponseType    string
}

func (this *OapiSnsConversationInfoRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiSnsConversationInfoRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiSnsConversationInfoRequest) SetOpenConversationId(openConversationId2 string) {
	this.OpenConversationId = openConversationId2
}
func (this *OapiSnsConversationInfoRequest) GetOpenConversationId() string {
	return this.OpenConversationId
}
func (this *OapiSnsConversationInfoRequest) GetApiMethodName() string {
	return "dingtalk.oapi.sns.conversation.info"
}
func (this *OapiSnsConversationInfoRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiSnsConversationInfoRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiSnsConversationInfoRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiSnsConversationInfoRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiSnsConversationInfoRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_conversation_id"] = this.OpenConversationId
	return txtParams
}
func (this *OapiSnsConversationInfoRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OpenConversationId, "openConversationId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiSnsConversationInfoRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiSnsConversationInfoRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiSnsConversationInfoResponse struct {
	taobao.TaobaoResponse
	Result  SnsOpenGroupInfoResponse `json:"result,omitempty"`
	Success bool                     `json:"success,omitempty"`
}
type SnsOpenGroupInfoResponse struct {
	Icon               string `json:"icon,omitempty"`
	OpenConversationId string `json:"open_conversation_id,omitempty"`
	OwnerUnionid       string `json:"owner_unionid,omitempty"`
	RobotWebHookUrl    string `json:"robot_web_hook_url,omitempty"`
	TemplateId         string `json:"template_id,omitempty"`
	Title              string `json:"title,omitempty"`
}
