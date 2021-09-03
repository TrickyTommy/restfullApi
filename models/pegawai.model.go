package models

import (
	"net/http"
	"restFullApi/db"
)

type Pegawai struct {
	ID int `json:"id"`
	Nama string `json:"nama"`
	Alamat string `json:"alamat"`
	Telepon string `json:"telepon"`
}

func FetchAllPegawai()  (Response, error){

	var obj Pegawai
	var arrobj []Pegawai
	var res Response

	// connection

	con := db.CreateCon()


	//sql statement

	 sqlStatment := "SELECT * FROM pegawai"

	 rows, err := con.Query(sqlStatment)
	 defer rows.Close()

	if err != nil {
		return res, err

	}

	for rows.Next() {
		err = rows.Scan(&obj.ID, &obj.Nama, &obj.Alamat, &obj.Telepon)
		if err != nil {
			return res, err

		}

		arrobj = append(arrobj, obj)


	}

	res.Status = http.StatusOK

	res.Message = "Success"
	res.Data = arrobj

	return res,nil
}

// input data ke 1 pegawai
func StorePegawai(nama,alamat,telepon string) (Response,error) {
	var res Response

	con := db.CreateCon()

	sqlStatement :=  "INSERT pegawai (nama, alamat, telepon) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res,err

	}

	result, err := stmt.Exec(nama,alamat,telepon)
	if err != nil {
		return res,err
	}

	lastInsertedId, er := result.LastInsertId()
	if er != nil {
		return res,err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id" :lastInsertedId,
	}
	return res,nil
}
//langkah 1update pegawai
func UpdatePegawai(id int, nama,alamat,telepon string)(Response, error)  {
	var res Response

	con := db.CreateCon()

	sqlStatment := "UPDATE pegawai SET nama = ?, alamat = ?, telepon = ? WHERE id = ?"

	stmt ,err := con.Prepare(sqlStatment)
	if err != nil {
		return res, err

	}
	result, err := stmt.Exec(nama,alamat,telepon,id)
	if err != nil {
		return res,err
	}

	//rows effected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}
	return res,nil
}

func DeletePegawai(id int)  (Response, error){

	var res Response
	 con := db.CreateCon()

	 sqlStatement := "DELETE FROM pegawai WHERE id = ?"

	 stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected , err := result.RowsAffected()
	if err != nil {
		return res,err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected" : rowsAffected,
	}

	return res,nil
}
