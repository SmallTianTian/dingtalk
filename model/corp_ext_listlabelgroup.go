package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpExtListlabelgroupsRequest() *CorpExtListlabelgroupsRequest {
	return &CorpExtListlabelgroupsRequest{}
}

type CorpExtListlabelgroupsRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpExtListlabelgroupsResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpExtListlabelgroupsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpExtListlabelgroupsRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *CorpExtListlabelgroupsRequest) GetOffset() int64 {
	return this.Offset
}
func (this *CorpExtListlabelgroupsRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *CorpExtListlabelgroupsRequest) GetSize() int64 {
	return this.Size
}
func (this *CorpExtListlabelgroupsRequest) GetApiMethodName() string {
	return "dingtalk.corp.ext.listlabelgroups"
}
func (this *CorpExtListlabelgroupsRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpExtListlabelgroupsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpExtListlabelgroupsRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpExtListlabelgroupsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpExtListlabelgroupsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpExtListlabelgroupsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *CorpExtListlabelgroupsRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpExtListlabelgroupsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpExtListlabelgroupsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpExtListlabelgroupsResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
