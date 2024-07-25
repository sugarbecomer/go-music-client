package dto

type KWSearchResp struct {
	ARTISTPIC     string    `json:"ARTISTPIC"`
	HIT           string    `json:"HIT"`
	HITMODE       string    `json:"HITMODE"`
	HITBUTOFFLINE string    `json:"HIT_BUT_OFFLINE"`
	MSHOW         string    `json:"MSHOW"`
	NEW           string    `json:"NEW"`
	PN            string    `json:"PN"`
	RN            string    `json:"RN"`
	SHOW          string    `json:"SHOW"`
	TOTAL         string    `json:"TOTAL"`
	UK            string    `json:"UK"`
	Abslist       []Abslist `json:"abslist"`
	Searchgroup   string    `json:"searchgroup"`
}

type Audiobookpayinfo struct {
	Download string `json:"download"`
	Play     string `json:"play"`
}

type Mvpayinfo struct {
	Download string `json:"download"`
	Play     string `json:"play"`
	Vid      string `json:"vid"`
}

type FeeType struct {
	Album   string `json:"album"`
	Bookvip string `json:"bookvip"`
	Song    string `json:"song"`
	Vip     string `json:"vip"`
}

type Paytagindex struct {
	AR501   int `json:"AR501"`
	DB      int `json:"DB"`
	F       int `json:"F"`
	H       int `json:"H"`
	HR      int `json:"HR"`
	L       int `json:"L"`
	S       int `json:"S"`
	ZP      int `json:"ZP"`
	ZPGA201 int `json:"ZPGA201"`
	ZPGA501 int `json:"ZPGA501"`
	ZPLY    int `json:"ZPLY"`
}

type PayInfo struct {
	CannotDownload   string      `json:"cannotDownload"`
	CannotOnlinePlay string      `json:"cannotOnlinePlay"`
	Download         string      `json:"download"`
	FeeType          FeeType     `json:"feeType"`
	Limitfree        string      `json:"limitfree"`
	ListenFragment   string      `json:"listen_fragment"`
	LocalEncrypt     string      `json:"local_encrypt"`
	Ndown            string      `json:"ndown"`
	Nplay            string      `json:"nplay"`
	OverseasNdown    string      `json:"overseas_ndown"`
	OverseasNplay    string      `json:"overseas_nplay"`
	Paytagindex      Paytagindex `json:"paytagindex"`
	Play             string      `json:"play"`
	RefrainEnd       string      `json:"refrain_end"`
	RefrainStart     string      `json:"refrain_start"`
	TipsIntercept    string      `json:"tips_intercept"`
}

type SUBLIST struct {
	AARTIST           string           `json:"AARTIST"`
	ALBUM             string           `json:"ALBUM"`
	ALBUMID           string           `json:"ALBUMID"`
	ARTIST            string           `json:"ARTIST"`
	ARTISTID          string           `json:"ARTISTID"`
	COPYRIGHT         string           `json:"COPYRIGHT"`
	CanSetRing        string           `json:"CanSetRing"`
	CanSetRingback    string           `json:"CanSetRingback"`
	DCTARGETID        string           `json:"DC_TARGETID"`
	DCTARGETTYPE      string           `json:"DC_TARGETTYPE"`
	DURATION          string           `json:"DURATION"`
	FORMATS           string           `json:"FORMATS"`
	HASECHO           string           `json:"HASECHO"`
	ISPOINT           string           `json:"IS_POINT"`
	MKVRID            string           `json:"MKVRID"`
	MP3NSIG1          string           `json:"MP3NSIG1"`
	MP3NSIG2          string           `json:"MP3NSIG2"`
	MP3RID            string           `json:"MP3RID"`
	MUSICRID          string           `json:"MUSICRID"`
	MUTIVER           string           `json:"MUTI_VER"`
	MVPIC             string           `json:"MVPIC"`
	NAME              string           `json:"NAME"`
	NEW               string           `json:"NEW"`
	NSIG1             string           `json:"NSIG1"`
	NSIG2             string           `json:"NSIG2"`
	ONLINE            string           `json:"ONLINE"`
	PAY               string           `json:"PAY"`
	PICPATH           string           `json:"PICPATH"`
	PLAYCNT           string           `json:"PLAYCNT"`
	SCORE100          string           `json:"SCORE100"`
	SIG1              string           `json:"SIG1"`
	SIG2              string           `json:"SIG2"`
	SONGNAME          string           `json:"SONGNAME"`
	SUBTITLE          string           `json:"SUBTITLE"`
	TAG               string           `json:"TAG"`
	AdSubtype         string           `json:"ad_subtype"`
	AdType            string           `json:"ad_type"`
	Allartistid       string           `json:"allartistid"`
	Audiobookpayinfo  Audiobookpayinfo `json:"audiobookpayinfo"`
	Barrage           string           `json:"barrage"`
	CacheStatus       string           `json:"cache_status"`
	ContentType       string           `json:"content_type"`
	Fpay              string           `json:"fpay"`
	Info              string           `json:"info"`
	IotInfo           string           `json:"iot_info"`
	Isdownload        string           `json:"isdownload"`
	Isshowtype        string           `json:"isshowtype"`
	Isstar            string           `json:"isstar"`
	Mp4Sig1           string           `json:"mp4sig1"`
	Mp4Sig2           string           `json:"mp4sig2"`
	Mvpayinfo         Mvpayinfo        `json:"mvpayinfo"`
	Originalsongtype  string           `json:"originalsongtype"`
	PayInfo           PayInfo          `json:"payInfo"`
	SpPrivilege       string           `json:"spPrivilege"`
	SubsStrategy      string           `json:"subsStrategy"`
	SubsText          string           `json:"subsText"`
	Terminal          string           `json:"terminal"`
	TmeMusicianAdtype string           `json:"tme_musician_adtype"`
	Tpay              string           `json:"tpay"`
	WebAlbumpicShort  string           `json:"web_albumpic_short"`
	WebArtistpicShort string           `json:"web_artistpic_short"`
	WebTimingonline   string           `json:"web_timingonline"`
}

