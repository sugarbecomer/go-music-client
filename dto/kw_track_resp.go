package dto

type KWTrackResp struct {
	Code       int             `json:"code"`
	Locationid string          `json:"locationid"`
	Data       KWTrackRespData `json:"data"`
	Msg        string          `json:"msg"`
}

type KWTrackRespData struct {
	Bitrate          int    `json:"bitrate"`
	User             string `json:"user"`
	Sig              string `json:"sig"`
	Type             string `json:"type"`
	Format           string `json:"format"`
	P2PAudiosourceid string `json:"p2p_audiosourceid"`
	Rid              int    `json:"rid"`
	Source           string `json:"source"`
	URL              string `json:"url"`
}
