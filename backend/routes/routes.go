package routes

import (
	"net/http"
	"github.com/pivot/controllers"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "../frontend/dist/index.html")
}

func SetupRoutes() {
  http.HandleFunc("/", serveIndex)
  //http.HandleFunc("/", controllers.HomePageHandler())
  http.HandleFunc("/login", controllers.HandleLogin())
  http.HandleFunc("/checkAuth", controllers.VerifyTokenHandler())
  http.HandleFunc("/jobs", controllers.HandleJobs())
  http.HandleFunc("/jobs/{id}", controllers.HandleJob())

  //serve static files
  staticFileServer := http.FileServer(http.Dir("../frontend/dist/assets"))
  http.Handle("/assets/", http.StripPrefix("/assets/", staticFileServer))

	//r.HandleFunc("/", controllers.HomePageHandler()).Methods(http.MethodGet)
	//r.HandleFunc("/{id}", controllers.NewPageHandler()).Methods(http.MethodGet)
	////r.HandleFunc("/login", controllers.LoginHandler()).Methods(http.MethodGet, http.MethodPost)
	//r.HandleFunc("/checkAuth", controllers.VerifyTokenHandler())

  //r.HandleFunc("/login", loginHandler()).Methods("GET")

	//r.HandleFunc("/v1/api/register", controllers.RegisterHandler()).Methods(http.MethodPost)

	////jobs
	//r.HandleFunc("/jobs", controllers.HandleJobs()).Methods(http.MethodGet, http.MethodPost)
	//r.HandleFunc("/jobs/{id}", controllers.HandleJob()).Methods(http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPatch)

	////las files
	//r.HandleFunc("/v1/api/{jobId}/files", controllers.HandleFiles()).Methods(http.MethodGet, http.MethodPost)
	//r.HandleFunc("/v1/api/{jobId}/files/{id}", controllers.HandleFile()).Methods(http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPatch)

	// photos
	// TODO
	//r.HandleFunc("/v1/api/photos/", controllers.HandlePhotos()).Methods(http.MethodGet, http.MethodPost)
	//r.HandleFunc("/v1/api/photos/{id}", controllers.HandlePhoto()).Methods(http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPatch)

	//poles
	// TODO
	//r.HandleFunc("/v1/api/poles/", controllers.HandlePoles()).Methods(http.MethodGet, http.MethodPost)
	//r.HandleFunc("/v1/api/poles/{id}", controllers.HandlePole()).Methods(http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPatch)

	//midspans
//	r.HandleFunc("/api/job/{jobId}/midspans", controllers.HandleJobMidspans()).Methods(http.MethodGet, http.MethodPost)
//	r.HandleFunc("/api/job/{jobId}/midspans/{id}", controllers.HandleJobMidspan()).Methods(http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPatch)
//
	//vegetation encroachments
	// TODO
	//r.HandleFunc("/v1/api/vegetations/", controllers.HandleVegetations()).Methods(http.MethodGet, http.MethodPost)
	//r.HandleFunc("/v1/api/vegetations/{id}", controllers.HandleVegetation()).Methods(http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPatch)

}
