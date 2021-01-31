package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpBlazersRemovemappingRequest() *CorpBlazersRemovemappingRequest {
	return &CorpBlazersRemovemappingRequest{}
}

type CorpBlazersRemovemappingRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpBlazersRemovemappingResponse
	BizId           string
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpBlazersRemovemappingRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpBlazersRemovemappingRequest) SetBizId(bizId2 string) {
	this.BizId = bizId2
}
func (this *CorpBlazersRemovemappingRequest) GetBizId() string {
	return this.BizId
}
func (this *CorpBlazersRemovemappingRequest) GetApiMethodName() string {
	return "dingtalk.corp.blazers.removemapping"
}
func (this *CorpBlazersRemovemappingRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpBlazersRemovemappingRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpBlazersRemovemappingRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpBlazersRemovemappingRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpBlazersRemovemappingRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpBlazersRemovemappingRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_id"] = this.BizId
	return txtParams
}
func (this *CorpBlazersRemovemappingRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpBlazersRemovemappingRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpBlazersRemovemappingRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpBlazersRemovemappingResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
