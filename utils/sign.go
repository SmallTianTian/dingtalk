package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"

	"github.com/SmallTianTian/dingtalk/model/taobao"
)

func ComputeSignature(accessSecret, data string) (string, *taobao.ApiRuleError) {
	sig := hmac.New(sha256.New, []byte(accessSecret))
	if _, err := sig.Write([]byte(data)); err != nil {
		// TODO
		return "", &taobao.ApiRuleError{ErrCode: "1", ErrMsg: err.Error()}
	}
	// encode(value, encoding).replace("+", "%20").replace("*", "%2A").replace("~", "%7E")
	// 				.replace("/", "%2F");
	str := base64.StdEncoding.EncodeToString(sig.Sum(nil))
	return str, nil
	// return url.PathEscape(str), nil
}

func GetCanonicalStringForIsv(timestamp int64, suiteTicket string) string {
	if suiteTicket == "" {
		return Int64S(timestamp)
	}
	return fmt.Sprintf("%d\n%s", timestamp, suiteTicket)
}

func ParamToQueryString(m map[string]string) string {
	var pkv []string
	for k, v := range m {
		if k != "" && v != "" {
			pkv = append(pkv, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
		}
	}
	return strings.Join(pkv, "&")
}
