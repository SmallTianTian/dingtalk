package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduHomeworkTopicCreateRequest() *OapiEduHomeworkTopicCreateRequest {
	return &OapiEduHomeworkTopicCreateRequest{}
}

type OapiEduHomeworkTopicCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkTopicCreateResponse
	TopHttpMethod   string
	TopResponseType string
	TopicItems      string
}

func (this *OapiEduHomeworkTopicCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkTopicCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkTopicCreateRequest) SetTopicItems(topicItems2 string) {
	this.TopicItems = topicItems2
}
func (this *OapiEduHomeworkTopicCreateRequest) GetTopicItems() string {
	return this.TopicItems
}
func (this *OapiEduHomeworkTopicCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.topic.create"
}
func (this *OapiEduHomeworkTopicCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkTopicCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkTopicCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkTopicCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkTopicCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["topic_items"] = this.TopicItems
	return txtParams
}
func (this *OapiEduHomeworkTopicCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckObjectMaxListSize(this.TopicItems, 20, "topicItems"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkTopicCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkTopicCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type SyncTopicItem struct {
	Analysis        string `json:"analysis,omitempty"`
	Answer          string `json:"answer,omitempty"`
	Attributes      string `json:"attributes,omitempty"`
	BizCode         string `json:"biz_code,omitempty"`
	Point           string `json:"point,omitempty"`
	QuestionContent string `json:"question_content,omitempty"`
	QuestionId      string `json:"question_id,omitempty"`
	QuestionType    string `json:"question_type,omitempty"`
}
type OapiEduHomeworkTopicCreateResponse struct {
	taobao.TaobaoResponse
	Result  []int64 `json:"result,omitempty"`
	Success bool    `json:"success,omitempty"`
}
