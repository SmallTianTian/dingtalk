package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiFugongProcessCodeGetRequest() *OapiFugongProcessCodeGetRequest {
	return &OapiFugongProcessCodeGetRequest{}
}

type OapiFugongProcessCodeGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiFugongProcessCodeGetResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiFugongProcessCodeGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiFugongProcessCodeGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiFugongProcessCodeGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.fugong.process_code.get"
}
func (this *OapiFugongProcessCodeGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiFugongProcessCodeGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiFugongProcessCodeGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiFugongProcessCodeGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiFugongProcessCodeGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiFugongProcessCodeGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiFugongProcessCodeGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiFugongProcessCodeGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiFugongProcessCodeGetResponse struct {
	taobao.TaobaoResponse
	Result []string `json:"result,omitempty"`
}
