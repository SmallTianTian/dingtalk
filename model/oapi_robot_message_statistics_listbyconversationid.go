package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiRobotMessageStatisticsListbyconversationidRequest() *OapiRobotMessageStatisticsListbyconversationidRequest {
	return &OapiRobotMessageStatisticsListbyconversationidRequest{}
}

type OapiRobotMessageStatisticsListbyconversationidRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiRobotMessageStatisticsListbyconversationidResponse
	ConversationIds string
	Page            int64
	PageSize        int64
	PushId          string
	ReadStatus      bool
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiRobotMessageStatisticsListbyconversationidRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) SetConversationIds(conversationIds2 string) {
	this.ConversationIds = conversationIds2
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) GetConversationIds() string {
	return this.ConversationIds
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) SetPage(page2 int64) {
	this.Page = page2
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) GetPage() int64 {
	return this.Page
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) SetPageSize(pageSize2 int64) {
	this.PageSize = pageSize2
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) GetPageSize() int64 {
	return this.PageSize
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) SetPushId(pushId2 string) {
	this.PushId = pushId2
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) GetPushId() string {
	return this.PushId
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) SetReadStatus(readStatus2 bool) {
	this.ReadStatus = readStatus2
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) GetReadStatus() bool {
	return this.ReadStatus
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) GetApiMethodName() string {
	return "dingtalk.oapi.robot.message.statistics.listbyconversationid"
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["conversation_ids"] = this.ConversationIds
	txtParams["page"] = this.Page
	txtParams["page_size"] = this.PageSize
	txtParams["push_id"] = this.PushId
	txtParams["read_status"] = this.ReadStatus
	return txtParams
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ConversationIds, "conversationIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiRobotMessageStatisticsListbyconversationidRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiRobotMessageStatisticsListbyconversationidResponse struct {
	taobao.TaobaoResponse
	Errcode int64                    `json:"errcode,omitempty"`
	Errmsg  string                   `json:"errmsg,omitempty"`
	Result  RobotMsgStatPageResponse `json:"result,omitempty"`
	Success bool                     `json:"success,omitempty"`
}
type MsgStatByCIdResVo struct {
	ConversationId string `json:"conversation_id,omitempty"`
	DingtalkId     string `json:"dingtalk_id,omitempty"`
	Name           string `json:"name,omitempty"`
	PushId         string `json:"push_id,omitempty"`
	ReadStatus     bool   `json:"read_status,omitempty"`
}
