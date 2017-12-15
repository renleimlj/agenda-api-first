package service

import(
    "net/http"
    "github.com/emicklei/go-restful"
    "github.com/satori/go.uuid"
)

type Userlog struct {
	Name, Password string
}

func UserRegister(container *restful.Container) {
    ws := new(restful.WebService)
    ws.
    Path("/user").
        Consumes(restful.MIME_XML, restful.MIME_JSON).
        Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

    ws.Route(ws.POST("/login").To(login))
    ws.Route(ws.DELETE("/logout/keys/{key}").To(logout))

    container.Add(ws)
}

func login(request *restful.Request, response *restful.Response) {
	usrlog := new(Userlog)
    err := request.ReadEntity(&usrlog)
    if err == nil {
        users := ListAllUsers()
        usr , ok := users[usrlog.Name]
        if !ok {
	        response.AddHeader("Content-Type", "text/plain")
	        response.WriteErrorString(http.StatusBadRequest, "Wrong name")
	    } else {
	    	if isLogin(usrlog.Name) {
	    		response.AddHeader("Content-Type", "text/plain")
	    		response.WriteErrorString(http.StatusUnauthorized, "Log out first")
	    	} else {
	        	if usrlog.Password != usr.Password {
	        		response.AddHeader("Content-Type", "text/plain")
	        		response.WriteErrorString(http.StatusBadRequest, "Wrong password")
	        	} else {
	        		key := uuid.NewV4().String()
	        		CreateKey(key, usr.Name)
	        		//response.WriteEntity(key)
	        		response.AddHeader("Content-Type", "text/plain")
	   				response.WriteHeaderAndEntity(http.StatusCreated, key)
	        	}
	        }
	    }
    } else {
        response.AddHeader("Content-Type", "text/plain")
        response.WriteErrorString(http.StatusInternalServerError, err.Error())
    }
}

func logout(request *restful.Request, response *restful.Response) {
    k := request.PathParameter("key")
    if FindUserbyKey(k) == "" {
    	response.AddHeader("Content-Type", "text/plain")
	    response.WriteErrorString(http.StatusUnauthorized, "Wrong key")
    } else {
	   	err := DeleteKey(k)
	   	if err != nil {
			response.AddHeader("Content-Type", "text/plain")
	       	response.WriteErrorString(http.StatusInternalServerError, err.Error())
	   	} else {
	   		response.AddHeader("Content-Type", "text/plain")
	   		response.WriteHeaderAndEntity(http.StatusNoContent, "byebye")
	   	}
    }
}











