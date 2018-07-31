package livy

import (
	"fmt"

	"github.com/pkg/errors"
)

// Returns all the active Batch sessions.
func (c *Client) GetBatches(from, size int) ([]Batch, error) {
	data := &GetBatchesRequest{
		From: from,
		Size: size,
	}

	req, err := c.NewRequest("GET", "/batches", data)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	var sessionsBatches GetBatchesResponse
	if err := c.Do(req, &sessionsBatches); err != nil {
		return nil, err
	}

	return sessionsBatches.Sessions, nil
}

// Start new Batch and return the created Batch object.
func (c *Client) PostBatch(batchRequest *PostBatchRequest) (*Batch, error) {
	req, err := c.NewRequest("POST", "/batches", batchRequest)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	var batch Batch
	if err := c.Do(req, &batch); err != nil {
		return nil, err
	}

	return &batch, nil
}

// Returns the Batch session information.
func (c *Client) GetBatch(batchID int) (*Batch, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("/batches/%d", batchID), nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	var batch Batch
	if err := c.Do(req, &batch); err != nil {
		return nil, err
	}

	return &batch, nil
}

// Returns the state of batch session.
func (c *Client) GetBatchState(batchID int) (string, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("/batches/%d/state", batchID), nil)
	if err != nil {
		return "", errors.Wrap(err, "creating request")
	}

	var batchState GetBatchStateResponse
	if err := c.Do(req, &batchState); err != nil {
		return "", err
	}

	return batchState.State, nil
}

// Kills the Batch job.
func (c *Client) KillBatchJob(batchID int) error {
	req, err := c.NewRequest("DELETE", fmt.Sprintf("/batches/%d", batchID), nil)
	if err != nil {
		return errors.Wrap(err, "creating request")
	}

	return c.Do(req, nil)
}

// Gets the log lines from this batch.
func (c *Client) GetBatchLogs(batchID, from, size int) ([]string, error) {
	data := &GetBatchLogsRequest{
		From: from,
		Size: size,
	}

	req, err := c.NewRequest("GET", fmt.Sprintf("/batches/%d/logs", batchID), data)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	var batchLogs GetBatchLogsResponse
	if err := c.Do(req, &batchLogs); err != nil {
		return nil, err
	}

	return batchLogs.Log, nil
}
