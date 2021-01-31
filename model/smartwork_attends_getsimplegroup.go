package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewSmartworkAttendsGetsimplegroupsRequest() *SmartworkAttendsGetsimplegroupsRequest {
	return &SmartworkAttendsGetsimplegroupsRequest{}
}

type SmartworkAttendsGetsimplegroupsRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            SmartworkAttendsGetsimplegroupsResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *SmartworkAttendsGetsimplegroupsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *SmartworkAttendsGetsimplegroupsRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *SmartworkAttendsGetsimplegroupsRequest) GetOffset() int64 {
	return this.Offset
}
func (this *SmartworkAttendsGetsimplegroupsRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *SmartworkAttendsGetsimplegroupsRequest) GetSize() int64 {
	return this.Size
}
func (this *SmartworkAttendsGetsimplegroupsRequest) GetApiMethodName() string {
	return "dingtalk.smartwork.attends.getsimplegroups"
}
func (this *SmartworkAttendsGetsimplegroupsRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *SmartworkAttendsGetsimplegroupsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *SmartworkAttendsGetsimplegroupsRequest) GetTopApiCallType() string {
	return "top"
}
func (this *SmartworkAttendsGetsimplegroupsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *SmartworkAttendsGetsimplegroupsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *SmartworkAttendsGetsimplegroupsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *SmartworkAttendsGetsimplegroupsRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *SmartworkAttendsGetsimplegroupsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *SmartworkAttendsGetsimplegroupsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SmartworkAttendsGetsimplegroupsResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
