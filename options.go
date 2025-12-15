package resultx

type MetaOption func(*Metadata)

func WithPagination(pagination Pagination) MetaOption {
	return func(m *Metadata) {
		m.Pagination = &pagination
	}
}
