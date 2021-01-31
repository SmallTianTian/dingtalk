package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiProcessTemplateUpgradeinfoQueryRequest() *OapiProcessTemplateUpgradeinfoQueryRequest {
	return &OapiProcessTemplateUpgradeinfoQueryRequest{}
}

type OapiProcessTemplateUpgradeinfoQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessTemplateUpgradeinfoQueryResponse
	ProcessCodes    string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessTemplateUpgradeinfoQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessTemplateUpgradeinfoQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessTemplateUpgradeinfoQueryRequest) SetProcessCodes(processCodes2 string) {
	this.ProcessCodes = processCodes2
}
func (this *OapiProcessTemplateUpgradeinfoQueryRequest) GetProcessCodes() string {
	return this.ProcessCodes
}
func (this *OapiProcessTemplateUpgradeinfoQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.template.upgradeinfo.query"
}
func (this *OapiProcessTemplateUpgradeinfoQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessTemplateUpgradeinfoQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessTemplateUpgradeinfoQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessTemplateUpgradeinfoQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessTemplateUpgradeinfoQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["process_codes"] = this.ProcessCodes
	return txtParams
}
func (this *OapiProcessTemplateUpgradeinfoQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ProcessCodes, "processCodes"); err != nil {
		return err
	}
	return nil
}
func (this *OapiProcessTemplateUpgradeinfoQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessTemplateUpgradeinfoQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiProcessTemplateUpgradeinfoQueryResponse struct {
	taobao.TaobaoResponse
	Result []Result `json:"result,omitempty"`
}
