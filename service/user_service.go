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
        Produces(restful.MIME_JSON, restful.MIME_XML)
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
            response.AddHeader("Content-Type", "text/plain")
            response.WriteErrorString(http.StatusBadRequest, err1.Error())
        } else {
            response.AddHeader("Content-Type", "application/json")
            response.WriteHeaderAndEntity(http.StatusCreated, usr)
        }
    } else {
        response.AddHeader("Content-Type", "text/plain")
        response.WriteErrorString(http.StatusInternalServerError, err.Error())
    }
}

func au(request *restful.Request, response *restful.Response) {
    users := ListAllUsers()
    response.WriteHeaderAndEntity(http.StatusOK, users)
}

