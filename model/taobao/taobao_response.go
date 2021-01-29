package taobao

type TaobaoResponse struct {
	// http info
	Code          string            `json:"-"`
	HeaderContent map[string]string `json:"-"`
	Body          string            `json:"-"`

	// common info
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	RequestId string `json:"requsestid"`
}

func (tr *TaobaoResponse) IsSuccess() bool {
	return tr.ErrCode == 0
}
