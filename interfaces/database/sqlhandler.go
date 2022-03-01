// 外側のレイヤーから持ってくる
// 登録しておく
package database

type SqlHandler interface {
    // Find(interface{}, ...interface{})
    // Exec(string, ...interface{})
    // First(interface{}, ...interface{})
    // Raw(string, ...interface{})
    // Create(interface{})
    // Save(interface{})
    // Delete(interface{})
    // Where(interface{}, ...interface{})
    Execute(string, ...interface{}) (Result, error)
    Query(string, ...interface{}) (Row, error)
}

type Result interface {
    LastInsertId() (int64, error)
    RowsAffected() (int64, error)
}

type Row interface {
    Scan(...interface{}) error
    Next() bool
    Close() error
}