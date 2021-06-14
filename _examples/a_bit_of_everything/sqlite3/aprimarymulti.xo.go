package sqlite3

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// APrimaryMulti represents a row from 'a_primary_multi'.
type APrimaryMulti struct {
	AKey  int            `json:"a_key"`  // a_key
	AText sql.NullString `json:"a_text"` // a_text
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the APrimaryMulti exists in the database.
func (apm *APrimaryMulti) Exists() bool {
	return apm._exists
}

// Deleted returns true when the APrimaryMulti has been marked for deletion from
// the database.
func (apm *APrimaryMulti) Deleted() bool {
	return apm._deleted
}

// Insert inserts the APrimaryMulti to the database.
func (apm *APrimaryMulti) Insert(ctx context.Context, db DB) error {
	switch {
	case apm._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case apm._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO a_primary_multi (` +
		`a_text` +
		`) VALUES (` +
		`?` +
		`)`
	// run
	logf(sqlstr, apm.AText)
	res, err := db.ExecContext(ctx, sqlstr, apm.AText)
	if err != nil {
		return err
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	} // set primary key
	apm.AKey = int(id)
	// set exists
	apm._exists = true
	return nil
}

// Update updates a APrimaryMulti in the database.
func (apm *APrimaryMulti) Update(ctx context.Context, db DB) error {
	switch {
	case !apm._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case apm._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE a_primary_multi SET ` +
		`a_text = ?` +
		` WHERE a_key = ?`
	// run
	logf(sqlstr, apm.AText, apm.AKey)
	if _, err := db.ExecContext(ctx, sqlstr, apm.AText, apm.AKey); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the APrimaryMulti to the database.
func (apm *APrimaryMulti) Save(ctx context.Context, db DB) error {
	if apm.Exists() {
		return apm.Update(ctx, db)
	}
	return apm.Insert(ctx, db)
}

// Delete deletes the APrimaryMulti from the database.
func (apm *APrimaryMulti) Delete(ctx context.Context, db DB) error {
	switch {
	case !apm._exists: // doesn't exist
		return nil
	case apm._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM a_primary_multi WHERE a_key = ?`
	// run
	logf(sqlstr, apm.AKey)
	if _, err := db.ExecContext(ctx, sqlstr, apm.AKey); err != nil {
		return logerror(err)
	}
	// set deleted
	apm._deleted = true
	return nil
}

// APrimaryMultiByAKey retrieves a row from 'a_primary_multi' as a APrimaryMulti.
//
// Generated from index 'a_primary_multi_a_key_pkey'.
func APrimaryMultiByAKey(ctx context.Context, db DB, aKey int) (*APrimaryMulti, error) {
	// query
	const sqlstr = `SELECT ` +
		`a_key, a_text ` +
		`FROM a_primary_multi ` +
		`WHERE a_key = ?`
	// run
	logf(sqlstr, aKey)
	apm := APrimaryMulti{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, aKey).Scan(&apm.AKey, &apm.AText); err != nil {
		return nil, logerror(err)
	}
	return &apm, nil
}
