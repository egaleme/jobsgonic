package job

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetJobs(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.GET("/api/jobs/", GetJobs)
	req, err := http.NewRequest("GET", "/api/jobs/", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
	}

}

func TestCreateJob(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.POST("/jobs", CreateJob)

	jobJson := `{"Id": 99, "Description": "banking operations", "CategoryId": 555, "Company":"uba","Position":"manager", "Location":"benin", "ExpiresAt":"12/12/2007", "CreatedAt":"12/12/2007","UpdatedAt":"12/12/2007", "PosterEmail": "askaleme@yahoo.com"}`

	req, err := http.NewRequest("POST", "/jobs", strings.NewReader(jobJson))
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", w.Code)
	}
}

func TestGetJobsClient(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.GET("/api/jobs/", GetJobs)

	server := httptest.NewServer(r)
	defer server.Close()
	urlStr := fmt.Sprintf("%s/api/jobs/", server.URL)
	req, err := http.NewRequest("GET", urlStr, nil)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d", res.StatusCode)
	}

}

func TestCreateJobsClient(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.POST("/jobs", CreateJob)

	server := httptest.NewServer(r)
	defer server.Close()
	urlStr := fmt.Sprintf("%s/jobs", server.URL)
	jobJson := `{"Id": 99, "Description": "banking operations", "CategoryId": 555, "Company":"uba","Position":"manager", "Location":"benin", "ExpiresAt":"12/12/2007", "CreatedAt":"12/12/2007","UpdatedAt":"12/12/2007", "PosterEmail": "askaleme@yahoo.com"}`

	body := strings.NewReader(jobJson)
	req, err := http.NewRequest("POST", urlStr, body)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", res.StatusCode)

	}
}
