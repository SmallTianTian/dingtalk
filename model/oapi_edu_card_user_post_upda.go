package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func NewOapiEduCardUserPostUpdateRequest() *OapiEduCardUserPostUpdateRequest {
	return &OapiEduCardUserPostUpdateRequest{}
}

type OapiEduCardUserPostUpdateRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduCardUserPostUpdateResponse
	TopHttpMethod   string
	TopResponseType string
	UpdatePostParam string
}

func (this *OapiEduCardUserPostUpdateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduCardUserPostUpdateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduCardUserPostUpdateRequest) SetUpdatePostParam(updatePostParam2 string) {
	this.UpdatePostParam = updatePostParam2
}
func (this *OapiEduCardUserPostUpdateRequest) GetUpdatePostParam() string {
	return this.UpdatePostParam
}
func (this *OapiEduCardUserPostUpdateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.card.user.post.update"
}
func (this *OapiEduCardUserPostUpdateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduCardUserPostUpdateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduCardUserPostUpdateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduCardUserPostUpdateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduCardUserPostUpdateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["update_post_param"] = this.UpdatePostParam
	return txtParams
}
func (this *OapiEduCardUserPostUpdateRequest) Check() *taobao.ApiRuleError {
	return nil
}
func (this *OapiEduCardUserPostUpdateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduCardUserPostUpdateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OpenUpdatePostParam struct {
	CardBizCode       string `json:"card_biz_code,omitempty"`
	CardBizId         string `json:"card_biz_id,omitempty"`
	CardId            string `json:"card_id,omitempty"`
	Content           string `json:"content,omitempty"`
	DetailUrl         string `json:"detail_url,omitempty"`
	EditUrl           string `json:"edit_url,omitempty"`
	Medias            string `json:"medias,omitempty"`
	MeteringNumber    string `json:"metering_number,omitempty"`
	PostId            int64  `json:"post_id,omitempty"`
	ReissueCard       bool   `json:"reissue_card,omitempty"`
	ShowName          string `json:"show_name,omitempty"`
	SourceType        string `json:"source_type,omitempty"`
	UnitOfMeasurement string `json:"unit_of_measurement,omitempty"`
	Userid            string `json:"userid,omitempty"`
}
type OapiEduCardUserPostUpdateResponse struct {
	taobao.TaobaoResponse
	Result  UpdatePostResponse `json:"result,omitempty"`
	Success bool               `json:"success,omitempty"`
}
type UpdatePostResponse struct {
	PostId int64 `json:"post_id,omitempty"`
}
