package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRobotMessageStatisticsListbypushidRequest() *OapiRobotMessageStatisticsListbypushidRequest {
	return &OapiRobotMessageStatisticsListbypushidRequest{}
}

type OapiRobotMessageStatisticsListbypushidRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRobotMessageStatisticsListbypushidResponse
	ConversationIds string
	Page            int64
	PageSize        int64
	PushId          string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRobotMessageStatisticsListbypushidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) SetConversationIds(conversationIds2 string) {
	this.ConversationIds = conversationIds2
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) GetConversationIds() string {
	return this.ConversationIds
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) SetPage(page2 int64) {
	this.Page = page2
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) GetPage() int64 {
	return this.Page
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) SetPushId(pushId2 string) {
	this.PushId = pushId2
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) GetPushId() string {
	return this.PushId
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.robot.message.statistics.listbypushid"
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["conversation_ids"] = this.ConversationIds
	txtParams["page"] = this.Page
	txtParams["page_size"] = this.PageSize
	txtParams["push_id"] = this.PushId
	return txtParams
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ConversationIds, "conversationIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRobotMessageStatisticsListbypushidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRobotMessageStatisticsListbypushidResponse struct {
	taobao.TaobaoResponse
	Errcode int64                    `json:"errcode,omitempty"`
	Errmsg  string                   `json:"errmsg,omitempty"`
	Result  RobotMsgStatPageResponse `json:"result,omitempty"`
	Success bool                     `json:"success,omitempty"`
}
type MsgStatByPushIdResVo struct {
	ConversationId         string `json:"conversation_id,omitempty"`
	GroupMemberCount       int64  `json:"group_member_count,omitempty"`
	GroupMemberUnreadCount int64  `json:"group_member_unread_count,omitempty"`
	PushId                 string `json:"push_id,omitempty"`
}
