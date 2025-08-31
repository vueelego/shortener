package dbrepo

type SessionRepo interface {
}

var _ SessionRepo = (*sessionRepo)(nil)

type sessionRepo struct{}
