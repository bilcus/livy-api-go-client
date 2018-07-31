package livy

import (
	"fmt"

	"github.com/pkg/errors"
)

// Returns all the statements in a session.
func (c *Client) GetSessionStatements(sessionID int) ([]Statement, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("/sessions/%d/statements", sessionID), nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	var sessionStatements GetSessionStatements
	if err := c.Do(req, &sessionStatements); err != nil {
		return nil, err
	}

	return sessionStatements.Statements, nil
}

// Runs a statement in a session and returns a statement object.
//
// If session kind is not specified or the submitted code is not the kind specified in session creation,
// this field should be filled with correct kind.
// Otherwise Livy will use kind specified in session creation as the default code kind.
func (c *Client) RunStatement(sessionID int, code string, kind CodeKind) (*Statement, error) {
	data := &PostStatementRequest{
		Code: code,
		Kind: kind,
	}

	req, err := c.NewRequest("POST", fmt.Sprintf("/sessions/%d/statements", sessionID), data)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	var statement Statement
	if err := c.Do(req, &statement); err != nil {
		return nil, err
	}

	return &statement, nil
}

// Returns a specified statement in a session.
func (c *Client) GetStatement(sessionID int, statementID int) (*Statement, error) {
	req, err := c.NewRequest("GET", fmt.Sprintf("/sessions/%d/statements/%d", sessionID, statementID), nil)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	var statement Statement
	if err := c.Do(req, &statement); err != nil {
		return nil, err
	}

	return &statement, nil
}

// Cancel the specified statement in this session.
func (c *Client) CancelStatement(sessionID, statementID int) error {
	req, err := c.NewRequest("POST", fmt.Sprintf("/sessions/%d/statements/%d/cancel", sessionID, statementID), nil)
	if err != nil {
		return errors.Wrap(err, "creating request")
	}

	return c.Do(req, nil)
}

// Runs a statement in a session.
func (c *Client) RunStatementGetCompletion(sessionID int, code string, kind CodeKind, cursor string) ([]string, error) {
	data := &PostStatementCompletionRequest{
		Code:   code,
		Kind:   kind,
		Cursor: cursor,
	}

	req, err := c.NewRequest("POST", fmt.Sprintf("/sessions/%d/completion", sessionID), data)
	if err != nil {
		return nil, errors.Wrap(err, "creating request")
	}

	var completion PostStatementCompletionResponse
	if err := c.Do(req, &completion); err != nil {
		return nil, err
	}

	return completion.Candidates, nil
}
