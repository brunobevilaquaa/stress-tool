package cli

import "stress-tool/internal/services"

type RootAdapter struct {
	service services.StressServiceInterface
}

func NewRootAdapter(service services.StressServiceInterface) *RootAdapter {
	return &RootAdapter{
		service: service,
	}
}

func (r *RootAdapter) Run(url string, requests int, concurrency int) error {
	return r.service.Run(url, requests, concurrency)
}
