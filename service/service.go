package service

import(
    "net/http"
    "github.com/emicklei/go-restful"
)

func MeetingRegister(container *restful.Container) {
    ws := new(restful.WebService)
    ws.
    Path("/meetings").
        Consumes(restful.MIME_XML, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

    ws.Route(ws.POST("").To(cm))
    ws.Route(ws.GET("/{m-title}").To(qm))

    container.Add(ws)
}

func cm(request *restful.Request, response *restful.Response) {
    met := new(Meeting)
    err := request.ReadEntity(&met)
    if err == nil {
        err1 := CreateMeeting(*met)
        if err1 != nil {
            response.WriteErrorString(http.StatusNotFound, err1.Error())
        } else {
            response.WriteEntity(met)
        }
    } else {
        response.AddHeader("Content-Type", "text/plain")
        response.WriteErrorString(http.StatusInternalServerError, err.Error())
    }
}

func qm(request *restful.Request, response *restful.Response) {
    t := request.PathParameter("m-title")
    meetings := ListAllMeetings()
    met , ok := meetings[t]
    if !ok {
        response.AddHeader("Content-Type", "text/plain")
        response.WriteErrorString(http.StatusNotFound, "User could not be found.")
    } else {
        response.WriteEntity(met)
    }
}




