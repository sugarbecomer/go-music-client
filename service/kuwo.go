package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"go-music-client/cst"
	"go-music-client/dto"
	"go-music-client/re"
	"go-music-client/utils"
	"math"
	"net/url"
	"strings"
)

type KuwoClient struct {
	ctx       context.Context
	ArrayMask [64]int64
}

// NewKuwoClient 创建一个kuwo客户端
func NewKuwoClient() *KuwoClient {
	var client = &KuwoClient{}
	for i := 0; i < 63; i++ {
		client.ArrayMask[i] = int64(math.Pow(float64(2), float64(i)))
	}
	client.ArrayMask[len(client.ArrayMask)-1] = -9223372036854775808
	return client
}
func (k *KuwoClient) base64Encrypt(str string) string {
	bs := k.encrypt([]byte(str), []byte(SECRET_KEY))
	str = base64.StdEncoding.EncodeToString(bs)
	return strings.ReplaceAll(str, "\n", "")
}

func (k *KuwoClient) encrypt(msg, key []byte) []byte {

	var l int64
	for i := 0; i < 8; i++ {
		l = l | int64(key[i])<<(i*8)
	}
	j := len(msg) / 8

	var arrLong1 [16]int64
	k.subKeys(l, arrLong1[:], 0)

	arrLong2 := make([]int64, j)
	for m := 0; m < j; m++ {
		for n := 0; n < 8; n++ {
			arrLong2[m] |= int64(msg[n+m*8]) << (n * 8)
		}
	}

	arrLong3 := make([]int64, (1+8*(j+1))/8)
	for i := 0; i < j; i++ {
		arrLong3[i] = k.DES64(arrLong1[:], arrLong2[i])
	}

	arrByte1 := make([]byte, len(msg)-j*8)
	copy(arrByte1, msg[j*8:])
	var l2 int64
	msgLen := len(msg)
	for i := 0; i < msgLen%8; i++ {
		l2 |= int64(arrByte1[i]) << (i * 8)
	}
	arrLong3[j] = k.DES64(arrLong1[:], l2)

	arrByte2 := make([]byte, 8*len(arrLong3))
	var i4 int

	for _, l3 := range arrLong3 {
		for i6 := 0; i6 < 8; i6++ {
			arrByte2[i4] = 0xff & byte((l3 >> (i6 * 8)))
			i4 = i4 + 1
		}
	}
	return arrByte2
}

func (k *KuwoClient) DES64(ls []int64, l int64) int64 {
	var out int64
	var SOut int64
	var pR [8]int64
	var pSource [2]int64
	var L, R int64

	out = k.bitTransform(arrayIP, 64, l)
	pSource[0] = 0xFFFFFFFF & out
	pSource[1] = (int64(-4294967296) & out) >> 32

	for i := 0; i < 16; i++ {
		R = pSource[1]
		R = k.bitTransform(arrayE, 64, R)
		R ^= ls[i]
		for j := 0; j < 8; j++ {
			pR[j] = 255 & (R >> (j * 8))
		}
		SOut = 0
		for sbi := 7; sbi > -1; sbi-- {
			SOut <<= 4
			SOut |= int64(matrixNSBox[sbi][int(pR[sbi])])
		}

		R = k.bitTransform(arrayP, 32, SOut)
		L = pSource[0]
		pSource[0] = pSource[1]
		pSource[1] = L ^ R
	}

	pSs := ReverseSlice(pSource[:])
	out = int64(-4294967296)&(pSs[1]<<32) | 0xFFFFFFFF&pSs[0]
	out = k.bitTransform(arrayIP_1, 64, out)
	return out
}

// 反转切片
func ReverseSlice(s []int64) []int64 {
	var ns []int64
	for i := 0; i < len(s); i++ {
		ns = append(ns, s[len(s)-i-1])
	}
	return ns
}

func (k *KuwoClient) subKeys(l int64, ls []int64, n int) []int64 {
	l2 := k.bitTransform(arrayPC_1, 56, l)
	for i := 0; i < 16; i++ {
		l2 = (l2&int64(arrayLsMask[arrayLs[i]]))<<(28-int64(arrayLs[i])) | (l2&int64(^arrayLsMask[arrayLs[i]]))>>arrayLs[i]
		ls[i] = k.bitTransform(arrayPC_2, 64, l2)
	}
	var j int
	for n == 1 && j < 8 {
		var t = ls[j]
		ls[j] = ls[15-j]
		ls[15-j] = t
	}
	return ls
}

