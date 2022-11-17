package abstractions

type GetQueries struct {
	Pagination   Pagination
	Associations string `query:"associations" json:"associations,omitempty"`   // json string
	Filter       string `query:"filter" json:"filter,omitempty" bson:"filter"` // json string
	Sort         string `query:"sort" json:"sort,omitempty"`
}

type GetByIdQueries struct {
	Associations string `query:"associations" json:"associations,omitempty"` // json string
}
