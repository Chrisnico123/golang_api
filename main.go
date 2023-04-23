package main

import (
	mahasiswa_controller "golang_jobhun_api/controllers/mahasiswa_controler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/mahasiswa" , mahasiswa_controller.FindAll)
	r.GET("/api/mahasiswa/:id" , mahasiswa_controller.FindId)
	r.PUT("/api/mahasiswa/:id" , mahasiswa_controller.Update)
	r.POST("/api/mahasiswa" , mahasiswa_controller.Create)
	r.DELETE("/api/mahasiswa/:id" , mahasiswa_controller.Delete)
	// r.GET("/api/mahasiswa/:id/*any" , mahasiswa_controller.AnyPath)

	r.Run()
}