package livy

const (
	CodeSpark   CodeKind = "spark"
	CodePySpark CodeKind = "pyspark"
	CodeSparkR  CodeKind = "sparkr"
	CodeSQL     CodeKind = "sql"

	StatementWaiting    StatementState = "waiting"
	StatementRunning    StatementState = "running"
	StatementAvailable  StatementState = "available"
	StatementError      StatementState = "error"
	StatementCancelling StatementState = "cancelling"
	StatementCancelled  StatementState = "cancelled"

	StatementOutputOK    StatementOutputStatus = "ok"
	StatementOutputError StatementOutputStatus = "error"
)

type CodeKind string

type StatementState string

type StatementOutputStatus string

// A statement represents the result of an execution statement.
type Statement struct {
	ID     int             `json:"id"`
	Code   string          `json:"code"`
	State  StatementState  `json:"state"`
	Output StatementOutput `json:"output"`
}

type StatementOutput struct {
	Status         StatementOutputStatus `json:"status"`
	ExecutionCount int                   `json:"execution_count"`
	Data           map[string]string     `json:"data"`
	Ename          string                `json:"ename"`
	Evalue         string                `json:"evalue"`
	Traceback      []string              `json:"traceback"`
}

type GetSessionStatements struct {
	Statements []Statement `json:"statements"`
}

type PostStatementRequest struct {
	Code string   `json:"code,omitempty"`
	Kind CodeKind `json:"kind,omitempty"`
}

type PostStatementCompletionRequest struct {
	Code   string   `json:"code,omitempty"`
	Kind   CodeKind `json:"kind,omitempty"`
	Cursor string   `json:"cursor,omitempty"`
}

type PostStatementCompletionResponse struct {
	Candidates []string `json:"candidates"`
}
