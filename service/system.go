package service

import "context"

type SystemService struct {
	ctx context.Context
}

// NewSystemService 创建SystemService实例
func NewSystemService() *SystemService {
	return &SystemService{}
}

// Startup 初始化 获取wails上下文对象
func (s *SystemService) Startup(ctx context.Context) {
	s.ctx = ctx
}
