package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpBlazersGetbizidRequest() *CorpBlazersGetbizidRequest {
	return &CorpBlazersGetbizidRequest{}
}

type CorpBlazersGetbizidRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpBlazersGetbizidResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpBlazersGetbizidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpBlazersGetbizidRequest) GetApiMethodName() string {
	return "dingtalk.corp.blazers.getbizid"
}
func (this *CorpBlazersGetbizidRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpBlazersGetbizidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpBlazersGetbizidRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpBlazersGetbizidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpBlazersGetbizidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpBlazersGetbizidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *CorpBlazersGetbizidRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpBlazersGetbizidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpBlazersGetbizidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpBlazersGetbizidResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
