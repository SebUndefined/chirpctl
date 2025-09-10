package shared

import "fmt"

type ListResult[T any] struct {
	TotalCount uint32
	Items      []T
}

func NewListResult[T any]() ListResult[T] {
	return ListResult[T]{
		TotalCount: 0,
		Items:      []T{},
	}
}

func NewListResultCap[T any](capHint int) ListResult[T] {
	if capHint < 0 {
		capHint = 0
	}
	return ListResult[T]{
		Items: make([]T, 0, capHint),
	}
}

// Add -
func (lr *ListResult[T]) Add(item T) {
	lr.Items = append(lr.Items, item)
}

// Merge -
func (lr *ListResult[T]) Merge(other ListResult[T]) {
	lr.Items = append(lr.Items, other.Items...)
	if other.TotalCount > lr.TotalCount { // TotalCount = “total serveur”
		lr.TotalCount = other.TotalCount // (ex.: pagination)
	}
}

func (lr *ListResult[T]) AddMany(items ...T) {
	lr.Items = append(lr.Items, items...)
}

func (lr ListResult[T]) Len() int {
	return len(lr.Items)
}

func (lr ListResult[T]) IsEmpty() bool {
	return len(lr.Items) == 0
}

type ListPaginationRequest struct {
	Limit  uint32
	Offset uint32
	Search string
}

func (lr ListPaginationRequest) String() string {
	return fmt.Sprintf("limit=%d, offset=%d search=%s", lr.Limit, lr.Offset, lr.Search)
}

func (lr ListPaginationRequest) GetSafeLimit() uint32 {
	if lr.Limit == 0 {
		return 100
	}
	if lr.Limit > 1000 {
		return 1000
	}
	return lr.Limit
}

type ListOrderByRequest struct {
	OrderBy     string
	OrderByDesc bool
}

type ListTagsFilterRequest struct {
	Tags map[string]string
}
