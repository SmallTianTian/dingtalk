package model

import (
	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiEduUserGetRequest() *OapiEduUserGetRequest {
	return &OapiEduUserGetRequest{}
}

type OapiEduUserGetRequest struct {
	taobao.TaobaoRequest
	taobao.SimpleResponse
	Resp            OapiEduUserGetResponse
	ClassId         int64
	Role            string
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiEduUserGetRequest) GetTopHttpMethod() string {
	return "POST"
}
func (this *OapiEduUserGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiEduUserGetRequest) SetClassId(classId2 int64) {
	this.ClassId = classId2
}
func (this *OapiEduUserGetRequest) GetClassId() int64 {
	return this.ClassId
}
func (this *OapiEduUserGetRequest) SetRole(role2 string) {
	this.Role = role2
}
func (this *OapiEduUserGetRequest) GetRole() string {
	return this.Role
}
func (this *OapiEduUserGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiEduUserGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiEduUserGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.edu.user.get"
}
func (this *OapiEduUserGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiEduUserGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiEduUserGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiEduUserGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiEduUserGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams["class_id"] = this.ClassId
	txtParams["role"] = this.Role
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiEduUserGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.ClassId, "classId"); err != nil {
		return err
	}
	return nil
}
func (this *OapiEduUserGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiEduUserGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiEduUserGetResponse struct {
	taobao.TaobaoResponse
	Result  Result `json:"result,omitempty"`
	Success bool   `json:"success,omitempty"`
}
type Detail struct {
	ClassId int64  `json:"class_id,omitempty"`
	Feature string `json:"feature,omitempty"`
	Name    string `json:"name,omitempty"`
	Role    string `json:"role,omitempty"`
	Unionid string `json:"unionid,omitempty"`
	Userid  string `json:"userid,omitempty"`
}
