package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewCorpBlazersGetbinddataRequest() *CorpBlazersGetbinddataRequest {
	return &CorpBlazersGetbinddataRequest{}
}

type CorpBlazersGetbinddataRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            CorpBlazersGetbinddataResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *CorpBlazersGetbinddataRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *CorpBlazersGetbinddataRequest) GetApiMethodName() string {
	return "dingtalk.corp.blazers.getbinddata"
}
func (this *CorpBlazersGetbinddataRequest) GetTopResponseType() string {
	return this.TopResponseType
}
func (this *CorpBlazersGetbinddataRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *CorpBlazersGetbinddataRequest) GetTopApiCallType() string {
	return "top"
}
func (this *CorpBlazersGetbinddataRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *CorpBlazersGetbinddataRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *CorpBlazersGetbinddataRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *CorpBlazersGetbinddataRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *CorpBlazersGetbinddataRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *CorpBlazersGetbinddataRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type CorpBlazersGetbinddataResponse struct {
	taobao.TaobaoResponse
	Result DingOpenResult `json:"result,omitempty"`
}
type DingOpenResult struct {
	DingOpenErrcode int64  `json:"ding_open_errcode,omitempty"`
	ErrorMsg        string `json:"error_msg,omitempty"`
	Result          string `json:"result,omitempty"`
	Success         bool   `json:"success,omitempty"`
}
