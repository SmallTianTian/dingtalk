package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRobotMessageStatisticsListRequest() *OapiRobotMessageStatisticsListRequest {
	return &OapiRobotMessageStatisticsListRequest{}
}

type OapiRobotMessageStatisticsListRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRobotMessageStatisticsListResponse
	Page            int64
	PageSize        int64
	PushIds         string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRobotMessageStatisticsListRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRobotMessageStatisticsListRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRobotMessageStatisticsListRequest) SetPage(page2 int64) {
	this.Page = page2
}
func (this *OapiRobotMessageStatisticsListRequest) GetPage() int64 {
	return this.Page
}
func (this *OapiRobotMessageStatisticsListRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiRobotMessageStatisticsListRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiRobotMessageStatisticsListRequest) SetPushIds(pushIds2 string) {
	this.PushIds = pushIds2
}
func (this *OapiRobotMessageStatisticsListRequest) GetPushIds() string {
	return this.PushIds
}
func (this *OapiRobotMessageStatisticsListRequest) GetApiMethodName() string {
	return "dingtalk.oapi.robot.message.statistics.list"
}
func (this *OapiRobotMessageStatisticsListRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRobotMessageStatisticsListRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRobotMessageStatisticsListRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRobotMessageStatisticsListRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRobotMessageStatisticsListRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["page"] = this.Page
	txtParams["page_size"] = this.PageSize
	txtParams["push_ids"] = this.PushIds
	return txtParams
}
func (this *OapiRobotMessageStatisticsListRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Page, "page"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRobotMessageStatisticsListRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRobotMessageStatisticsListRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRobotMessageStatisticsListResponse struct {
	taobao.TaobaoResponse
	Errcode int64                    `json:"errcode,omitempty"`
	Errmsg  string                   `json:"errmsg,omitempty"`
	Result  RobotMsgStatPageResponse `json:"result,omitempty"`
	Success bool                     `json:"success,omitempty"`
}
type MsgStatResVo struct {
	PushId                      string `json:"push_id,omitempty"`
	ReachGroupCount             int64  `json:"reach_group_count,omitempty"`
	ReachGroupMemberCount       int64  `json:"reach_group_member_count,omitempty"`
	ReachGroupMemberUnreadCount int64  `json:"reach_group_member_unread_count,omitempty"`
}
type PaginationVo struct {
	CurrentPage    int64 `json:"current_page,omitempty"`
	TotalPage      int64 `json:"total_page,omitempty"`
	TotalRecordNum int64 `json:"total_record_num,omitempty"`
}
type RobotMsgStatPageResponse struct {
	List       []MsgStatResVo `json:"list,omitempty"`
	Pagination PaginationVo   `json:"pagination,omitempty"`
}