func (k *KuwoClient) bitTransform(arrInt []int, n int, l int64) int64 {
	var l2 int64
	for i := 0; i < n; i++ {
		if (arrInt[i] < 0) || ((l & k.ArrayMask[arrInt[i]]) == 0) {
			continue
		}
		l2 |= k.ArrayMask[i]
	}
	return l2
}
func KuWwSearchApi(pageNo, pageSize int, key string) string {
	escape := url.QueryEscape(key)
	return fmt.Sprintf(KWSearchApi, pageNo, pageSize, escape)
}

// SearchMusic 搜索音乐(实现musicClient interface)
func (k *KuwoClient) SearchMusic(kw string) *dto.KWResp {
	api := KuWwSearchApi(0, 10, kw)
	resp := utils.HttpGetWithHeader(api, KWSearchHead)
	if resp.Err != nil {
		panic(resp.Err)
	}
	log.Info(string(resp.Data))
	kwResp := new(dto.KWResp)
	err := json.Unmarshal(resp.Data, kwResp)
	if err != nil {
		panic(err)
	}
	return kwResp
}

// GetMusicUrl 获取音乐url(实现musicClient interface)
func (k *KuwoClient) GetMusicUrl(mid, quality string) string {
	switch quality {
	case cst.STAND:
		quality = "128kmp3"
	case cst.HIGHT:
		quality = "320kmp3"
	case cst.FLAC:
		quality = "2000kflac"
	}

	api := NewTrackUrlApi(mid, quality)
	resp := utils.HttpGetWithHeader(api, KWGetTrackUrlHead)
	if resp.SetCookie != "" {
		submatch := re.KWTokenReg.FindAllStringSubmatch(resp.SetCookie, -1)
		kwToken := submatch[0][1]
		csrf = kwToken
	}
	r := new(dto.KWTrackResp)
	err := json.Unmarshal(resp.Data, &r)
	if err != nil {
		panic(err)
	}
	if r.Data.URL == "" || strings.Contains(r.Data.URL, "/4141006416.mp3") ||
		strings.Contains(r.Data.URL, "/2272659253.mp3") ||
		strings.Contains(r.Data.URL, "2015636967.aac") {
		return k.GetMusicUrl3(mid, quality)
	}

	return r.Data.URL
}

// GetMusicUrl3 获取音乐url(实现musicClient interface)
func (k *KuwoClient) GetMusicUrl3(mid, quality string) string {
	switch quality {
	case cst.STAND:
		quality = "corp=kuwo&p2p=1&type=convert_url2&format=mp3&rid=" + mid
	case cst.HIGHT:
		quality = fmt.Sprintf("user=0&android_id=0&prod=kwplayer_ar_9.3.1.3&corp=kuwo&newver=3&vipver=9.3.1.3&source=oppo&p2p=1&notrace=0&type=convert_url2&format=flac|mp3|aac&sig=0&rid=%s&priority=bitrate&loginUid=0&network=WIFI&loginSid=0&mode=download", mid)
	case cst.FLAC:
		quality = "corp=kuwo&p2p=1&type=convert_url2&format=flac&rid=" + mid
	}

	api := NewTrackUrl3Api(k.base64Encrypt(quality))
	log.Info("api:", api)
	resp := utils.HttpGetWithHeader(api, KWGetTrackUrlHead)
	if resp.SetCookie != "" {
		submatch := re.KWTokenReg.FindAllStringSubmatch(resp.SetCookie, -1)
		kwToken := submatch[0][1]
		csrf = kwToken
	}

	str := string(resp.Data)
	allString := re.KWUrlReg.FindAllString(str, -1)
	if len(allString) == 0 {
		return ""
	}
	return allString[0]
}

func NewTrackUrl3Api(key string) string {
	return fmt.Sprintf(TrackUrl3Api, key)
}

func NewTrackUrlApi(mid, param string) string {
	return fmt.Sprintf(TrackUrlApi, mid, param)
}

const (
	SECRET_KEY = "ylzsxkwm"

	DES_MODE_DECRYPT int = 1

	KWSearchApi  = "https://search.kuwo.cn/r.s?pn=%d&rn=%d&all=%s&ft=music&newsearch=1&alflac=1&itemset=web_2013&client=kt&cluster=0&vermerge=1&rformat=json&encoding=utf8&show_copyright_off=1&pcmp4=1&ver=mbox&plat=pc&vipver=MUSIC_9.2.0.0_W6&devid=11404450&newver=1&issubtitle=1&pcjson=1"
	TrackUrl3Api = "http://nmobi.kuwo.cn/mobi.s?f=kuwo&q=%s"
	TrackUrlApi  = "https://mobi.kuwo.cn/mobi.s?f=web&source=kwplayer_ar_5.1.0.0_B_jiakong_vh.apk&type=convert_url_with_sign&rid=%s&br=%s"
)

