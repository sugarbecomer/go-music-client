package service

import "context"

type MusicClient interface {
	SearchMusic(ctx context.Context, keyword string) ([]string, error)
	GetMusicUrl(mid, quality string) string
	GetMusicUrl3(mid, quality string) string
}
