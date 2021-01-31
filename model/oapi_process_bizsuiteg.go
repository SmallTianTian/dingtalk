package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiProcessBizsuiteGetRequest() *OapiProcessBizsuiteGetRequest {
	return &OapiProcessBizsuiteGetRequest{}
}

type OapiProcessBizsuiteGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiProcessBizsuiteGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiProcessBizsuiteGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiProcessBizsuiteGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiProcessBizsuiteGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiProcessBizsuiteGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiProcessBizsuiteGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.process.bizsuite.get"
}
func (this *OapiProcessBizsuiteGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiProcessBizsuiteGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiProcessBizsuiteGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiProcessBizsuiteGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiProcessBizsuiteGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiProcessBizsuiteGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiProcessBizsuiteGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiProcessBizsuiteGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OAPIFormDataReqVO struct {
	ExtendValue string `json:"extend_value,omitempty"`
	Label       string `json:"label,omitempty"`
	Value       string `json:"value,omitempty"`
}
type BaseSuiteRequest struct {
	ActionType   string              `json:"action_type,omitempty"`
	Agentid      int64               `json:"agentid,omitempty"`
	BizType      string              `json:"biz_type,omitempty"`
	FormDataList []OAPIFormDataReqVO `json:"form_data_list,omitempty"`
	ProcessCode  string              `json:"process_code,omitempty"`
	SeqId        int64               `json:"seq_id,omitempty"`
	Userid       string              `json:"userid,omitempty"`
}
type OapiProcessBizsuiteGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64             `json:"errcode,omitempty"`
	Errmsg  string            `json:"errmsg,omitempty"`
	Result  BaseSuiteResponse `json:"result,omitempty"`
	Success bool              `json:"success,omitempty"`
}
type FormDataResponseVO struct {
	BizAlias    string `json:"biz_alias,omitempty"`
	ExtendValue string `json:"extend_value,omitempty"`
	Key         string `json:"key,omitempty"`
	Lable       string `json:"lable,omitempty"`
	Value       string `json:"value,omitempty"`
}
type BaseSuiteResponse struct {
	FormDataList []FormDataResponseVO `json:"form_data_list,omitempty"`
	SeqId        int64                `json:"seq_id,omitempty"`
}
