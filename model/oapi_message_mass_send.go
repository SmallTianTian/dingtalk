package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiMessageMassSendRequest() *OapiMessageMassSendRequest {
	return &OapiMessageMassSendRequest{}
}

type OapiMessageMassSendRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiMessageMassSendResponse
	DepIdList       string
	IsPreview       bool
	IsToAll         bool
	MediaId         string
	MsgBody         string
	MsgType         string
	TextContent     string
	TopHttpMethod   string
	TopResponseType string
	Unionid         string
	UseridList      string
	Uuid            string
}

func (this *OapiMessageMassSendRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiMessageMassSendRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiMessageMassSendRequest) SetDepIdList(depIdList2 string) {
	this.DepIdList = depIdList2
}
func (this *OapiMessageMassSendRequest) GetDepIdList() string {
	return this.DepIdList
}
func (this *OapiMessageMassSendRequest) SetIsPreview(isPreview2 bool) {
	this.IsPreview = isPreview2
}
func (this *OapiMessageMassSendRequest) GetIsPreview() bool {
	return this.IsPreview
}
func (this *OapiMessageMassSendRequest) SetIsToAll(isToAll2 bool) {
	this.IsToAll = isToAll2
}
func (this *OapiMessageMassSendRequest) GetIsToAll() bool {
	return this.IsToAll
}
func (this *OapiMessageMassSendRequest) SetMediaId(mediaId2 string) {
	this.MediaId = mediaId2
}
func (this *OapiMessageMassSendRequest) GetMediaId() string {
	return this.MediaId
}
func (this *OapiMessageMassSendRequest) SetMsgBody(msgBody2 string) {
	this.MsgBody = msgBody2
}
func (this *OapiMessageMassSendRequest) GetMsgBody() string {
	return this.MsgBody
}
func (this *OapiMessageMassSendRequest) SetMsgType(msgType2 string) {
	this.MsgType = msgType2
}
func (this *OapiMessageMassSendRequest) GetMsgType() string {
	return this.MsgType
}
func (this *OapiMessageMassSendRequest) SetTextContent(textContent2 string) {
	this.TextContent = textContent2
}
func (this *OapiMessageMassSendRequest) GetTextContent() string {
	return this.TextContent
}
func (this *OapiMessageMassSendRequest) SetUnionid(unionid2 string) {
	this.Unionid = unionid2
}
func (this *OapiMessageMassSendRequest) GetUnionid() string {
	return this.Unionid
}
func (this *OapiMessageMassSendRequest) SetUseridList(useridList2 string) {
	this.UseridList = useridList2
}
func (this *OapiMessageMassSendRequest) GetUseridList() string {
	return this.UseridList
}
func (this *OapiMessageMassSendRequest) SetUuid(uuid2 string) {
	this.Uuid = uuid2
}
func (this *OapiMessageMassSendRequest) GetUuid() string {
	return this.Uuid
}
func (this *OapiMessageMassSendRequest) GetApiMethodName() string {
	return "dingtalk.oapi.message.mass.send"
}
func (this *OapiMessageMassSendRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiMessageMassSendRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiMessageMassSendRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiMessageMassSendRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiMessageMassSendRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["dep_id_list"] = this.DepIdList
	txtParams["is_preview"] = this.IsPreview
	txtParams["is_to_all"] = this.IsToAll
	txtParams["media_id"] = this.MediaId
	txtParams["msg_body"] = this.MsgBody
	txtParams["msg_type"] = this.MsgType
	txtParams["text_content"] = this.TextContent
	txtParams["unionid"] = this.Unionid
	txtParams["userid_list"] = this.UseridList
	txtParams["uuid"] = this.Uuid
	return txtParams
}
func (this *OapiMessageMassSendRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.DepIdList, 2000, "depIdList"); err != nil {
		return err
	}
	return nil
}
func (this *OapiMessageMassSendRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiMessageMassSendRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type FormItem struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
type RichText struct {
	Num  string `json:"num,omitempty"`
	Unit string `json:"unit,omitempty"`
}
type Button struct {
	ActionUrl string `json:"action_url,omitempty"`
	Title     string `json:"title,omitempty"`
}
type MessageBody struct {
	ActionCard ActionCard `json:"action_card,omitempty"`
	File       File       `json:"file,omitempty"`
	Link       Link       `json:"link,omitempty"`
	Markdown   Markdown   `json:"markdown,omitempty"`
	Oa         OA         `json:"oa,omitempty"`
	Voice      Voice      `json:"voice,omitempty"`
}
type OapiMessageMassSendResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	TaskId  string `json:"task_id,omitempty"`
}
