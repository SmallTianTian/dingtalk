package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsStatisticsJobListRequest() *OapiAtsStatisticsJobListRequest {
	return &OapiAtsStatisticsJobListRequest{}
}

type OapiAtsStatisticsJobListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsStatisticsJobListResponse
	BizCode         string
	Cursor          string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsStatisticsJobListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsStatisticsJobListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsStatisticsJobListRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsStatisticsJobListRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsStatisticsJobListRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiAtsStatisticsJobListRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiAtsStatisticsJobListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAtsStatisticsJobListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAtsStatisticsJobListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.statistics.job.list"
}
func (this *OapiAtsStatisticsJobListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsStatisticsJobListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsStatisticsJobListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsStatisticsJobListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsStatisticsJobListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiAtsStatisticsJobListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsStatisticsJobListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsStatisticsJobListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsStatisticsJobListResponse struct {
	taobao.TaobaoResponse
	Result AtsPageResult `json:"result,omitempty"`
}
