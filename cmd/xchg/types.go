package main

import "database/sql"

// Order 代表用戶的外幣資產的變動。
//
// 由於 Order 是站在外幣資產的角度，所以買進外幣時 Foreign 會是正的，賣出則會是負的。
// 同理，買進外幣時的 Local 會是負的，賣出則會是正的。
type Order struct {
	Time    int64   `sdm:"order_when" json:"when"` // 交易的 timestamp
	Foreign float64 `sdm:"foreign_currency" json:"foreign"`
	Local   float64 `sdm:"local_currency" json:"local"`
	Code    string  `sdm:"currency_code" json:"code"` // 外幣代碼 (三碼英文)
}

// This will be called at very begining of the program, just fail fast.
func initTable(db *sql.DB) error {
	qstr := `CREATE TABLE IF NOT EXISTS orders (
  order_when INTEGER,
  foreign_currency DOUBLE,
  local_currency DOUBLE,
  currency_code VARCHAR(3),
  CONSTRAINT order_pk PRIMARY KEY (order_when,currency_code)
)`

	if _, err := db.Exec(qstr); err != nil {
		return err
	}

	return nil
}
