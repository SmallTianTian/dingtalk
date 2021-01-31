package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiCollectionFormListRequest() *OapiCollectionFormListRequest {
	return &OapiCollectionFormListRequest{}
}

type OapiCollectionFormListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiCollectionFormListResponse
	BizType         int64
	Creator         string
	Offset          int64
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiCollectionFormListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiCollectionFormListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiCollectionFormListRequest) SetBizType(bizType2 int64) {
	this.BizType = bizType2
}
func (this *OapiCollectionFormListRequest) GetBizType() int64 {
	return this.BizType
}
func (this *OapiCollectionFormListRequest) SetCreator(creator2 string) {
	this.Creator = creator2
}
func (this *OapiCollectionFormListRequest) GetCreator() string {
	return this.Creator
}
func (this *OapiCollectionFormListRequest) SetOffset(offset2 int64) {
	this.Offset = offset2
}
func (this *OapiCollectionFormListRequest) GetOffset() int64 {
	return this.Offset
}
func (this *OapiCollectionFormListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiCollectionFormListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiCollectionFormListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.collection.form.list"
}
func (this *OapiCollectionFormListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiCollectionFormListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiCollectionFormListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiCollectionFormListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiCollectionFormListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_type"] = this.BizType
	txtParams["creator"] = this.Creator
	txtParams["offset"] = this.Offset
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiCollectionFormListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Offset, "offset"); err != nil {
		return err
	}
	return nil
}
func (this *OapiCollectionFormListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiCollectionFormListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiCollectionFormListResponse struct {
	taobao.TaobaoResponse
	Result PageResult `json:"result,omitempty"`
}
type FormSchemaResponse struct {
	Creator  string              `json:"creator,omitempty"`
	FormCode string              `json:"form_code,omitempty"`
	Memo     string              `json:"memo,omitempty"`
	Name     string              `json:"name,omitempty"`
	Setting  FormSchemaSettingVo `json:"setting,omitempty"`
}
