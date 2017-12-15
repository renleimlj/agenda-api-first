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

    ws.Route(ws.POST("/keys/{key}").To(cm))
    ws.Route(ws.GET("/{m-title}/keys/{key}").To(qm))
    ws.Route(ws.DELETE("/{m-title}/keys/{key}").To(dm))
    container.Add(ws)
}

func cm(request *restful.Request, response *restful.Response) {
    met := new(Meeting)
    err := request.ReadEntity(&met)
    if err == nil {
        k := request.PathParameter("key")
        if FindUserbyKey(k) == "" {
            response.AddHeader("Content-Type", "text/plain")
            response.WriteErrorString(http.StatusUnauthorized, "Wrong key")
        } else {
            if met.Spon != FindUserbyKey(k) {
                response.AddHeader("Content-Type", "text/plain")
                response.WriteErrorString(http.StatusBadRequest, "You should be the sponser")
            } else {
                err1 := CreateMeeting(*met)
                if err1 != nil {
                    response.AddHeader("Content-Type", "text/plain")
                    response.WriteErrorString(http.StatusBadRequest, "Name conflict")
                } else {
                    response.AddHeader("Content-Type", "application/json")
                    response.WriteHeaderAndEntity(http.StatusCreated, met)
                }
            }
        }
    } else {
        response.AddHeader("Content-Type", "text/plain")
        response.WriteErrorString(http.StatusInternalServerError, err.Error())
    }
}

func qm(request *restful.Request, response *restful.Response) {
    t := request.PathParameter("m-title")
    k := request.PathParameter("key")
    if FindUserbyKey(k) == "" {
        response.AddHeader("Content-Type", "text/plain")
        response.WriteErrorString(http.StatusUnauthorized, "Wrong key")
    } else {
        meetings := ListAllMeetings()
        met , ok := meetings[t]
        if !ok {
            response.AddHeader("Content-Type", "text/plain")
            response.WriteErrorString(http.StatusBadRequest, "No result")
        } else {
            response.AddHeader("Content-Type", "application/json")
            response.WriteHeaderAndEntity(http.StatusOK, met)
        }
    }
}

func dm(request *restful.Request, response *restful.Response) {
    t := request.PathParameter("m-title")
    k := request.PathParameter("key")
    if FindUserbyKey(k) == "" {
        response.AddHeader("Content-Type", "text/plain")
        response.WriteErrorString(http.StatusUnauthorized, "Wrong key")
    } else {
        meetings := ListAllMeetings()
        met , ok := meetings[t]
        if ok {
            if met.Spon != FindUserbyKey(k) {
                response.AddHeader("Content-Type", "text/plain")
                response.WriteErrorString(http.StatusUnauthorized, "You are not sponser of this meeting")
                return
            }
        }
        err := DeleteMeeting(t)
        if err != nil {
            response.AddHeader("Content-Type", "text/plain")
            response.WriteErrorString(http.StatusInternalServerError, "")
        } else {
            response.AddHeader("Content-Type", "text/plain")
            response.WriteHeaderAndEntity(http.StatusNoContent, "succeed")
        }
    }
}