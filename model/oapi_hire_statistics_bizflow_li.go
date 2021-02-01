package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHireStatisticsBizflowListRequest() *OapiHireStatisticsBizflowListRequest {
	return &OapiHireStatisticsBizflowListRequest{}
}

type OapiHireStatisticsBizflowListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHireStatisticsBizflowListResponse
	Cursor          string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiHireStatisticsBizflowListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHireStatisticsBizflowListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHireStatisticsBizflowListRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiHireStatisticsBizflowListRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiHireStatisticsBizflowListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiHireStatisticsBizflowListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiHireStatisticsBizflowListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hire.statistics.bizflow.list"
}
func (this *OapiHireStatisticsBizflowListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHireStatisticsBizflowListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHireStatisticsBizflowListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHireStatisticsBizflowListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHireStatisticsBizflowListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiHireStatisticsBizflowListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Size, "size"); err != nil {
		return err
	}
	return nil
}
func (this *OapiHireStatisticsBizflowListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHireStatisticsBizflowListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHireStatisticsBizflowListResponse struct {
	taobao.TaobaoResponse
	Result DdAtsPageResult `json:"result,omitempty"`
}
