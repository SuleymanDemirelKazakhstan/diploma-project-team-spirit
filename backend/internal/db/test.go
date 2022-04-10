package db

func (db *Layer) CheckPostgresVersion() (string, error) {
	var v string
	sqlStatement := `SELECT version()`
	row := db.DB.QueryRow(sqlStatement)
	if err := row.Scan(&v); err != nil {
		return "", err
	}
	return v, nil
}
