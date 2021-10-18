package restore_logs

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client/client"
)

const restoreLogsServiceEndpoint string = "%s/archive/restore"

const (
	RestoreStatusInProgress    = "IN_PROGRESS"
	RestoreStatusActive        = "ACTIVE"
	RestoreStatusLimitExceeded = "LIMIT_EXCEEDED"
	RestoreStatusAborted       = "ABORTED"
	RestoreStatusFailed        = "FAILED"
	RestoreStatusDeleted       = "DELETED"
	RestoreStatusExpired       = "EXPIRED"

	listRestoreOperations  = "ListRestoreOperations"
	getRestoreOperation    = "GetRestoreOperation"
	deleteRestoreOperation = "DeleteRestoreOperation"
)

type RestoreClient struct {
	*client.Client
}

type InitiateRestore struct {
	AccountName string `json:"accountName"` // Name of the restored account
	StartTime   int64  `json:"startTime"`
	EndTime     int64  `json:"endTime"`
}

type RestoreOperation struct {
	Id               int32   `json:"id"`        // ID of the restore operation in Logz.io
	AccountId        int32   `json:"accountId"` // ID of the restored account in Logz.io
	AccountName      string  `json:"accountName"`
	RestoredVolumeGb float64 `json:"restoredVolumeGb,omitempty"` // nullable
	Status           string  `json:"status,omitempty"`
	StartTime        float64 `json:"startTime"`
	EndTime          float64 `json:"endTime"`
	CreatedAt        float64 `json:"createdAt"`
	StartedAt        float64 `json:"startedAt,omitempty"`  // nullable
	FinishedAt       float64 `json:"finishedAt,omitempty"` // nullable
	ExpiresAt        float64 `json:"expiresAt,omitempty"`  // nullable
}

// New Creates a new entry point into the restore logs functions, accepts the user's logz.io API token and base url
func New(apiToken string, baseUrl string) (*RestoreClient, error) {
	if len(apiToken) == 0 {
		return nil, fmt.Errorf("API token not defined")
	}
	if len(baseUrl) == 0 {
		return nil, fmt.Errorf("Base URL not defined")
	}

	c := &RestoreClient{
		Client: client.New(apiToken, baseUrl),
	}
	return c, nil
}
