package tmp

import (
	"context"

	tmpModel "github.com/MrHuxu/xmock/tmp/model"
)

// Usage ...
type Usage interface {
	GetPage(context.Context, int, int, []string, string, string, ...interface{}) (results []*tmpModel.Usage, totalRows int64, err error)
	Count(ctx context.Context, cond string, args ...interface{}) (cnt int64, err error)
	GetAll(ctx context.Context, preloads []string, order, cond string, args ...interface{}) (results []*tmpModel.Usage, err error)
	GetOne(ctx context.Context, cond *tmpModel.Usage, preloads []string) (record *tmpModel.Usage, err error)
	Add(ctx context.Context, record *tmpModel.Usage) (result *tmpModel.Usage, err error)
	Update(ctx context.Context, cond, updates *tmpModel.Usage) (rowsAffected int64, err error)
	Upsert(ctx context.Context, cond, record *tmpModel.Usage) (result *tmpModel.Usage, err error)
	Delete(ctx context.Context, cond *tmpModel.Usage) (rowsAffected int64, err error)
}
