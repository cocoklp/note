IPV4:
^([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(\.([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])){3}$

IPV6:
^(?i)[1-9a-f][0-9a-f]{0,3}(:[0-9a-f]{1,4}){7}$
  (?i): 匹配时不区分大小写





```
go run reg.go 'SDK-HMAC-SHA256 Access=dddddddddddddddddddd, SignedHeaders=1234567123456712345671234567, Signature=1234567812345678123456781234567812345678123456781234567812345678'


func parseAuthHeader(header string) (ak string, signHeader string, sig string) {

	defer func() {
		if err1 := recover(); err1 != nil {
			log.Error("panic handled:", err1)
		}
	}()
	authRegexp := regexp.MustCompile(util.AuthHeaderRegex)
	matchVars := authRegexp.FindStringSubmatch(header)
	if len(matchVars) <= util.MaxMatchVarSize {
		return "", "", ""
	}
	return matchVars[1], matchVars[2], matchVars[3]
}
const AuthHeaderRegex string = `^SDK-HMAC-SHA256 Access=([\w=+/]{20}), SignedHeaders=([^, ]{28}), Signature=([^, ]{64})$`

```

