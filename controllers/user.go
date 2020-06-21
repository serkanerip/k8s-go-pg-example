package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/serkanerip/go-pg/models"
	"github.com/serkanerip/go-pg/responses"
)

// ListUsers return user list as json
func (s *Server) ListUsers(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	user := models.User{}

	users, err := user.FindAllUsers(s.DB)
	if err != nil {
		responses.ERROR(resp, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(resp, http.StatusOK, users)
}

// CreateUser handle create user api endpoint
func (s *Server) CreateUser(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		responses.ERROR(resp, http.StatusUnprocessableEntity, err)
		return
	}

	fmt.Println(string(body))
	user := models.User{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(resp, http.StatusUnprocessableEntity, err)
		return
	}

	savedUser, err := user.SaveUser(s.DB)

	responses.JSON(resp, http.StatusCreated, struct {
		MSG  string      `json:"msg"`
		USER models.User `json:"user"`
	}{
		MSG: "created", USER: *savedUser,
	})
}
