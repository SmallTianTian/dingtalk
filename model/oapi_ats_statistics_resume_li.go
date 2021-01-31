package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsStatisticsResumeListRequest() *OapiAtsStatisticsResumeListRequest {
	return &OapiAtsStatisticsResumeListRequest{}
}

type OapiAtsStatisticsResumeListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsStatisticsResumeListResponse
	BizCode         string
	Cursor          string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsStatisticsResumeListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsStatisticsResumeListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsStatisticsResumeListRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsStatisticsResumeListRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsStatisticsResumeListRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiAtsStatisticsResumeListRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiAtsStatisticsResumeListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAtsStatisticsResumeListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAtsStatisticsResumeListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.statistics.resume.list"
}
func (this *OapiAtsStatisticsResumeListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsStatisticsResumeListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsStatisticsResumeListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsStatisticsResumeListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsStatisticsResumeListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiAtsStatisticsResumeListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsStatisticsResumeListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsStatisticsResumeListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsStatisticsResumeListResponse struct {
	taobao.TaobaoResponse
	Errcode int64         `json:"errcode,omitempty"`
	Errmsg  string        `json:"errmsg,omitempty"`
	Result  AtsPageResult `json:"result,omitempty"`
}
