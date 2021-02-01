package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduRecommendReturnRequest() *OapiEduRecommendReturnRequest {
	return &OapiEduRecommendReturnRequest{}
}

type OapiEduRecommendReturnRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduRecommendReturnResponse
	ClassId         int64
	OutContentId    string
	OutTxId         string
	ResultType      int64
	ResultValue     string
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

func (this *OapiEduRecommendReturnRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduRecommendReturnRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduRecommendReturnRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduRecommendReturnRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduRecommendReturnRequest) SetOutContentId(outContentId2 string) {
	this.OutContentId = outContentId2
}
func (this *OapiEduRecommendReturnRequest) GetOutContentId() string {
	return this.OutContentId
}
func (this *OapiEduRecommendReturnRequest) SetOutTxId(outTxId2 string) {
	this.OutTxId = outTxId2
}
func (this *OapiEduRecommendReturnRequest) GetOutTxId() string {
	return this.OutTxId
}
func (this *OapiEduRecommendReturnRequest) SetResultType(resultType2 int64) {
	this.ResultType = resultType2
}
func (this *OapiEduRecommendReturnRequest) GetResultType() int64 {
	return this.ResultType
}
func (this *OapiEduRecommendReturnRequest) SetResultValue(resultValue2 string) {
	this.ResultValue = resultValue2
}
func (this *OapiEduRecommendReturnRequest) GetResultValue() string {
	return this.ResultValue
}
func (this *OapiEduRecommendReturnRequest) SetReturnUrl(returnUrl2 string) {
	this.ReturnUrl = returnUrl2
}
func (this *OapiEduRecommendReturnRequest) GetReturnUrl() string {
	return this.ReturnUrl
}
func (this *OapiEduRecommendReturnRequest) SetSubjectCode(subjectCode2 string) {
	this.SubjectCode = subjectCode2
}
func (this *OapiEduRecommendReturnRequest) GetSubjectCode() string {
	return this.SubjectCode
}
func (this *OapiEduRecommendReturnRequest) SetSummary(summary2 string) {
	this.Summary = summary2
}
func (this *OapiEduRecommendReturnRequest) GetSummary() string {
	return this.Summary
}
func (this *OapiEduRecommendReturnRequest) SetTextbookCode(textbookCode2 string) {
	this.TextbookCode = textbookCode2
}
func (this *OapiEduRecommendReturnRequest) GetTextbookCode() string {
	return this.TextbookCode
}
func (this *OapiEduRecommendReturnRequest) SetThumbnail(thumbnail2 string) {
	this.Thumbnail = thumbnail2
}
func (this *OapiEduRecommendReturnRequest) GetThumbnail() string {
	return this.Thumbnail
}
func (this *OapiEduRecommendReturnRequest) SetTitle(title2 string) {
	this.Title = title2
}
func (this *OapiEduRecommendReturnRequest) GetTitle() string {
	return this.Title
}
func (this *OapiEduRecommendReturnRequest) SetType(type2 string) {
	this.Type = type2
}
func (this *OapiEduRecommendReturnRequest) GetType() string {
	return this.Type
}
func (this *OapiEduRecommendReturnRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduRecommendReturnRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduRecommendReturnRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.recommend.return"
}
func (this *OapiEduRecommendReturnRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduRecommendReturnRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduRecommendReturnRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduRecommendReturnRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduRecommendReturnRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams["out_content_id"] = this.OutContentId
	txtParams["out_tx_id"] = this.OutTxId
	txtParams["result_type"] = this.ResultType
	txtParams["result_value"] = this.ResultValue
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
func (this *OapiEduRecommendReturnRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.OutContentId, "outContentId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduRecommendReturnRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduRecommendReturnRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduRecommendReturnResponse struct {
	taobao.TaobaoResponse
	Success bool `json:"success,omitempty"`
}
