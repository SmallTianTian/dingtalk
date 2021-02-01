package model

import (
	"time"

	"github.com/SmallTianTian/dingtalk/model/taobao"
	"github.com/SmallTianTian/dingtalk/utils"
)

func NewOapiUserGetRequest() *OapiUserGetRequest {
	return &OapiUserGetRequest{}
}

type OapiUserGetRequest struct {
	taobao.SimpleTaobaoRequest
	taobao.SimpleResponse
	Resp            OapiUserGetResponse
	TopHttpMethod   string
	TopResponseType string
	Userid          string
}

func (this *OapiUserGetRequest) GetTopHttpMethod() string {
	return "GET"
}
func (this *OapiUserGetRequest) GetTopResponseType() string {
	return taobao.RESPONSE_TYPE_DINGTALK_OAPI
}
func (this *OapiUserGetRequest) SetUserid(userid2 string) {
	this.Userid = userid2
}
func (this *OapiUserGetRequest) GetUserid() string {
	return this.Userid
}
func (this *OapiUserGetRequest) GetApiMethodName() string {
	return "dingtalk.oapi.user.get"
}
func (this *OapiUserGetRequest) SetTopResponseType(topResponseType2 string) {
	this.TopResponseType = topResponseType2
}
func (this *OapiUserGetRequest) GetTopApiCallType() string {
	return "oapi"
}
func (this *OapiUserGetRequest) SetTopHttpMethod(topHttpMethod2 string) {
	this.TopHttpMethod = topHttpMethod2
}
func (this *OapiUserGetRequest) SetHttpMethod(httpMethod string) {
	this.SetTopHttpMethod(httpMethod)
}
func (this *OapiUserGetRequest) GetTextParams() map[string]interface{} {
	txtParams := make(map[string]interface{})
	txtParams[taobao.DATA_OUTGOING_USER_ID] = this.Userid
	return txtParams
}
func (this *OapiUserGetRequest) Check() *taobao.ApiRuleError {
	if err := utils.CheckNotEmpty(this.Userid, taobao.DATA_OUTGOING_USER_ID); err != nil {
		return err
	}
	return nil
}
func (this *OapiUserGetRequest) GetRespInstance() interface{} {
	return &this.Resp
}
func (this *OapiUserGetRequest) GetTaobaoResp() *taobao.TaobaoResponse {
	return &this.Resp.TaobaoResponse
}

type OapiUserGetResponse struct {
	taobao.TaobaoResponse
	Active            bool    `json:"active,omitempty"`
	AssociatedUnionId string  `json:"associatedUnionId,omitempty"`
	Avatar            string  `json:"avatar,omitempty"`
	Department        []int64 `json:"department,omitempty"`
	DingId            string  `json:"dingId,omitempty"`
	Email             string  `json:"email,omitempty"`

	Extattr            string    `json:"extattr,omitempty"`
	HiredDate          time.Time `json:"hiredDate,omitempty"`
	InviteMobile       string    `json:"inviteMobile,omitempty"`
	IsAdmin            bool      `json:"isAdmin,omitempty"`
	IsBoss             bool      `json:"isBoss,omitempty"`
	IsCustomizedPortal bool      `json:"isCustomizedPortal,omitempty"`
	IsHide             bool      `json:"isHide,omitempty"`
	IsLeaderInDepts    string    `json:"isLeaderInDepts,omitempty"`
	IsLimited          bool      `json:"isLimited,omitempty"`
	IsSenior           bool      `json:"isSenior,omitempty"`
	Jobnumber          string    `json:"jobnumber,omitempty"`
	ManagerUserId      string    `json:"managerUserId,omitempty"`
	MemberView         bool      `json:"memberView,omitempty"`
	Mobile             string    `json:"mobile,omitempty"`
	MobileHash         string    `json:"mobileHash,omitempty"`
	Name               string    `json:"name,omitempty"`
	Nickname           string    `json:"nickname,omitempty"`
	OpenId             string    `json:"openId,omitempty"`
	OrderInDepts       string    `json:"orderInDepts,omitempty"`
	OrgEmail           string    `json:"orgEmail,omitempty"`
	Position           string    `json:"position,omitempty"`
	RealAuthed         bool      `json:"realAuthed,omitempty"`
	Remark             string    `json:"remark,omitempty"`
	Roles              []Roles   `json:"roles,omitempty"`
	StateCode          string    `json:"stateCode,omitempty"`
	Tel                string    `json:"tel,omitempty"`
	Unionid            string    `json:"unionid,omitempty"`
	Userid             string    `json:"userid,omitempty"`
	WorkPlace          string    `json:"workPlace,omitempty"`
}
