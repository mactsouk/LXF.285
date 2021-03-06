package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var PORT = ":8008"
var UPLOADPATH = "./data/"

func uploadHandler(c *gin.Context) {
	fileID := c.PostForm("fileID")

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Error in form: %s", err))
		return
	}

	filename := UPLOADPATH + filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s\n", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded with fileID = %s.", file.Filename, fileID))
}

func defaultHandler(c *gin.Context) {
	c.String(http.StatusNotFound, "Connecting from "+c.ClientIP())
}

func main() {
	fmt.Println("Starting HTTP server!")
	router := gin.Default()
	router.Static("/", "./public")
	router.POST("/upload", uploadHandler)

	router.NoRoute(defaultHandler)
	router.Run(PORT)
}
