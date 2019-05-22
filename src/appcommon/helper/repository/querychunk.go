package repository

func (q *DbQuery) SelectQueryChunk(query string, value ...interface{}) *DbQuery {
	return q
}

// FromQueryChunk function concate "from" statement into sql query string
func (q *DbQuery) FromQueryChunk(name string) *DbQuery {
	*q = *q + " from " + DbQuery(name)
	return q
}

// MultiFromQueryCheck function concate "from" statement with
// multiple tables into sql query string
func (q *DbQuery) MultiFromQueryCheck(names ...string) *DbQuery {
	query := " from "
	for i, v := range names {
		query += v
		if i < len(names)-1 {
			query += ", "
		}
	}
	*q += DbQuery(query)
	return q
}

// FinishQuery function puts the stopper into sql query string
func (q *DbQuery) FinishQuery() *DbQuery {
	*q = *q + DbQuery(";")
	return q
}
