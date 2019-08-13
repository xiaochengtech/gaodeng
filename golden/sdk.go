package golden

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var BaseUrl = map[string]string{"test": "http://182.254.219.106:8400", "prod": "https://openapi.fapiaoer.cn"}

func NewSdk(appkey, appsecret, env string) *Sdk {
	return &Sdk{
		env:       env,
		appkey:    appkey,
		appsecret: appsecret,
	}
}

type Sdk struct {
	env       string
	appkey    string
	appsecret string
}

func (this *Sdk) HttpPost(url string, post map[string]interface{}) (*http.Response, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	url = this.getBaseUrl() + url + "?appkey=" + this.appkey + "&timestamp=" + timestamp + "&signature=" + this.generateSign(post, timestamp)
	json := jsoniter.Config{
		MarshalFloatWith6Digits: true,
		EscapeHTML:              false,
		SortMapKeys:             true, //本身高灯平台仅要求对最外层json key进行asci码升序排序，但map是无序且随机的，所以签名和post数据均排序以保持一致
		UseNumber:               true,
		DisallowUnknownFields:   false,
		CaseSensitive:           true,
	}.Froze()
	b, _ := json.Marshal(post)
	return http.Post(url, "application/json", bytes.NewBuffer(b))
}

func (this *Sdk) generateSign(post map[string]interface{}, timestamp string) string {
	var originStrBuilder bytes.Buffer
	originStrBuilder.WriteString(this.appkey)
	originStrBuilder.WriteString(timestamp)
	var keys []string
	for key, _ := range post {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var valueStrBuilder bytes.Buffer
	var keyLen = len(keys)
	for i := 0; i < keyLen; i++ {
		key := keys[i]
		valueStrBuilder.WriteString(key)
		valueStrBuilder.WriteString("=")
		value := post[key]
		switch value.(type) {
		case string:
			valueStrBuilder.WriteString(value.(string))
		case int:
			v := value.(int)
			valueStrBuilder.WriteString(strconv.Itoa(v))
		case bool:
			if value.(bool) {
				valueStrBuilder.WriteString("1")
			} else {
				valueStrBuilder.WriteString("0")
			}
		default:
			json := jsoniter.Config{
				MarshalFloatWith6Digits: true,
				EscapeHTML:              false,
				SortMapKeys:             true, //本身高灯平台仅要求对最外层json key进行asci码升序排序，但map是无序且随机的，所以签名和post数据均排序以保持一致
				UseNumber:               true,
				DisallowUnknownFields:   false,
				CaseSensitive:           true,
			}.Froze()
			s, _ := json.Marshal(value)
			valueStrBuilder.WriteString(string(s))
		}
		if i < keyLen-1 {
			valueStrBuilder.WriteString("&")
		}
	}
	valueString := url.QueryEscape(valueStrBuilder.String())
	valueString = strings.Replace(valueString, "+", "%20", -1)
	originStrBuilder.WriteString(valueString)
	//fmt.Println(valueString)
	originStrBuilder.WriteString(this.appsecret)
	h := md5.New()
	//fmt.Println( string( originStrBuilder.Bytes() ))
	h.Write(originStrBuilder.Bytes())
	md5Value := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	//fmt.Println(md5Value)
	return md5Value
}

func (this *Sdk) getBaseUrl() string {
	if this.env == "test" {
		return BaseUrl["test"]
	} else {
		return BaseUrl["prod"]
	}
}
