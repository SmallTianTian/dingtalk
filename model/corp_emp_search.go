package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpEmpSearchRequest() *CorpEmpSearchRequest {
	return &CorpEmpSearchRequest{}
}

type CorpEmpSearchRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            CorpEmpSearchResponse
	Keyword         string
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpEmpSearchRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpEmpSearchRequest) SetKeyword(keyword2 string) {
	this.Keyword = keyword2
}
func (this *CorpEmpSearchRequest) GetKeyword() string {
	return this.Keyword
}
func (this *CorpEmpSearchRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *CorpEmpSearchRequest) GetOffset() int64 {
	return this.Offset
}
func (this *CorpEmpSearchRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *CorpEmpSearchRequest) GetSize() int64 {
	return this.Size
}
func (this *CorpEmpSearchRequest) GetApiMethodName() string {
	return "dingtalk.corp.emp.search"
}
func (this *CorpEmpSearchRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpEmpSearchRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpEmpSearchRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpEmpSearchRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpEmpSearchRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpEmpSearchRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["keyword"] = this.Keyword
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *CorpEmpSearchRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpEmpSearchRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpEmpSearchRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpEmpSearchResponse struct {
	taobao.TaobaoResponse
	Result string `json:"result,omitempty"`
}
