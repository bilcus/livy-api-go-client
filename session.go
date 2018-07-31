package livy

import (
	"fmt"

	"github.com/pkg/errors"
)

// Returns all the active interactive sessions.
func (c *Client) GetSessions(from, size int) ([]Session, error) {
	data := &GetSessionsRequest{
		From: from,
		Size: size,
	}

	req, err := c.NewRequest("GET", "/sessions", data)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	var sessionsResponse GetSessionsResponse
	if err := c.Do(req, &sessionsResponse); err != nil {
		return nil, err
	}

	return sessionsResponse.Sessions, nil
}

// Creates a new interactive Scala, Python, or R shell in the cluster.
// Returns created Session.
//
// Starting with version 0.5.0-incubating the kind field is not required.
// To be compatible with previous versions users can still specify this with spark, pyspark or sparkr,
// implying that the submitted code snippet is the corresponding kind.
func (c *Client) StartSession(startSessionRequest StartSessionRequest) (*Session, error) {
	req, err := c.NewRequest("POST", "/sessions", startSessionRequest)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	var session Session
	if err := c.Do(req, &session); err != nil {
		return nil, err
	}

	return &session, nil
}

// Returns the session information.
func (c *Client) GetSession(sessionID int) (*Session, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("/sessions/%d", sessionID), nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	var session Session
	if err := c.Do(req, &session); err != nil {
		return nil, err
	}

	return &session, nil
}

// Returns the state of session.
func (c *Client) GetSessionState(sessionID int) (SessionState, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("/sessions/%d/state", sessionID), nil)
	if err != nil {
		return "", errors.Wrap(err, "creating request")
	}

	var sessionStateResponse GetSessionStateResponse
	if err := c.Do(req, &sessionStateResponse); err != nil {
		return "", err
	}

	return sessionStateResponse.State, nil
}

// Kills the Session job.
func (c *Client) KillSession(sessionID int) error {
	req, err := c.NewRequest("DELETE", fmt.Sprintf("/sessions/%d", sessionID), nil)
	if err != nil {
		return errors.Wrap(err, "creating request")
	}

	return c.Do(req, nil)
}

// Gets the log lines from this session.
func (c *Client) GetSessionLogs(sessionID, from, size int) ([]string, error) {
	data := &GetSessionLogsRequest{
		From: from,
		Size: size,
	}

	req, err := c.NewRequest("GET", fmt.Sprintf("/sessions/%d/log", sessionID), data)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	var sessionLogs GetSessionLogsResponse
	if err := c.Do(req, &sessionLogs); err != nil {
		return nil, err
	}

	return sessionLogs.Log, nil
}
