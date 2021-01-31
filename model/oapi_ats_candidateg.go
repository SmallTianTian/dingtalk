package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiAtsCandidateGetRequest() *OapiAtsCandidateGetRequest {
	return &OapiAtsCandidateGetRequest{}
}

type OapiAtsCandidateGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiAtsCandidateGetResponse
	BizCode         string
	CandidateId     string
	TopHttpMethod   string
	TopResponseType string
}

func (this *OapiAtsCandidateGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiAtsCandidateGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiAtsCandidateGetRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiAtsCandidateGetRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiAtsCandidateGetRequest) SetCandidateId(candidateId2 string) {
	this.CandidateId = candidateId2
}
func (this *OapiAtsCandidateGetRequest) GetCandidateId() string {
	return this.CandidateId
}
func (this *OapiAtsCandidateGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.ats.candidate.get"
}
func (this *OapiAtsCandidateGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiAtsCandidateGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiAtsCandidateGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiAtsCandidateGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiAtsCandidateGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["candidate_id"] = this.CandidateId
	return txtParams
}
func (this *OapiAtsCandidateGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.BizCode, "bizCode"); err != nil {
		return err
	}
	return nil
}
func (this *OapiAtsCandidateGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiAtsCandidateGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiAtsCandidateGetResponse struct {
	taobao.TaobaoResponse
	Errcode int64          `json:"errcode,omitempty"`
	Errmsg  string         `json:"errmsg,omitempty"`
	Result  TopCandidateVO `json:"result,omitempty"`
}
type TopCandidateVO struct {
	CandidateId     string `json:"candidate_id,omitempty"`
	CorpId          string `json:"corp_id,omitempty"`
	GmtCreateMils   int64  `json:"gmt_create_mils,omitempty"`
	GmtModifiedMils int64  `json:"gmt_modified_mils,omitempty"`
	Name            string `json:"name,omitempty"`
}
