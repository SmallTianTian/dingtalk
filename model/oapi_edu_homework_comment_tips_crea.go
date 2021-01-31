package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduHomeworkCommentTipsCreateRequest() *OapiEduHomeworkCommentTipsCreateRequest {
	return &OapiEduHomeworkCommentTipsCreateRequest{}
}

type OapiEduHomeworkCommentTipsCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkCommentTipsCreateResponse
	Attributes      string
	Audio           string
	BizCode         string
	Content         string
	Media           string
	Photo           string
	SortOrder       int64
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduHomeworkCommentTipsCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) SetAttributes(attributes2 string) {
	this.Attributes = attributes2
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) GetAttributes() string {
	return this.Attributes
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) SetAudio(audio2 string) {
	this.Audio = audio2
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) GetAudio() string {
	return this.Audio
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) SetContent(content2 string) {
	this.Content = content2
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) GetContent() string {
	return this.Content
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) SetMedia(media2 string) {
	this.Media = media2
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) GetMedia() string {
	return this.Media
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) SetPhoto(photo2 string) {
	this.Photo = photo2
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) GetPhoto() string {
	return this.Photo
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) SetSortOrder(sortOrder2 int64) {
	this.SortOrder = sortOrder2
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) GetSortOrder() int64 {
	return this.SortOrder
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.comment.tips.create"
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["attributes"] = this.Attributes
	txtParams["audio"] = this.Audio
	txtParams["biz_code"] = this.BizCode
	txtParams[taobao.DATA_CONTENT] = this.Content
	txtParams["media"] = this.Media
	txtParams["photo"] = this.Photo
	txtParams["sort_order"] = this.SortOrder
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkCommentTipsCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduHomeworkCommentTipsCreateResponse struct {
	taobao.TaobaoResponse
	Result  int64 `json:"result,omitempty"`
	Success bool  `json:"success,omitempty"`
}
