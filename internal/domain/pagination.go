package domain

type Pagination struct {
	Limit  int
	Offset int
}

type CreatePagination struct {
	First  string
	After  string
	Before string
}

type PageInfo struct {
	StartCursor     string
	EndCursor       string
	HasNextPage     string
	HasPreviousPage string
}

type Edges struct {
	Cursor string
	Notes  []interface{}
}

type PaginationResponse struct {
	Total int
	Edges
	PageInfo
}
