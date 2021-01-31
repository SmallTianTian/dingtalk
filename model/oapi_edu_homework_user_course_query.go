package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduHomeworkUserCourseQueryRequest() *OapiEduHomeworkUserCourseQueryRequest {
	return &OapiEduHomeworkUserCourseQueryRequest{}
}

type OapiEduHomeworkUserCourseQueryRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkUserCourseQueryResponse
	BizCode         string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduHomeworkUserCourseQueryRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkUserCourseQueryRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkUserCourseQueryRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkUserCourseQueryRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkUserCourseQueryRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduHomeworkUserCourseQueryRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduHomeworkUserCourseQueryRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.user.course.query"
}
func (this *OapiEduHomeworkUserCourseQueryRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkUserCourseQueryRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkUserCourseQueryRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkUserCourseQueryRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkUserCourseQueryRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduHomeworkUserCourseQueryRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkUserCourseQueryRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkUserCourseQueryRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduHomeworkUserCourseQueryResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
type TemplateContentDto struct {
	Photo string `json:"photo,omitempty"`
	Title string `json:"title,omitempty"`
}
