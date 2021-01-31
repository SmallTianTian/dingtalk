package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpExtListRequest() *CorpExtListRequest {
	return &CorpExtListRequest{}
}

type CorpExtListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpExtListResponse
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpExtListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpExtListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *CorpExtListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *CorpExtListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *CorpExtListRequest) GetSize() int64 {
	return this.Size
}
func (this *CorpExtListRequest) GetApiMethodName() string {
	return "dingtalk.corp.ext.list"
}
func (this *CorpExtListRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpExtListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpExtListRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpExtListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpExtListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpExtListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *CorpExtListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpExtListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpExtListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpExtListResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
