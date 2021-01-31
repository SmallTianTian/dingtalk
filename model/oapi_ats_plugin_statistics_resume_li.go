package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsPluginStatisticsResumeListRequest() *OapiAtsPluginStatisticsResumeListRequest {
	return &OapiAtsPluginStatisticsResumeListRequest{}
}

type OapiAtsPluginStatisticsResumeListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsPluginStatisticsResumeListResponse
	BizCode         string
	Cursor          string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsPluginStatisticsResumeListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsPluginStatisticsResumeListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsPluginStatisticsResumeListRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsPluginStatisticsResumeListRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsPluginStatisticsResumeListRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiAtsPluginStatisticsResumeListRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiAtsPluginStatisticsResumeListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAtsPluginStatisticsResumeListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAtsPluginStatisticsResumeListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.plugin.statistics.resume.list"
}
func (this *OapiAtsPluginStatisticsResumeListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsPluginStatisticsResumeListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsPluginStatisticsResumeListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsPluginStatisticsResumeListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsPluginStatisticsResumeListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiAtsPluginStatisticsResumeListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsPluginStatisticsResumeListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsPluginStatisticsResumeListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsPluginStatisticsResumeListResponse struct {
	taobao.TaobaoResponse
	Errcode int64         `json:"errcode,omitempty"`
	Errmsg  string        `json:"errmsg,omitempty"`
	Result  AtsPageResult `json:"result,omitempty"`
}
type TopResumeStatisticsVo struct {
	BizCode         string `json:"biz_code,omitempty"`
	CandidateId     string `json:"candidate_id,omitempty"`
	Channel         string `json:"channel,omitempty"`
	CorpId          string `json:"corp_id,omitempty"`
	GmtCreateMils   int64  `json:"gmt_create_mils,omitempty"`
	GmtModifiedMils int64  `json:"gmt_modified_mils,omitempty"`
	ResumeId        string `json:"resume_id,omitempty"`
}
