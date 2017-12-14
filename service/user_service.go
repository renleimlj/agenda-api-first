package service

import(
    "net/http"
    "github.com/emicklei/go-restful"
)

func UsersRegister(container *restful.Container) {
    ws := new(restful.WebService)
    ws.
    Path("/users").
        Consumes(restful.MIME_XML, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

    ws.Route(ws.POST("").To(cu))
    ws.Route(ws.GET("").To(au))

    container.Add(ws)
}

func cu(request *restful.Request, response *restful.Response) {
    usr := new(User)
    err := request.ReadEntity(&usr)
    if err == nil {
        err1 := CreateUser(*usr)
        if err1 != nil {
            response.WriteErrorString(http.StatusNotFound, err1.Error())
        } else {
            response.WriteEntity(usr)
        }
    } else {
        response.AddHeader("Content-Type", "text/plain")
        response.WriteErrorString(http.StatusInternalServerError, err.Error())
    }
}

func au(request *restful.Request, response *restful.Response) {
    users := ListAllUsers()
    response.WriteEntity(users)
}

