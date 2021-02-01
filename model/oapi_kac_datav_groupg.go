package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiKacDatavGroupGetRequest() *OapiKacDatavGroupGetRequest {
	return &OapiKacDatavGroupGetRequest{}
}

type OapiKacDatavGroupGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiKacDatavGroupGetResponse
	Request         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiKacDatavGroupGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiKacDatavGroupGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiKacDatavGroupGetRequest) SetRequest(request2 string) {
	this.Request = request2
}
func (this *OapiKacDatavGroupGetRequest) GetRequest() string {
	return this.Request
}
func (this *OapiKacDatavGroupGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.kac.datav.group.get"
}
func (this *OapiKacDatavGroupGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiKacDatavGroupGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiKacDatavGroupGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiKacDatavGroupGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiKacDatavGroupGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.QM_ROOT_TAG] = this.Request
	return txtParams
}
func (this *OapiKacDatavGroupGetRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiKacDatavGroupGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiKacDatavGroupGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type GroupStatisticalSummaryRequest struct {
	DataId string `json:"data_id,omitempty"`
}
type OapiKacDatavGroupGetResponse struct {
	taobao.TaobaoResponse
	Result GroupStatisticalSummaryResponse `json:"result,omitempty"`
}
type GroupStatisticalSummaryResponse struct {
	DeptGroupCnt  int64 `json:"dept_group_cnt,omitempty"`
	InnerGroupCnt int64 `json:"inner_group_cnt,omitempty"`
}
