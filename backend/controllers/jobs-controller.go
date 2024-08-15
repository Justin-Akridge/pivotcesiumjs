package controllers

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/pivot/middleware"
  "github.com/pivot/models"
  "github.com/pivot/utils"
)

/***********************
* route /api/jobs
***********************/

func HandleJobs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
      case http.MethodGet:
        handleGetJobs(w, r)
        
      case http.MethodPost:
        handlePostJob(w, r)
        
      default:
        fmt.Fprintln(w, "method not allowed %s", r.Method)
    }
	}
}

func handleGetJobs(w http.ResponseWriter, r *http.Request) {
  claims, ok := r.Context().Value(middleware.ClaimsContextKey).(*utils.Claims)
  if !ok || claims == nil {
    log.Printf("Unable to retrieve claims from context")
    http.Error(w, "Unable to get claims from context", http.StatusInternalServerError)
    return
  }

  jobs, err := models.GetAllJobs(claims.CompanyId)
  if err != nil {
    log.Printf("Error retrieving jobs: %v", err)
    http.Error(w, "Unable to retrieve jobs", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(w).Encode(jobs); err != nil {
    log.Printf("Error encoding jobs to JSON: %v", err)
    http.Error(w, "Unable to encode jobs", http.StatusInternalServerError)
    return
  }
}

func handlePostJob(w http.ResponseWriter, r *http.Request) {
  claims, ok := r.Context().Value(middleware.ClaimsContextKey).(*utils.Claims)
  if !ok || claims == nil {
    log.Printf("Unable to retrieve claims from context")
    http.Error(w, "Unable to get claims from context", http.StatusInternalServerError)
    return
  }

  err := r.ParseMultipartForm(10 << 20)
  if err != nil {
    log.Printf("Failed to parse form: %v", err)
    http.Error(w, "Failed to parse form", http.StatusBadRequest)
    return
  }

  jobName := r.FormValue("jobName")
  clientName := r.FormValue("clientName")

  newJobRequest := models.CreateJobRequest {
    JobName: jobName,
    CompanyId: claims.CompanyId,
    ClientName: clientName,
  }

  jobId, err := models.CreateJob(newJobRequest)
  if err != nil {
    fmt.Println("error creating job")
    http.Error(w, "error creating new job in database", http.StatusInternalServerError)
    return
  }

  fmt.Println("JOBID: ", jobId)
  redirectURL := fmt.Sprintf("/%s", jobId.String())
  http.Redirect(w, r, redirectURL, http.StatusSeeOther)
  // convert uuid to string
  //jobIDStr := job.ID.String()
  //auditErr := models.UpdateJobAudit(jobIDStr, claims.Name, "CREATE", "Created job")
  //if auditErr != nil {
  //  log.Printf("Error updating job audit: %v", auditErr)
  //}

  //w.Header().Set("Content-Type", "application/json")
  //w.WriteHeader(http.StatusCreated) // 201
  //err = json.NewEncoder(w).Encode(job)
  //if err != nil {
  //  fmt.Println("error encoding job")
  //  http.Error(w, "error encoding new job, failure to return to client", http.StatusInternalServerError)
  //  return
  //}
}


/***********************
* route /api/jobs/${id}
***********************/

func HandleJob() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    jobId := vars["id"]

    if jobId == "" {
      http.Error(w, "no job id", http.StatusBadRequest)
      return
    }
    switch r.Method {
      case http.MethodGet:
        handleGetJob(w, r, jobId)

      case http.MethodDelete:
        handleDeleteJob(w, r, jobId)
        
      case http.MethodPut:
        handlePutJob(w, r, jobId)
        
      case http.MethodPatch:
        handlePatchJob(w, r, jobId)

      default:
        fmt.Fprintln(w, "method not allowed %s", r.Method)
    }
	}
}

func handleDeleteJob(w http.ResponseWriter, r *http.Request, jobId string) {
  claims, ok := r.Context().Value(middleware.ClaimsContextKey).(*utils.Claims)
  if !ok || claims == nil {
    log.Printf("Unable to retrieve claims from context")
    http.Error(w, "Unable to get claims from context", http.StatusInternalServerError)
    return
  }

  // delete job from database
  err := models.DeleteJob(jobId, claims.CompanyId)
  if err != nil {
    log.Printf("Error deleting job: %v", err)
    http.Error(w, "Failed to delete job", http.StatusInternalServerError)
    return
  }

  // delete any files in the s3 bucket
  //err = models.DeleteFilesFromS3Bucket(claims.CompanyId, jobId)
  //if err != nil {
  //  log.Printf("Error deleting files from s3 bucket: %v", err)
  //  http.Error(w, "Error deleting files from s3 bucket", http.StatusInternalServerError)
  //  return
  //}
  w.WriteHeader(http.StatusOK)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(map[string]string{"message": "Job deleted successfully"})
}


func handleGetJob(w http.ResponseWriter, r *http.Request, jobId string) {
  job, err := models.GetJob(jobId)
  if err != nil {
    http.Error(w, "Job not found", http.StatusNotFound)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(job)
}

func handlePutJob(w http.ResponseWriter, r *http.Request, jobId string) {}

func handlePatchJob(w http.ResponseWriter, r *http.Request, jobId string) {}

