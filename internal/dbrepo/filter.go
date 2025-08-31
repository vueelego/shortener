package dbrepo

import "math"

type Filter struct {
	PageNum     int   `json:"pageNum"`
	PageSize    int   `json:"pageSize"`
	TotalRows   int64 `json:"totalRows"`
	LastPageNum int   `json:"lastPageNum"`
}

func (f *Filter) Limit() int {
	if f.PageSize < 10 {
		f.PageSize = 10
	}
	if f.PageSize > 100 {
		f.PageSize = 100
	}

	return f.PageSize
}

func (f *Filter) Offset() int {
	if f.PageNum < 1 {
		f.PageNum = 1
	}

	return (f.PageNum - 1) * f.PageSize
}

func (f *Filter) GetLastPageNum() int {
	if f.TotalRows <= 0 {
		return 0
	} else {
		return int(math.Ceil(float64(f.TotalRows) / float64(f.PageSize)))
	}
}