type Abslist struct {
	AARTIST           string           `json:"AARTIST"`
	ALBUM             string           `json:"ALBUM"`
	ALBUMID           string           `json:"ALBUMID"`
	ARTIST            string           `json:"ARTIST"`
	ARTISTID          string           `json:"ARTISTID"`
	COPYRIGHT         string           `json:"COPYRIGHT"`
	CanSetRing        string           `json:"CanSetRing"`
	CanSetRingback    string           `json:"CanSetRingback"`
	DCTARGETID        string           `json:"DC_TARGETID"`
	DCTARGETTYPE      string           `json:"DC_TARGETTYPE"`
	DURATION          string           `json:"DURATION"`
	FORMATS           string           `json:"FORMATS"`
	HASECHO           string           `json:"HASECHO"`
	ISPOINT           string           `json:"IS_POINT"`
	MKVRID            string           `json:"MKVRID"`
	MP3NSIG1          string           `json:"MP3NSIG1"`
	MP3NSIG2          string           `json:"MP3NSIG2"`
	MP3RID            string           `json:"MP3RID"`
	MUSICRID          string           `json:"MUSICRID"`
	MUTIVER           string           `json:"MUTI_VER"`
	MVPIC             string           `json:"MVPIC"`
	NAME              string           `json:"NAME"`
	NEW               string           `json:"NEW"`
	NSIG1             string           `json:"NSIG1"`
	NSIG2             string           `json:"NSIG2"`
	ONLINE            string           `json:"ONLINE"`
	PAY               string           `json:"PAY"`
	PICPATH           string           `json:"PICPATH"`
	PLAYCNT           string           `json:"PLAYCNT"`
	SCORE100          string           `json:"SCORE100"`
	SIG1              string           `json:"SIG1"`
	SIG2              string           `json:"SIG2"`
	SONGNAME          string           `json:"SONGNAME"`
	SUBLIST           []SUBLIST        `json:"SUBLIST"`
	SUBTITLE          string           `json:"SUBTITLE"`
	TAG               string           `json:"TAG"`
	AdSubtype         string           `json:"ad_subtype"`
	AdType            string           `json:"ad_type"`
	Allartistid       string           `json:"allartistid"`
	Audiobookpayinfo  Audiobookpayinfo `json:"audiobookpayinfo"`
	Barrage           string           `json:"barrage"`
	CacheStatus       string           `json:"cache_status"`
	ContentType       string           `json:"content_type"`
	Fpay              string           `json:"fpay"`
	Info              string           `json:"info"`
	IotInfo           string           `json:"iot_info"`
	Isdownload        string           `json:"isdownload"`
	Isshowtype        string           `json:"isshowtype"`
	Isstar            string           `json:"isstar"`
	Mp4Sig1           string           `json:"mp4sig1"`
	Mp4Sig2           string           `json:"mp4sig2"`
	Mvpayinfo         Mvpayinfo        `json:"mvpayinfo"`
	Originalsongtype  string           `json:"originalsongtype"`
	PayInfo           PayInfo          `json:"payInfo"`
	SpPrivilege       string           `json:"spPrivilege"`
	SubsStrategy      string           `json:"subsStrategy"`
	SubsText          string           `json:"subsText"`
	Terminal          string           `json:"terminal"`
	TmeMusicianAdtype string           `json:"tme_musician_adtype"`
	Tpay              string           `json:"tpay"`
	WebAlbumpicShort  string           `json:"web_albumpic_short"`
	WebArtistpicShort string           `json:"web_artistpic_short"`
	WebTimingonline   string           `json:"web_timingonline"`
}
