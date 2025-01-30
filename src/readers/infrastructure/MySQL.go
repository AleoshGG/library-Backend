package infrastructure

import (
	"fmt"
	"library-Backend/src/core"
	"library-Backend/src/readers/domain"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()

	if conn.Err != "" {
		fmt.Println("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) CreateReader(reader domain.Reader) (uint, error) {
	query := "INSERT INTO readers (first_name, last_name, email, phone_number, account_status) VALUES (?,?,?,?,?)"

	res, err := mysql.conn.ExecutePreparedQuery(query, reader.First_name, reader.Last_name, reader.Email, reader.Phone_number, reader.Account_status)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}

	id, _ := res.LastInsertId() 

	return uint(id), nil 
}

func (mysql *MySQL) GetReaderByName(name string) []domain.Reader {
	query := "SELECT * FROM readers WHERE first_name LIKE CONCAT('%', ?, '%') OR last_name LIKE CONCAT('%', ?, '%')"

	var readers []domain.Reader

	rows := mysql.conn.FetchRows(query, name, name)

	if rows == nil {
        fmt.Println("No se pudieron obtener los datos.")
        return readers
    }

	defer rows.Close()

	for rows.Next() {
		var r domain.Reader
		rows.Scan(&r.Id_reader, &r.First_name, &r.Last_name, &r.Email, &r.Phone_number, &r.Account_status)

		readers = append(readers, r)
	}
	
	return readers
}

func (mysql *MySQL) GetAllReaders() []domain.Reader {
	query := "SELECT * FROM readers LIMIT 100"

	var readers []domain.Reader

	rows := mysql.conn.FetchRows(query)

	if rows == nil {
        fmt.Println("No se pudieron obtener los datos.")
        return readers
    }

	defer rows.Close()

	for rows.Next() {
		var r domain.Reader
		rows.Scan(&r.Id_reader, &r.First_name, &r.Last_name, &r.Email, &r.Phone_number, &r.Account_status)

		readers = append(readers, r)
	}
	
	return readers
}