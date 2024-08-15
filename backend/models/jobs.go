package models

import (
  "fmt"
  "time"
  "github.com/google/uuid"
  "github.com/pivot/database"
  _ "github.com/lib/pq"
)

type CreateJobRequest struct {
  JobName     string `json:"name"`
  CompanyId   string `json:"company_id"`
  ClientName  string `json:"client_name"`
}

type Job struct {
	ID          uuid.UUID `json:"id"`
  JobName     string    `json:"name"`
  CompanyId   uuid.UUID `json:"company_id"`
  ClientName  string    `json:"client_name"`
  CreatedAt   time.Time `json:"created_at"`
}

func CreateJob(job CreateJobRequest) (uuid.UUID, error) {
  companyUUID, err := uuid.Parse(job.CompanyId)

	if err != nil {
    fmt.Println("ERROR: ", err)
		return uuid.Nil, err
	}

  stmt, err := database.DB.Prepare(`
		INSERT INTO jobs (company_id, name, client_name)
		VALUES ($1, $2, $3)
		RETURNING id
	`)

  if err != nil {
    fmt.Println("ERROR: ", err)
    return uuid.Nil, err
  }
  defer stmt.Close()


  //var createdAt time.Time
  var jobId uuid.UUID 

  err = stmt.QueryRow(companyUUID, job.JobName, job.ClientName).Scan(&jobId)
	if err != nil {
    fmt.Println("ERROR: ", err)
		return uuid.Nil, err
	}
  fmt.Println("jobid: ",jobId)
  
  return jobId, nil
  //return Job {
  //  ID:          jobId,
  //  JobName:     job.JobName,
  //  CompanyId:   companyUUID,
  //  ClientName:  job.ClientName,
  //  CreatedAt:   createdAt,
  //}, nil
}

func GetJob(jobId string) (string, error) {
  return "", nil
}

func DeleteJob(jobId, companyId string) error {
  query := "DELETE FROM jobs WHERE id = $1 AND company_id = $2"

  result, err := database.DB.Exec(query, jobId, companyId)
  if err != nil {
    return fmt.Errorf("failed to execute delete query: %v", err)
  }

  rowsAffected, err := result.RowsAffected()
  if err != nil {
    return fmt.Errorf("failed to retrieve affected rows: %v", err)
  }

  if rowsAffected == 0 {
    return fmt.Errorf("no rows deleted, jobId: %s, companyId: %s", jobId, companyId)
  }

  fmt.Printf("Deleted %d rows from jobs\n", rowsAffected)
  return nil
}
//func DeleteJob(jobId, companyId string) (error) {
//  tx, err := database.DB.Begin()
//  if err != nil {
//    return fmt.Errorf("failed to begin transaction: %v", err)
//  }
//
//  query := "DELETE FROM job_metadata WHERE job_id = $1"
//  _, err = tx.Exec(query, jobId)
//  if err != nil {
//    tx.Rollback()
//    return fmt.Errorf("failed to delete job metadata from database: %v", err)
//  }
//
//  query = "DELETE FROM job_audit WHERE job_id = $1"
//  _, err = tx.Exec(query, jobId)
//  if err != nil {
//    tx.Rollback()
//    return fmt.Errorf("failed to delete job audit from database: %v", err)
//  }
//
//  query = "DELETE FROM jobs WHERE id = $1 AND company_id = $2"
//  _, err = tx.Exec(query, jobId)
//  if err != nil {
//    tx.Rollback()
//    return fmt.Errorf("failed to delete job metadata from database: %v", err)
//  }
//
//  err = tx.Commit()
//  if err != nil {
//    return fmt.Errorf("failed to commit transaction: %v", err)
//  }
//
//  //result, err := database.DB.Exec(query, jobId, companyId)
//  //if err != nil {
//  //  return fmt.Errorf("failed to delete job: %v", err)
//  //}
//
//  // Check if any rows were affected
//  //rowsAffected, err := result.RowsAffected()
//  //if err != nil {
//  //  return fmt.Errorf("failed to check rows affected: %v", err)
//  //}
//  //if rowsAffected == 0 {
//  //  return fmt.Errorf("job not found or does not belong to the company")
//  //}
//
//  return nil
//}

func GetAllJobs(companyId string) ([]Job, error) {
  query := "SELECT id, name, company_id, client_name, status, created_at FROM jobs WHERE company_id = $1"

  stmt, err := database.DB.Prepare(query)
  if err != nil {
    return nil, fmt.Errorf("failed to prepare statement: %v", err)
  }
  defer stmt.Close()

  rows, err := stmt.Query(companyId)
  if err != nil {
    return nil, fmt.Errorf("failed to execute query: %v", err)
  }

  defer rows.Close()

  var jobs []Job

  for rows.Next() {
    var job Job
    if err := rows.Scan(&job.ID, &job.JobName, &job.CompanyId, &job.ClientName, &job.CreatedAt); err != nil {
      return nil, fmt.Errorf("failed to scan row: %v", err)
    }
    jobs = append(jobs, job)
  }

  if err := rows.Err(); err != nil {
    return nil, fmt.Errorf("error occurred while interating over rows: %v", err)
  }

  return jobs, nil
}

func UpdateJobAudit(jobId, name, action, changes string) (error) {
  // return error if not successful
  validActions := map[string]bool{
    "CREATE": true,
    "UPDATE": true,
    "DELETE": true,
  }

  if !validActions[action] {
    return fmt.Errorf("invalid action: %s", action)
  }

  query := `
    INSERT INTO job_audit (job_id, modified_by, action, changes)
    VALUES ($1, $2, $3, $4)
  `

  _, err := database.DB.Exec(query, jobId, name, action, changes)
  if err != nil {
    return fmt.Errorf("error inserting into job_audit table: %w", err)
  }

  return nil
}

