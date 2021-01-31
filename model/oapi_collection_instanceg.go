package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
	"time"
)

func NewOapiCollectionInstanceGetRequest() *OapiCollectionInstanceGetRequest {
	return &OapiCollectionInstanceGetRequest{}
}

type OapiCollectionInstanceGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCollectionInstanceGetResponse
	BizType         int64
	FormInstanceId  string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCollectionInstanceGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCollectionInstanceGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCollectionInstanceGetRequest) SetBizType(bizType2 int64) {
	this.BizType = bizType2
}
func (this *OapiCollectionInstanceGetRequest) GetBizType() int64 {
	return this.BizType
}
func (this *OapiCollectionInstanceGetRequest) SetFormInstanceId(formInstanceId2 string) {
	this.FormInstanceId = formInstanceId2
}
func (this *OapiCollectionInstanceGetRequest) GetFormInstanceId() string {
	return this.FormInstanceId
}
func (this *OapiCollectionInstanceGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.collection.instance.get"
}
func (this *OapiCollectionInstanceGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCollectionInstanceGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCollectionInstanceGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCollectionInstanceGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCollectionInstanceGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_type"] = this.BizType
	txtParams["formInstance_id"] = this.FormInstanceId
	return txtParams
}
func (this *OapiCollectionInstanceGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.FormInstanceId, "formInstanceId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCollectionInstanceGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCollectionInstanceGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCollectionInstanceGetResponse struct {
	taobao.TaobaoResponse
	Result FormInstanceVo `json:"result,omitempty"`
}
type FormData struct {
	Key   string `json:"key,omitempty"`
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}
type FormInstanceVo struct {
	CreateTime time.Time  `json:"create_time,omitempty"`
	Creator    string     `json:"creator,omitempty"`
	FormCode   string     `json:"form_code,omitempty"`
	FormList   []FormData `json:"form_list,omitempty"`
	ModifyTime time.Time  `json:"modify_time,omitempty"`
	Title      string     `json:"title,omitempty"`
}
