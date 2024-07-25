package dto

type KWLyricResp struct {
	Code      int             `json:"code"`
	Msg       string          `json:"msg"`
	ReqID     string          `json:"reqId"`
	TID       string          `json:"tId"`
	Data      KWLyricRespData `json:"data"`
	ProfileID string          `json:"profileId"`
	CurTime   int64           `json:"curTime"`
	Success   bool            `json:"success"`
}

type Lrclist struct {
	LineLyric string `json:"lineLyric"`
	Time      string `json:"time"`
}

type KWLyricRespData struct {
	Lrclist []Lrclist `json:"lrclist"`
}
