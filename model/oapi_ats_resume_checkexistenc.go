package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsResumeCheckexistenceRequest() *OapiAtsResumeCheckexistenceRequest {
	return &OapiAtsResumeCheckexistenceRequest{}
}

type OapiAtsResumeCheckexistenceRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp             OapiAtsResumeCheckexistenceResponse
	BizCode          string
	ResumeDetailInfo string
	TopHttpMethod    string
	TopResponseType  string
}

func (this *OapiAtsResumeCheckexistenceRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsResumeCheckexistenceRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsResumeCheckexistenceRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsResumeCheckexistenceRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsResumeCheckexistenceRequest) SetResumeDetailInfo(resumeDetailInfo2 string) {
	this.ResumeDetailInfo = resumeDetailInfo2
}
func (this *OapiAtsResumeCheckexistenceRequest) GetResumeDetailInfo() string {
	return this.ResumeDetailInfo
}
func (this *OapiAtsResumeCheckexistenceRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.resume.checkexistence"
}
func (this *OapiAtsResumeCheckexistenceRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsResumeCheckexistenceRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsResumeCheckexistenceRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsResumeCheckexistenceRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsResumeCheckexistenceRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["resume_detail_info"] = this.ResumeDetailInfo
	return txtParams
}
func (this *OapiAtsResumeCheckexistenceRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsResumeCheckexistenceRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsResumeCheckexistenceRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsResumeCheckexistenceResponse struct {
	taobao.TaobaoResponse
	Result ResumeCheckResultVo `json:"result,omitempty"`
}
type ResumeCheckResultVo struct {
	Existed      bool     `json:"existed,omitempty"`
	ResumeIdList []string `json:"resume_id_list,omitempty"`
}
