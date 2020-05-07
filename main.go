package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main()  {
	router := mux.NewRouter()

	router.HandleFunc("/api/upload", uploadFile).Methods("POST")

	err := http.ListenAndServe(":6000", router)

	if err != nil {
		fmt.Println("error")
	}
}

func uploadFile(response http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10*1024*1024)

	file, handle, err := request.FormFile("file")

	if err != nil {
		fmt.Println("error in file")
	}

	defer file.Close()

	fmt.Println("file info")
	fmt.Println("name :", handle.Filename)
	fmt.Println("headr :", handle.Header)
	fmt.Println("size: ", handle.Size)

	//request -> body= templateContect

}