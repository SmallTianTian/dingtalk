package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsPluginStatisticsJobListRequest() *OapiAtsPluginStatisticsJobListRequest {
	return &OapiAtsPluginStatisticsJobListRequest{}
}

type OapiAtsPluginStatisticsJobListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsPluginStatisticsJobListResponse
	BizCode         string
	Cursor          string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsPluginStatisticsJobListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsPluginStatisticsJobListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsPluginStatisticsJobListRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsPluginStatisticsJobListRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsPluginStatisticsJobListRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiAtsPluginStatisticsJobListRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiAtsPluginStatisticsJobListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiAtsPluginStatisticsJobListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiAtsPluginStatisticsJobListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.plugin.statistics.job.list"
}
func (this *OapiAtsPluginStatisticsJobListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsPluginStatisticsJobListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsPluginStatisticsJobListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsPluginStatisticsJobListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsPluginStatisticsJobListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiAtsPluginStatisticsJobListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsPluginStatisticsJobListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsPluginStatisticsJobListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsPluginStatisticsJobListResponse struct {
	taobao.TaobaoResponse
	Errcode int64         `json:"errcode,omitempty"`
	Errmsg  string        `json:"errmsg,omitempty"`
	Result  AtsPageResult `json:"result,omitempty"`
}
type TopJobStatisticsVo struct {
	BizCode         string `json:"biz_code,omitempty"`
	CorpId          string `json:"corp_id,omitempty"`
	CreatorUserid   string `json:"creator_userid,omitempty"`
	GmtCreateMils   int64  `json:"gmt_create_mils,omitempty"`
	GmtModifiedMils int64  `json:"gmt_modified_mils,omitempty"`
	HeadCount       int64  `json:"head_count,omitempty"`
	JobId           string `json:"job_id,omitempty"`
	MainDeptId      int64  `json:"main_dept_id,omitempty"`
	Name            string `json:"name,omitempty"`
	OwnerUserid     string `json:"owner_userid,omitempty"`
	Status          int64  `json:"status,omitempty"`
}
