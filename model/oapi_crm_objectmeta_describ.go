package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCrmObjectmetaDescribeRequest() *OapiCrmObjectmetaDescribeRequest {
	return &OapiCrmObjectmetaDescribeRequest{}
}

type OapiCrmObjectmetaDescribeRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCrmObjectmetaDescribeResponse
	Name            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCrmObjectmetaDescribeRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCrmObjectmetaDescribeRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCrmObjectmetaDescribeRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiCrmObjectmetaDescribeRequest) GetName() string {
	return this.Name
}
func (this *OapiCrmObjectmetaDescribeRequest) GetApiMethodName() string {
	return "dingtalk.oapi.crm.objectmeta.describe"
}
func (this *OapiCrmObjectmetaDescribeRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCrmObjectmetaDescribeRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCrmObjectmetaDescribeRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCrmObjectmetaDescribeRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCrmObjectmetaDescribeRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["name"] = this.Name
	return txtParams
}
func (this *OapiCrmObjectmetaDescribeRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Name, "name"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCrmObjectmetaDescribeRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCrmObjectmetaDescribeRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCrmObjectmetaDescribeResponse struct {
	taobao.TaobaoResponse
	Result DObject `json:"result,omitempty"`
}
