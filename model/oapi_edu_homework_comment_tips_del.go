package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduHomeworkCommentTipsDeleteRequest() *OapiEduHomeworkCommentTipsDeleteRequest {
	return &OapiEduHomeworkCommentTipsDeleteRequest{}
}

type OapiEduHomeworkCommentTipsDeleteRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduHomeworkCommentTipsDeleteResponse
	BizCode         string
	TipIds          string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduHomeworkCommentTipsDeleteRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) SetBizCode(bizCode2 string) {
	this.BizCode = bizCode2
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) GetBizCode() string {
	return this.BizCode
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) SetTipIds(tipIds2 string) {
	this.TipIds = tipIds2
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) GetTipIds() string {
	return this.TipIds
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.homework.comment.tips.delete"
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["biz_code"] = this.BizCode
	txtParams["tip_ids"] = this.TipIds
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckStringMaxListSize(this.TipIds, 999, "tipIds"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduHomeworkCommentTipsDeleteRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduHomeworkCommentTipsDeleteResponse struct {
	taobao.TaobaoResponse
	Errcode int64  `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Result  bool   `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
