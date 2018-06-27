package userservice

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
)

// User ...
type User struct {
	ID, Name string
}

var users []User

func init() {
	users = append(users, User{ID: "1", Name: "John"})
	users = append(users, User{ID: "2", Name: "Bill"})
	users = append(users, User{ID: "3", Name: "Sarah"})
}

// New ...
func New() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	service.Route(service.GET("").To(ListUser))
	service.Route(service.GET("/{user-id}").To(FindUser))
	service.Route(service.POST("/{user-id}").To(UpdateUser))
	service.Route(service.PUT("/{user-id}").To(CreateUser))
	service.Route(service.DELETE("/{user-id}").To(RemoveUser))

	return service
}

// ListUser ...
func ListUser(request *restful.Request, response *restful.Response) {
	response.WriteEntity(users)
}

// FindUser ...
func FindUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	for _, item := range users {
		if item.ID == id {
			response.WriteEntity(item)
			return
		}
	}
	response.WriteEntity(&User{})
}

// UpdateUser ...
func UpdateUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	usr := new(User)
	err := request.ReadEntity(&usr)

	if err == nil {
		usr.ID = id
		for index, item := range users {
			if item.ID == usr.ID {
				users = append(users[:index], users[index+1:]...)
				users = append(users, *usr)
				break
			}
		}
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// CreateUser ...
func CreateUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	usr := new(User)
	err := request.ReadEntity(&usr)
	if err == nil {
		usr.ID = id
		users = append(users, *usr)
		response.WriteEntity(usr)
	} else {
		response.WriteError(http.StatusInternalServerError, err)
	}
}

// RemoveUser ...
func RemoveUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	for index, item := range users {
		if item.ID == id {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	response.WriteEntity(users)
}
