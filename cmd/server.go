package main

import (
	"encoding/json"
	"github.com/henriqueholtz/imersion-fullcycle-golang/model"
	"net/http"
)

var courseList = model.NewCourses()

func main() {
	print("Hello! Server is running at 8888 port.")
	http.HandleFunc("/courses", CourseListHanlder)
	http.HandleFunc("/courses", CourseListHanlder)
	http.ListenAndServe(":8888", nil)
}

func CourseListHanlder(w http.ResponseWriter, r *http.Request) {
	courseJson, _ := json.Marshal(courseList)
	w.Write([]byte(courseJson))
}
