package service

import(
    //"log"
    //"net/http"
    "github.com/emicklei/go-restful"
)

type MeetingResource struct {
    Meetings map[string]Meeting
}

func (m MeetingResource) MeetingRegister(container *restful.Container) {
    ws := new(restful.WebService)
    ws.
    Path("/meetings").
        Consumes(restful.MIME_XML, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

    //ws.Route(ws.GET("/{user-id}").To(u.findUser))
    //ws.Route(ws.POST("").To(u.updateUser))
    //ws.Route(ws.PUT("/{user-id}").To(u.createUser))
    //ws.Route(ws.DELETE("/{user-id}").To(u.removeUser))

    ws.Route(ws.POST("").To(m.cm))
    ws.Route(ws.GET("").To(m.am))

    container.Add(ws)
}

func (m *MeetingResource) cm(request *restful.Request, response *restful.Response) {

}

func (m *MeetingResource) am(request *restful.Request, response *restful.Response) {

}




