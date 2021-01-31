package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduRecommendCreateRequest() *OapiEduRecommendCreateRequest {
	return &OapiEduRecommendCreateRequest{}
}

type OapiEduRecommendCreateRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduRecommendCreateResponse
	ClassId         int64
	OutContentId    string
	PeriodCode      string
	ReturnUrl       string
	SubjectCode     string
	Summary         string
	TextbookCode    string
	Thumbnail       string
	Title           string
	TopHttpMethod   string
	TopResponseType string
	Type            string
	Userid          string
}

func (this *OapiEduRecommendCreateRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduRecommendCreateRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduRecommendCreateRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduRecommendCreateRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduRecommendCreateRequest) SetOutContentId(outContentId2 string) {
	this.OutContentId = outContentId2
}
func (this *OapiEduRecommendCreateRequest) GetOutContentId() string {
	return this.OutContentId
}
func (this *OapiEduRecommendCreateRequest) SetPeriodCode(periodCode2 string) {
	this.PeriodCode = periodCode2
}
func (this *OapiEduRecommendCreateRequest) GetPeriodCode() string {
	return this.PeriodCode
}
func (this *OapiEduRecommendCreateRequest) SetReturnUrl(returnUrl2 string) {
	this.ReturnUrl = returnUrl2
}
func (this *OapiEduRecommendCreateRequest) GetReturnUrl() string {
	return this.ReturnUrl
}
func (this *OapiEduRecommendCreateRequest) SetSubjectCode(subjectCode2 string) {
	this.SubjectCode = subjectCode2
}
func (this *OapiEduRecommendCreateRequest) GetSubjectCode() string {
	return this.SubjectCode
}
func (this *OapiEduRecommendCreateRequest) SetSummary(summary2 string) {
	this.Summary = summary2
}
func (this *OapiEduRecommendCreateRequest) GetSummary() string {
	return this.Summary
}
func (this *OapiEduRecommendCreateRequest) SetTextbookCode(textbookCode2 string) {
	this.TextbookCode = textbookCode2
}
func (this *OapiEduRecommendCreateRequest) GetTextbookCode() string {
	return this.TextbookCode
}
func (this *OapiEduRecommendCreateRequest) SetThumbnail(thumbnail2 string) {
	this.Thumbnail = thumbnail2
}
func (this *OapiEduRecommendCreateRequest) GetThumbnail() string {
	return this.Thumbnail
}
func (this *OapiEduRecommendCreateRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiEduRecommendCreateRequest) GetTitle() string {
	return this.Title
}
func (this *OapiEduRecommendCreateRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *OapiEduRecommendCreateRequest) GetType() string {
	return this.Type
}
func (this *OapiEduRecommendCreateRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduRecommendCreateRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduRecommendCreateRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.recommend.create"
}
func (this *OapiEduRecommendCreateRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduRecommendCreateRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduRecommendCreateRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduRecommendCreateRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduRecommendCreateRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams["out_content_id"] = this.OutContentId
	txtParams["period_code"] = this.PeriodCode
	txtParams["return_url"] = this.ReturnUrl
	txtParams["subject_code"] = this.SubjectCode
	txtParams["summary"] = this.Summary
	txtParams["textbook_code"] = this.TextbookCode
	txtParams["thumbnail"] = this.Thumbnail
	txtParams["title"] = this.Title
	txtParams["type"] = this.Type
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduRecommendCreateRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OutContentId, "outContentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduRecommendCreateRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduRecommendCreateRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduRecommendCreateResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Success bool   `json:"success,omitempty"`
}
