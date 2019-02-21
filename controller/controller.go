package controller

import (
	"encoding/json"
	"fmt"
	"gorm_http_server_demo/model"
	"net/http"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

// Welcome ...
func Welcome(w http.ResponseWriter, req *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	fmt.Fprintf(w, "Welcome to the home page!")
}

// CreateUser ...
func CreateUser(w http.ResponseWriter, req *http.Request) {
	returnValue := make(map[string]interface{})
	if req.Method != "POST" {
		returnValue["errmesg"] = "Not a POST method"
		returnMessage, _ := json.Marshal(returnValue)
		fmt.Fprintln(w, string(returnMessage))
		return
	}
	if err := req.ParseForm(); err != nil {
		returnValue["errmesg"] = "ParseForm() error"
		returnMessage, _ := json.Marshal(returnValue)
		fmt.Fprintf(w, string(returnMessage))
		return
	}

	user := model.Gousers{}
	uuid, _ := uuid.NewV4()
	userid := uuid.String()
	user.Userid = userid
	user.Name = req.FormValue("name")
	user.Gender = req.FormValue("gender")
	user.Age, _ = strconv.ParseInt(req.FormValue("age"), 10, 64)
	returnValue = user.Create()

	jsonEncode, _ := json.Marshal(returnValue)
	fmt.Fprintf(w, string(jsonEncode))

}

// ReadUser ...
func ReadUser(w http.ResponseWriter, req *http.Request) {
	returnValue := make(map[string]interface{})
	if req.Method != "POST" {
		returnValue["errmesg"] = "Not a POST method"
		returnMessage, _ := json.Marshal(returnValue)
		fmt.Fprintln(w, string(returnMessage))
		return
	}
	if err := req.ParseForm(); err != nil {
		returnValue["errmesg"] = "ParseForm() error"
		returnMessage, _ := json.Marshal(returnValue)
		fmt.Fprintf(w, string(returnMessage))
		return
	}

	user := model.Gousers{}
	user.Userid = req.FormValue("userid")

	returnValue = user.Read()
	if returnValue["userid"] == "" {
		returnValue["errmesg"] = "Member not found, please check your userid"
	} else {
		returnValue["errmesg"] = nil
	}
	jsonEncode, _ := json.Marshal(returnValue)
	fmt.Fprintf(w, string(jsonEncode))
}

// UpdateUser ...
func UpdateUser(w http.ResponseWriter, req *http.Request) {
	returnValue := make(map[string]interface{})
	if req.Method != "POST" {
		returnValue["errmesg"] = "Not a POST method"
		returnMessage, _ := json.Marshal(returnValue)
		fmt.Fprintln(w, string(returnMessage))
		return
	}
	if err := req.ParseForm(); err != nil {
		returnValue["errmesg"] = "ParseForm() error"
		returnMessage, _ := json.Marshal(returnValue)
		fmt.Fprintf(w, string(returnMessage))
		return
	}

	user := model.Gousers{}
	user.Userid = req.FormValue("userid")

	var value string
	keys := []string{"userid", "name", "gender", "age"}
	for _, k := range keys {
		value = req.FormValue(k)
		if value != "" {
			switch k {
			case "age":
				user.Age, _ = strconv.ParseInt(value, 10, 64)
			case "name":
				user.Name = value
			case "gender":
				user.Gender = value
			}
		}
	}
	returnValue = user.Update()
	jsonEncode, _ := json.Marshal(returnValue)
	fmt.Fprintf(w, string(jsonEncode))
}

// DeleteUser ...
func DeleteUser(w http.ResponseWriter, req *http.Request) {
	returnValue := make(map[string]interface{})
	if req.Method != "POST" {
		returnValue["errmesg"] = "Not a POST method"
		returnMessage, _ := json.Marshal(returnValue)
		fmt.Fprintln(w, string(returnMessage))
		return
	}
	if err := req.ParseForm(); err != nil {
		returnValue["errmesg"] = "ParseForm() error"
		returnMessage, _ := json.Marshal(returnValue)
		fmt.Fprintf(w, string(returnMessage))
		return
	}

	user := model.Gousers{}
	user.Userid = req.FormValue("userid")
	returnValue = user.Delete()
	jsonEncode, _ := json.Marshal(returnValue)
	fmt.Fprintf(w, string(jsonEncode))
}
