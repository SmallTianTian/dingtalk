package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiHirePluginStatisticsBizflowListRequest() *OapiHirePluginStatisticsBizflowListRequest {
	return &OapiHirePluginStatisticsBizflowListRequest{}
}

type OapiHirePluginStatisticsBizflowListRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiHirePluginStatisticsBizflowListResponse
	Cursor          string
	Size            int64
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiHirePluginStatisticsBizflowListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiHirePluginStatisticsBizflowListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiHirePluginStatisticsBizflowListRequest) SetCursor(cursor2 string) {
	this.Cursor = cursor2
}
func (this *OapiHirePluginStatisticsBizflowListRequest) GetCursor() string {
	return this.Cursor
}
func (this *OapiHirePluginStatisticsBizflowListRequest) SetSize(size2 int64) {
	this.Size = size2
}
func (this *OapiHirePluginStatisticsBizflowListRequest) GetSize() int64 {
	return this.Size
}
func (this *OapiHirePluginStatisticsBizflowListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.hire.plugin.statistics.bizflow.list"
}
func (this *OapiHirePluginStatisticsBizflowListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiHirePluginStatisticsBizflowListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiHirePluginStatisticsBizflowListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiHirePluginStatisticsBizflowListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiHirePluginStatisticsBizflowListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["cursor"] = this.Cursor
	txtParams["size"] = this.Size
	return txtParams
}
func (this *OapiHirePluginStatisticsBizflowListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Size, "size"); err != nil {
		return err
	}
	return nil
}
func (this *OapiHirePluginStatisticsBizflowListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiHirePluginStatisticsBizflowListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiHirePluginStatisticsBizflowListResponse struct {
	taobao.TaobaoResponse
	Result DdAtsPageResult `json:"result,omitempty"`
}
type TopBizFlowStatisticsVo struct {
	CandidateId     string `json:"candidate_id,omitempty"`
	CorpId          string `json:"corp_id,omitempty"`
	CreatorUserid   string `json:"creator_userid,omitempty"`
	FlowId          string `json:"flow_id,omitempty"`
	FlowStatus      int64  `json:"flow_status,omitempty"`
	GmtCreateMils   int64  `json:"gmt_create_mils,omitempty"`
	GmtModifiedMils int64  `json:"gmt_modified_mils,omitempty"`
	JobId           string `json:"job_id,omitempty"`
	OwnerUserid     string `json:"owner_userid,omitempty"`
	RecruitId       string `json:"recruit_id,omitempty"`
	ResumeId        string `json:"resume_id,omitempty"`
}
type DdAtsPageResult struct {
	HasMore    bool                     `json:"has_more,omitempty"`
	List       []TopBizFlowStatisticsVo `json:"list,omitempty"`
	NextCursor string                   `json:"next_cursor,omitempty"`
}
