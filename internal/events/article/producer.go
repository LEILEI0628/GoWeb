package article

import "context"

type Producer interface {
	ProduceReadEvent(ctx context.Context, evt ReadEvent) error
}

type ReadEvent struct {
	Uid int64
	Aid int64
}
