package resultx

const (
	defaultLimit = 10
	maxLimit     = 100
)

type PaginationRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type Pagination struct {
	Total      int64 `json:"total"`
	Page       int   `json:"page"`
	TotalPages int   `json:"totalPages"`
	Limit      int   `json:"limit"`
	Offset     int   `json:"offset"`
	HasNext    bool  `json:"hasNext"`
	HasPrev    bool  `json:"hasPrev"`
}

func NewPaginationRequest(offset int, limit int) PaginationRequest {
	if offset < 0 {
		offset = 0
	}
	if limit < 1 {
		limit = defaultLimit
	}
	if limit > maxLimit {
		limit = maxLimit
	}
	return PaginationRequest{
		Offset: offset,
		Limit:  limit,
	}
}

func (p PaginationRequest) GetOffset() int {
	return p.Offset
}

func (p PaginationRequest) GetLimit() int {
	return p.Limit
}

func NewPagination(total int64, offset int, limit int) Pagination {
	page := 1
	totalPages := 0
	if limit > 0 {
		page = (offset / limit) + 1
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}

	hasNext := int64(offset+limit) < total
	hasPrev := offset > 0

	return Pagination{
		Total:      total,
		Page:       page,
		TotalPages: totalPages,
		Offset:     offset,
		Limit:      limit,
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}
}
