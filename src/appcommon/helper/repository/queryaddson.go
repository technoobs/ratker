package repository

// Action --> Scope

// DbQuery is the raw sql query statement
type DbQuery string

// BaseQuery is the base sql query
type BaseQuery struct {
	Query DbQuery
	Error error
}

// BaseSearchQuery is the base sql query used for SEARCH
type BaseSearchQuery struct {
	Query       BaseQuery
	SelectField string // select
	Table       string // from
	Scope       string // where
}

// BaseUpdateQuery is the base sql query used for UPDATE
type BaseUpdateQuery struct {
	Query       BaseQuery
	Table       string // update
	UpdateField string // set
	Scope       string // where
}

// BaseDeleteQuery is the base sql query used for DELETE
type BaseDeleteQuery struct {
	Query BaseQuery
	Table string // delete from
	Scope string // where
}

// BaseInsertQuery is the base sql query used for INSERT
type BaseInsertQuery struct {
	Query       BaseQuery
	Table       string
	InsertField string
	InsertValue string
}

// TODO: Create extension functions to build custome raw query string

// GetQuery function sets the basic query for selected table
func (query *BaseQuery) GetQuery(name string) *BaseQuery {
	return query
}
