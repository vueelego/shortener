package dbrepo

type QuotaRepo interface {
}

var _ QuotaRepo = (*quotaRepo)(nil)

type quotaRepo struct{}
