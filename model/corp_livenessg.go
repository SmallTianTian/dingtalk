package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpLivenessGetRequest() *CorpLivenessGetRequest {
	return &CorpLivenessGetRequest{}
}

type CorpLivenessGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpLivenessGetResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpLivenessGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpLivenessGetRequest) GetApiMethodName() string {
	return "dingtalk.corp.liveness.get"
}
func (this *CorpLivenessGetRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpLivenessGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpLivenessGetRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpLivenessGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpLivenessGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpLivenessGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *CorpLivenessGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpLivenessGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpLivenessGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpLivenessGetResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type OrgLivenessVo struct {
	Liveness string `json:"liveness,omitempty"`
}
