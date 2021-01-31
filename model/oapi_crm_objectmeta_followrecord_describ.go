package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCrmObjectmetaFollowrecordDescribeRequest() *OapiCrmObjectmetaFollowrecordDescribeRequest {
	return &OapiCrmObjectmetaFollowrecordDescribeRequest{}
}

type OapiCrmObjectmetaFollowrecordDescribeRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmObjectmetaFollowrecordDescribeResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmObjectmetaFollowrecordDescribeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectmetaFollowrecordDescribeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectmetaFollowrecordDescribeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectmeta.followrecord.describe"
}
func (this *OapiCrmObjectmetaFollowrecordDescribeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectmetaFollowrecordDescribeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectmetaFollowrecordDescribeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectmetaFollowrecordDescribeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectmetaFollowrecordDescribeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiCrmObjectmetaFollowrecordDescribeRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCrmObjectmetaFollowrecordDescribeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectmetaFollowrecordDescribeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectmetaFollowrecordDescribeResponse struct {
	taobao.TaobaoResponse
	Errcode int64   `json:"errcode,omitempty"`
	Errmsg  string  `json:"errmsg,omitempty"`
	Result  DObject `json:"result,omitempty"`
}
