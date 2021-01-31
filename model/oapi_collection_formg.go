package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCollectionFormGetRequest() *OapiCollectionFormGetRequest {
	return &OapiCollectionFormGetRequest{}
}

type OapiCollectionFormGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCollectionFormGetResponse
	ActionDate      string
	FormCode        string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCollectionFormGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCollectionFormGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCollectionFormGetRequest) SetActionDate(actionDate2 string) {
	this.ActionDate = actionDate2
}
func (this *OapiCollectionFormGetRequest) GetActionDate() string {
	return this.ActionDate
}
func (this *OapiCollectionFormGetRequest) SetFormCode(formCode2 string) {
	this.FormCode = formCode2
}
func (this *OapiCollectionFormGetRequest) GetFormCode() string {
	return this.FormCode
}
func (this *OapiCollectionFormGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.collection.form.get"
}
func (this *OapiCollectionFormGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCollectionFormGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCollectionFormGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCollectionFormGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCollectionFormGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["action_date"] = this.ActionDate
	txtParams["form_code"] = this.FormCode
	return txtParams
}
func (this *OapiCollectionFormGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ActionDate, "actionDate"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCollectionFormGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCollectionFormGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCollectionFormGetResponse struct {
	taobao.TaobaoResponse
	Result PageResult `json:"result,omitempty"`
}
