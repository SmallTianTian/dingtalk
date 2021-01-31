package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsJobListbyownerRequest() *OapiAtsJobListbyownerRequest {
	return &OapiAtsJobListbyownerRequest{}
}

type OapiAtsJobListbyownerRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsJobListbyownerResponse
	BizCode         string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiAtsJobListbyownerRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsJobListbyownerRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsJobListbyownerRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsJobListbyownerRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsJobListbyownerRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiAtsJobListbyownerRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiAtsJobListbyownerRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.job.listbyowner"
}
func (this *OapiAtsJobListbyownerRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsJobListbyownerRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsJobListbyownerRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsJobListbyownerRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsJobListbyownerRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiAtsJobListbyownerRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsJobListbyownerRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsJobListbyownerRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsJobListbyownerResponse struct {
	taobao.TaobaoResponse
	Result []JobSimpleVo `json:"result,omitempty"`
}
