package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"time"
)

func NewOapiBlackboardListidsRequest() *OapiBlackboardListidsRequest {
	return &OapiBlackboardListidsRequest{}
}

type OapiBlackboardListidsRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiBlackboardListidsResponse
	QueryRequest    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiBlackboardListidsRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiBlackboardListidsRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiBlackboardListidsRequest) SetQueryRequest(queryRequest2 string) {
	this.QueryRequest = queryRequest2
}
func (this *OapiBlackboardListidsRequest) GetQueryRequest() string {
	return this.QueryRequest
}
func (this *OapiBlackboardListidsRequest) GetApiMethodName() string {
	return "dingtalk.oapi.blackboard.listids"
}
func (this *OapiBlackboardListidsRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiBlackboardListidsRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiBlackboardListidsRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiBlackboardListidsRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiBlackboardListidsRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["query_request"] = this.QueryRequest
	return txtParams
}
func (this *OapiBlackboardListidsRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiBlackboardListidsRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiBlackboardListidsRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiBlackboardQueryVo struct {
	CategoryId      string    `json:"category_id,omitempty"`
	EndTime         time.Time `json:"end_time,omitempty"`
	OperationUserid string    `json:"operation_userid,omitempty"`
	Page            int64     `json:"page,omitempty"`
	PageSize        int64     `json:"page_size,omitempty"`
	StartTime       time.Time `json:"start_time,omitempty"`
}
type OapiBlackboardListidsResponse struct {
	taobao.TaobaoResponse
	Errcode int64    `json:"errcode,omitempty"`
	Errmsg  string   `json:"errmsg,omitempty"`
	Result  []string `json:"result,omitempty"`
	Success bool     `json:"success,omitempty"`
}
