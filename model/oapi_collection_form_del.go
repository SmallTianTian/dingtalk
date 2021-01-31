package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCollectionFormDeleteRequest() *OapiCollectionFormDeleteRequest {
	return &OapiCollectionFormDeleteRequest{}
}

type OapiCollectionFormDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCollectionFormDeleteResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCollectionFormDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCollectionFormDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCollectionFormDeleteRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiCollectionFormDeleteRequest) GetRequest() string {
	return this.Request
}
func (this *OapiCollectionFormDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.collection.form.delete"
}
func (this *OapiCollectionFormDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCollectionFormDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCollectionFormDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCollectionFormDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCollectionFormDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiCollectionFormDeleteRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCollectionFormDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCollectionFormDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type FormSchemaDeleteRequest struct {
	BizType  int64  `json:"biz_type,omitempty"`
	FormCode string `json:"form_code,omitempty"`
	Userid   string `json:"userid,omitempty"`
}
type OapiCollectionFormDeleteResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
