package models

type Mahasiswa struct {
	Id       int    `form:"id" json:"id"`
	Nama     string `form:"nama" json:"nama"`
	Usia     int    `form:"usia" json:"usia"`
	Gender   int    `form:"gender" json:"gender"`
	TglRegis string `form:"tgl_registrasi" json:"tgl_registrasi"`
	Jurusan  string `form:"jurusan" json:"jurusan"`
	Hobi     string `form:"hobi" json:"hobi"`
}
