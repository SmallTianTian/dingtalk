package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiCircleEnworkUpdateRequest() *OapiCircleEnworkUpdateRequest {
	return &OapiCircleEnworkUpdateRequest{}
}

type OapiCircleEnworkUpdateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCircleEnworkUpdateResponse
	OpenUpdateDto   string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCircleEnworkUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCircleEnworkUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCircleEnworkUpdateRequest) SetOpenUpdateDto(openUpdateDto2 string) {
	this.OpenUpdateDto = openUpdateDto2
}
func (this *OapiCircleEnworkUpdateRequest) GetOpenUpdateDto() string {
	return this.OpenUpdateDto
}
func (this *OapiCircleEnworkUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.circle.enwork.update"
}
func (this *OapiCircleEnworkUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCircleEnworkUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCircleEnworkUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCircleEnworkUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCircleEnworkUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["open_update_dto"] = this.OpenUpdateDto
	return txtParams
}
func (this *OapiCircleEnworkUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiCircleEnworkUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCircleEnworkUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenUpdateWorkPostDTO struct {
	OpenId        string `json:"open_id,omitempty"`
	PigaiAnalysis string `json:"pigai_analysis,omitempty"`
	PostId        string `json:"post_id,omitempty"`
	Similarity    string `json:"similarity,omitempty"`
	VersionId     int64  `json:"version_id,omitempty"`
	WorkId        string `json:"work_id,omitempty"`
}
type OapiCircleEnworkUpdateResponse struct {
	taobao.TaobaoResponse
	Result bool `json:"result,omitempty"`
}
