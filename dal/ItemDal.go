package dal

import "crud/model"

func InsertItem(item model.Item) (int64, int64, error) {
	GetConnection()
	sqlQuery := "Insert items set item_name=?, unit_price=?, amount=?, item_status=?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, 0, err
	}

	res, err := stmt.Exec(item.ItemName, item.UnitPrice, item.Amount, item.ItemStatus)

	if err != nil {
		return 0, 0, err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, 0, err
	}
	lastInsertId, err := res.LastInsertId()

	return rowsAffected, lastInsertId, err

}
