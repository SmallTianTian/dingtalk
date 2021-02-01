package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessCleanRequest() *OapiProcessCleanRequest {
	return &OapiProcessCleanRequest{}
}

type OapiProcessCleanRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessCleanResponse
	Corpid          string
	ProcessCode     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessCleanRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessCleanRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessCleanRequest) SetCorpid(corpid2 string) {
	this.Corpid = corpid2
}
func (this *OapiProcessCleanRequest) GetCorpid() string {
	return this.Corpid
}
func (this *OapiProcessCleanRequest) SetProcessCode(processCode2 string) {
	this.ProcessCode = processCode2
}
func (this *OapiProcessCleanRequest) GetProcessCode() string {
	return this.ProcessCode
}
func (this *OapiProcessCleanRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.clean"
}
func (this *OapiProcessCleanRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessCleanRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessCleanRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessCleanRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessCleanRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["corpid"] = this.Corpid
	txtParams["process_code"] = this.ProcessCode
	return txtParams
}
func (this *OapiProcessCleanRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProcessCode, "processCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessCleanRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessCleanRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessCleanResponse struct {
	taobao.TaobaoResponse
}
