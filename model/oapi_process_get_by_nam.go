package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessGetByNameRequest() *OapiProcessGetByNameRequest {
	return &OapiProcessGetByNameRequest{}
}

type OapiProcessGetByNameRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessGetByNameResponse
	Name            string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessGetByNameRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessGetByNameRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessGetByNameRequest) SetName(name2 string) {
	this.Name = name2
}
func (this *OapiProcessGetByNameRequest) GetName() string {
	return this.Name
}
func (this *OapiProcessGetByNameRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.get_by_name"
}
func (this *OapiProcessGetByNameRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessGetByNameRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessGetByNameRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessGetByNameRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessGetByNameRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["name"] = this.Name
	return txtParams
}
func (this *OapiProcessGetByNameRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Name, "name"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessGetByNameRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessGetByNameRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessGetByNameResponse struct {
	taobao.TaobaoResponse

	ProcessCode string `json:"process_code,omitempty"`
}
