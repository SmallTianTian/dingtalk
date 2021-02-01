package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCollectionFormStopRequest() *OapiCollectionFormStopRequest {
	return &OapiCollectionFormStopRequest{}
}

type OapiCollectionFormStopRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCollectionFormStopResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCollectionFormStopRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCollectionFormStopRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCollectionFormStopRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiCollectionFormStopRequest) GetRequest() string {
	return this.Request
}
func (this *OapiCollectionFormStopRequest) GetApiMethodName() string {
	return "dingtalk.oapi.collection.form.stop"
}
func (this *OapiCollectionFormStopRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCollectionFormStopRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCollectionFormStopRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCollectionFormStopRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCollectionFormStopRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiCollectionFormStopRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCollectionFormStopRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCollectionFormStopRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type FormSchemaStopRequest struct {
	BizType  int64  `json:"biz_type,omitempty"`
	FormCode string `json:"form_code,omitempty"`
	Userid   string `json:"userid,omitempty"`
}
type OapiCollectionFormStopResponse struct {
	taobao.TaobaoResponse
	Result  string `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
