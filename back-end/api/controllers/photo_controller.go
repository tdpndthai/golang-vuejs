package controllers

import (
	_ "encoding/json"
	"fmt"
	_ "fmt"
	"io/ioutil"
	_ "io/ioutil"
	_ "mime/multipart"
	"net/http"
	_ "net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tdpndthai/golang-vuejs/api/models"
	_ "github.com/tdpndthai/golang-vuejs/api/models"
	"github.com/tdpndthai/golang-vuejs/api/responses"
	_ "github.com/tdpndthai/golang-vuejs/api/responses"
)

func (server *Server) UploadFile(w http.ResponseWriter, r *http.Request) {
	//limit 10MB
	r.ParseMultipartForm(10 * 1024 * 1024)
	file, handler, err := r.FormFile("myfile")

	if err != nil {
		fmt.Println(err)
		return
	}
	//trì hoãn việc thực hiện file.Close() cho đến khi hàm xung quanh trả giá trị về
	defer file.Close()

	fmt.Println("File Info")
	fmt.Println("File name: ", handler.Filename)
	fmt.Println("File size: ", handler.Size)
	fmt.Println("File type: ", handler.Header.Get("Content-Type"))

	//upload file,tham số truyền vào gồm đường dẫn + tên ảnh sẽ lưu
	tempFile, err2 := ioutil.TempFile("uploads", "upload-*.jpg")
	//fmt.Println("file name upload",tempFile.Name())
	filePath := tempFile.Name()
	if err2 != nil {
		fmt.Println(err2)
	}
	defer tempFile.Close()

	fileBytes, err3 := ioutil.ReadAll(file)
	if err3 != nil {
		fmt.Println(err3)
	}

	tempFile.Write(fileBytes)
	photo := models.Photo{
		Src: filePath,
	}

	photoCreated, err := photo.SavePhoto(server.DB)

	//fmt.Println("Xong")
	//add multipart form
	responses.JSON(w, http.StatusCreated, photoCreated)
}

func (server *Server) GetFileById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	photo := models.Photo{}
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	result, err1 := photo.FindPhotoById(server.DB, uint32(uid))
	if err1 != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, result)
}
