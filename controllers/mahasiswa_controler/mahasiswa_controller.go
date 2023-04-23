package mahasiswa_controller

import (
	"golang_jobhun_api/helper"
	"golang_jobhun_api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context){
	var mahasiswa models.Mahasiswa
	var arrMahasiswa []models.Mahasiswa

	db := models.Connect()
	defer db.Close()

	sql1 := "SELECT mahasiswa.id , mahasiswa.nama , mahasiswa.usia , mahasiswa.gender , mahasiswa.tgl_registrasi, jurusan.jurusan FROM mahasiswa LEFT JOIN jurusan ON mahasiswa.id=jurusan.id ORDER BY mahasiswa.id;"
	sql2 := "SELECT hobi.hobi FROM mahasiswa_hobi JOIN hobi ON mahasiswa_hobi.id_mahasiswa = hobi.id WHERE mahasiswa_hobi.id_mahasiswa = ?"
	
	rows ,err := db.Query(sql1)
	helper.PanicError(err)
	for rows.Next() {
		err := rows.Scan(&mahasiswa.Id ,&mahasiswa.Nama ,&mahasiswa.Usia,&mahasiswa.Gender,&mahasiswa.TglRegis , &mahasiswa.Jurusan)
		if  err != nil {
			panic(err)
		} else {
			mhs_id := &mahasiswa.Id
		
			rows1 ,err := db.Query(sql2 , mhs_id)
			helper.PanicError(err)
			if rows1.Next() {
				err := rows1.Scan(&mahasiswa.Hobi)
				helper.PanicError(err)
			}
			arrMahasiswa = append(arrMahasiswa, mahasiswa)
		}
	}
	c.AbortWithStatusJSON(http.StatusOK , gin.H{
		"mahasiswa" : arrMahasiswa,
		"name" : "Success Get All Data",
		"message" : "Success Get All Data",
		"status" : 200,
	})
}

func FindId(c *gin.Context){
	var mahasiswa models.Mahasiswa
	id := c.Param("id")

	db := models.Connect()
	defer db.Close()

	sql1 := "SELECT mahasiswa.id , mahasiswa.nama , mahasiswa.usia , mahasiswa.gender , mahasiswa.tgl_registrasi, jurusan.jurusan FROM mahasiswa LEFT JOIN jurusan ON mahasiswa.id=jurusan.id WHERE mahasiswa.id=?"
	sql2 := "SELECT hobi.hobi FROM mahasiswa_hobi JOIN hobi ON mahasiswa_hobi.id_hobi = hobi.id WHERE mahasiswa_hobi.id_mahasiswa = ?"

	rows ,err := db.Query(sql1,id)
	helper.PanicError(err)

	if rows.Next() {
		err := rows.Scan(&mahasiswa.Id ,&mahasiswa.Nama ,&mahasiswa.Usia,&mahasiswa.Gender,&mahasiswa.TglRegis , &mahasiswa.Jurusan)
		helper.PanicError(err)
	}
	rows ,err = db.Query(sql2,id)
	helper.PanicError(err)

	if rows.Next() {
		err := rows.Scan(&mahasiswa.Hobi)
		if  err != nil {
			panic(err)
		}
	}
	c.AbortWithStatusJSON(http.StatusOK , gin.H{
		"mahasiswa" : mahasiswa,
		"name" : "Success Get Data By Id",
		"message" : "Success Get Data By Id",
		"status" : 200,
	})
}

func Create(c *gin.Context){
	
	nama := c.PostForm("nama")
	usia := c.PostForm("usia")
	gender := c.PostForm("gender")

	tgl_registrasi := time.Now()
	jurusan := c.PostForm("jurusan")
	hobi := c.PostForm("hobi")
	db := models.Connect()
    defer db.Close()
	
	sql := "INSERT INTO mahasiswa (nama , usia , gender , tgl_registrasi) VALUES (?,?,?,?)"
	mhs, err := db.Exec( sql,nama ,usia, gender , tgl_registrasi)
	helper.PanicError(err)

	id, err := mhs.LastInsertId()
	helper.PanicError(err)
	sql = "INSERT INTO mahasiswa_hobi (id_mahasiswa , id_hobi) VALUES (?,?)"
	_, err = db.Exec( sql , id , hobi)


	sql = "INSERT INTO jurusan (jurusan) VALUES (?)"
	_, err = db.Exec( sql,jurusan)
	if err != nil {
		panic(err)
	} else {
		c.AbortWithStatusJSON(http.StatusOK , gin.H{
			"name" : "Success Create Data",
			"message" : "Success Create Data",
			"status" : 200,
			"id" : id,
		})
	}
}

func Update(c *gin.Context){
	id := c.Param("id")
	nama := c.PostForm("nama")
	usia := c.PostForm("usia")
	gender := c.PostForm("gender")
	tgl_registrasi := time.Now()
	hobi := c.PostForm("hobi")
	jurusan := c.PostForm("jurusan")
	db := models.Connect()
    defer db.Close()
	
	sql := "UPDATE mahasiswa SET nama = ? , usia = ? , gender = ? , tgl_registrasi = ? WHERE id = ?"
	_,err := db.Exec( sql,nama ,usia, gender , tgl_registrasi , id)
	helper.PanicError(err)

	sql = "UPDATE mahasiswa_hobi SET id_hobi = ? WHERE id_mahasiswa = ?"
	_, err = db.Exec( sql ,hobi ,id)
	helper.PanicError(err)
	sql = "UPDATE jurusan SET jurusan = ? WHERE id = ?"
	_, err = db.Exec( sql,jurusan , id)
	if err != nil {
		panic(err)
	} else {
		c.AbortWithStatusJSON(http.StatusOK , gin.H{
			"name" : "Success Update Data",
			"message" : "Success Update Data",
			"status" : 200,
		})
	}
}

func Delete(c *gin.Context){
	sql := "DELETE FROM mahasiswa,jurusan , mahasiswa_hobi USING mahasiswa , jurusan , mahasiswa_hobi WHERE mahasiswa.id=? AND jurusan.id=? AND mahasiswa_hobi.id_mahasiswa = ?"
	id := c.Param("id")

	db := models.Connect()
	defer db.Close()

	_ , err := db.Query(sql , id , id , id)
	if err != nil {
		panic(err)
	} else {
		c.AbortWithStatusJSON(http.StatusOK , gin.H{
			"name" : "Success Delete Data",
			"message" : "Success Delete Data",
			"status" : 200,
		})
	}
}

func AnyPath(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK , gin.H{
		"name" : "not found",
		"message" : "page not found",
		"status" : 404,
	})
}