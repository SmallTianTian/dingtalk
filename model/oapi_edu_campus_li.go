package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduCampusListRequest() *OapiEduCampusListRequest {
	return &OapiEduCampusListRequest{}
}

type OapiEduCampusListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCampusListResponse
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiEduCampusListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCampusListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCampusListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.campus.list"
}
func (this *OapiEduCampusListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCampusListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCampusListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCampusListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCampusListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	return txtParams
}
func (this *OapiEduCampusListRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduCampusListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCampusListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduCampusListResponse struct {
	taobao.TaobaoResponse
	Result  []CampusResponse `json:"result,omitempty"`
	Success bool             `json:"success,omitempty"`
}