var (
	arrayE = []int{
		31, 0, DES_MODE_DECRYPT, 2, 3, 4, -1, -1, 3, 4, 5, 6, 7, 8, -1, -1, 7, 8, 9, 10, 11, 12, -1, -1, 11, 12, 13, 14,
		15, 16, -1, -1, 15, 16, 17, 18, 19, 20, -1, -1, 19, 20, 21, 22, 23, 24, -1, -1, 23, 24, 25, 26, 27, 28, -1, -1,
		27, 28, 29, 30, 31, 30, -1, -1,
	}

	arrayIP = []int{
		57, 49, 41, 33, 25, 17, 9, DES_MODE_DECRYPT, 59, 51, 43, 35, 27, 19, 11, 3, 61, 53, 45, 37, 29, 21, 13, 5, 63,
		55, 47, 39, 31, 23, 15, 7, 56, 48, 40, 32, 24, 16, 8, 0, 58, 50, 42, 34, 26, 18, 10, 2, 60, 52, 44, 36, 28, 20,
		12, 4, 62, 54, 46, 38, 30, 22, 14, 6,
	}

	arrayIP_1 = []int{
		39, 7, 47, 15, 55, 23, 63, 31, 38, 6, 46, 14, 54, 22, 62, 30, 37, 5, 45, 13, 53, 21, 61, 29, 36, 4, 44, 12, 52,
		20, 60, 28, 35, 3, 43, 11, 51, 19, 59, 27, 34, 2, 42, 10, 50, 18, 58, 26, 33, DES_MODE_DECRYPT, 41, 9, 49, 17,
		57, 25, 32, 0, 40, 8, 48, 16, 56, 24,
	}

	arrayLs = []byte{1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1}

	arrayLsMask = []int{0, 0x100001, 0x300003}

	arrayP = []int{15, 6, 19, 20, 28, 11, 27, 16,
		0, 14, 22, 25, 4, 17, 30, 9,
		1, 7, 23, 13, 31, 26, 2, 8,
		18, 12, 29, 5, 21, 10, 3, 24}

	arrayPC_1 = []int{
		56, 48, 40, 32, 24, 16, 8, 0,
		57, 49, 41, 33, 25, 17, 9, 1,
		58, 50, 42, 34, 26, 18, 10, 2,
		59, 51, 43, 35, 62, 54, 46, 38,
		30, 22, 14, 6, 61, 53, 45, 37,
		29, 21, 13, 5, 60, 52, 44, 36,
		28, 20, 12, 4, 27, 19, 11, 3,
	}

	arrayPC_2 = []int{
		13, 16, 10, 23, 0, 4, -1, -1,
		2, 27, 14, 5, 20, 9, -1, -1,
		22, 18, 11, 3, 25, 7, -1, -1,
		15, 6, 26, 19, 12, 1, -1, -1,
		40, 51, 30, 36, 46, 54, -1, -1,
		29, 39, 50, 44, 32, 47, -1, -1,
		43, 48, 38, 55, 33, 52, -1, -1,
		45, 41, 49, 35, 28, 31, -1, -1,
	}

	matrixNSBox = [][]byte{{
		14, 4, 3, 15, 2, 13, 5, 3,
		13, 14, 6, 9, 11, 2, 0, 5,
		4, 1, 10, 12, 15, 6, 9, 10,
		1, 8, 12, 7, 8, 11, 7, 0,
		0, 15, 10, 5, 14, 4, 9, 10,
		7, 8, 12, 3, 13, 1, 3, 6,
		15, 12, 6, 11, 2, 9, 5, 0,
		4, 2, 11, 14, 1, 7, 8, 13}, {
		15, 0, 9, 5, 6, 10, 12, 9,
		8, 7, 2, 12, 3, 13, 5, 2,
		1, 14, 7, 8, 11, 4, 0, 3,
		14, 11, 13, 6, 4, 1, 10, 15,
		3, 13, 12, 11, 15, 3, 6, 0,
		4, 10, 1, 7, 8, 4, 11, 14,
		13, 8, 0, 6, 2, 15, 9, 5,
		7, 1, 10, 12, 14, 2, 5, 9}, {
		10, 13, 1, 11, 6, 8, 11, 5,
		9, 4, 12, 2, 15, 3, 2, 14,
		0, 6, 13, 1, 3, 15, 4, 10,
		14, 9, 7, 12, 5, 0, 8, 7,
		13, 1, 2, 4, 3, 6, 12, 11,
		0, 13, 5, 14, 6, 8, 15, 2,
		7, 10, 8, 15, 4, 9, 11, 5,
		9, 0, 14, 3, 10, 7, 1, 12}, {
		7, 10, 1, 15, 0, 12, 11, 5,
		14, 9, 8, 3, 9, 7, 4, 8,
		13, 6, 2, 1, 6, 11, 12, 2,
		3, 0, 5, 14, 10, 13, 15, 4,
		13, 3, 4, 9, 6, 10, 1, 12,
		11, 0, 2, 5, 0, 13, 14, 2,
		8, 15, 7, 4, 15, 1, 10, 7,
		5, 6, 12, 11, 3, 8, 9, 14}, {
		2, 4, 8, 15, 7, 10, 13, 6,
		4, 1, 3, 12, 11, 7, 14, 0,
		12, 2, 5, 9, 10, 13, 0, 3,
		1, 11, 15, 5, 6, 8, 9, 14,
		14, 11, 5, 6, 4, 1, 3, 10,
		2, 12, 15, 0, 13, 2, 8, 5,
		11, 8, 0, 15, 7, 14, 9, 4,
		12, 7, 10, 9, 1, 13, 6, 3}, {
		12, 9, 0, 7, 9, 2, 14, 1,
		10, 15, 3, 4, 6, 12, 5, 11,
		1, 14, 13, 0, 2, 8, 7, 13,
		15, 5, 4, 10, 8, 3, 11, 6,
		10, 4, 6, 11, 7, 9, 0, 6,
		4, 2, 13, 1, 9, 15, 3, 8,
		15, 3, 1, 14, 12, 5, 11, 0,
		2, 12, 14, 7, 5, 10, 8, 13}, {
		4, 1, 3, 10, 15, 12, 5, 0,
		2, 11, 9, 6, 8, 7, 6, 9,
		11, 4, 12, 15, 0, 3, 10, 5,
		14, 13, 7, 8, 13, 14, 1, 2,
		13, 6, 14, 9, 4, 1, 2, 14,
		11, 13, 5, 0, 1, 10, 8, 3,
		0, 11, 3, 5, 9, 4, 15, 2,
		7, 8, 12, 15, 10, 7, 6, 12}, {
		13, 7, 10, 0, 6, 9, 5, 15,
		8, 4, 3, 10, 11, 14, 12, 5,
		2, 11, 9, 6, 15, 12, 0, 3,
		4, 1, 14, 13, 1, 2, 7, 8,
		1, 2, 12, 15, 10, 4, 0, 3,
		13, 14, 6, 9, 7, 8, 9, 6,
		15, 1, 5, 12, 3, 10, 14, 5,
		8, 7, 11, 0, 4, 13, 2, 11}}

	csrf = ""
)
var (
	KWSearchHead = map[string]string{
		"user_agent": `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.50`,
		"accept":     `text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7`,
		"referer":    `http://kuwo.cn/search/list?key=%E6%81%90%E9%BE%99%E6%8A%97%E7%8B%BC8`,
		"csrf":       csrf,
		"Secret":     "13261c0dccfeac48dd7a8b33de9fd1bb59e7bcd1fbda77ed3a2e42bce5fc7e0f0036d507",
		"Cross":      "e5191b2eb629a3da9dc6868755a3e779",
		"Cookie":     "ga=GA1.2.1860922824.1635265329; Hm_lvt_cdb524f42f0ce19b169a8071123a4797=1663159268; gid=9ed7ed0b-8d4b-4167-8c9d-f1f2c55642f7; Hm_token=et7csP3xeQfeadZsDEazXEpEXhmjTC4k; Hm_Iuvt_cdb524f42f0ce19b169b8072123a4727=Mzfa6zAAcAfszyHFdREYF7KfBRNmAEi4",
	}

	KWGetTrackUrlHead = map[string]string{
		"user_agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.50",
		"accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"referer":         "https://www.kuwo.cn/search/list?key=",
		"csrf":            csrf,
		"accept_encoding": "gzip",
		"host":            "nmobi.kuwo.cn",
		"connection":      "Keep-Alive",
	}
)
