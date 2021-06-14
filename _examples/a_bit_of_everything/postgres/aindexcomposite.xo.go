package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// AIndexComposite represents a row from 'public.a_index_composite'.
type AIndexComposite struct {
	AKey1 sql.NullInt64 `json:"a_key1"` // a_key1
	AKey2 sql.NullInt64 `json:"a_key2"` // a_key2

}

// AIndexCompositesByAKey1AKey2 retrieves a row from 'public.a_index_composite' as a AIndexComposite.
//
// Generated from index 'a_index_composite_idx'.
func AIndexCompositesByAKey1AKey2(ctx context.Context, db DB, aKey1, aKey2 sql.NullInt64) ([]*AIndexComposite, error) {
	// query
	const sqlstr = `SELECT ` +
		`a_key1, a_key2 ` +
		`FROM public.a_index_composite ` +
		`WHERE a_key1 = $1 AND a_key2 = $2`
	// run
	logf(sqlstr, aKey1, aKey2)
	rows, err := db.QueryContext(ctx, sqlstr, aKey1, aKey2)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*AIndexComposite
	for rows.Next() {
		aic := AIndexComposite{}
		// scan
		if err := rows.Scan(&aic.AKey1, &aic.AKey2); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &aic)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
