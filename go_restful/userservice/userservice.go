package userservice

import (
	"errors"
	"net/http"

	"github.com/emicklei/go-restful"
)

//User comment
type User struct {
	Id, Name string
}

var u = User{Id: "101", Name: "ramesh"}

//New comment
func New() *restful.WebService {
	service := new(restful.WebService)
	service.Path("/users").Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	service.Route(service.GET("/{user-id}").To(FindUser))
	service.Route(service.POST("").To(UpdateUser))
	service.Route(service.PUT("/{user-id}").To(CreateUser))
	service.Route(service.DELETE("/{user-id}").To(RemoveUser))

	return service
}

//FindUser comment
func FindUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	// here you would fetch user from some persistence system
	if u.Id == id {
		usr := &User{Id: u.Id, Name: u.Name}
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, errors.New("User not found"))
	}
	//usr := &User{Id: id, Name: "John Doe"}
	//response.WriteEntity(usr)
}

//UpdateUser comment
func UpdateUser(request *restful.Request, response *restful.Response) {
	usr := new(User)
	err := request.ReadEntity(&usr)
	// here you would update the user with some persistence system
	u = *usr
	if err != nil {
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

//CreateUser comment
func CreateUser(request *restful.Request, response *restful.Response) {
	usr := User{Id: request.PathParameter("user-id")}
	err := request.ReadEntity(&usr)
	// here you would create the user with some persistence system
	u = User{Id: usr.Id, Name: usr.Name}
	if err == nil {
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

//RemoveUser comment
func RemoveUser(request *restful.Request, response *restful.Response) {
	// here you would delete the user from some persistence system
	u = User{}
}
