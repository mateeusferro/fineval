package auditly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mateeusferro/fineval/internal/config"
	"github.com/mateeusferro/fineval/internal/domain"
)

func Send(logData domain.AuditlyInput) bool {
	auditlyHost := config.EnvVariable("AUDITLY_HOST")
	data, err := json.Marshal(logData)

	if err != nil {
		fmt.Println("Error in JSON: ", err)
		return false
	}

	resp, err := http.Post(
		auditlyHost+"/log",
		"application/json",
		bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("Error requesting auditly: ", err)
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return false
	}

	log.Print(body)
	return resp.StatusCode == 201
}
