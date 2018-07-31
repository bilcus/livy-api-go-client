package livy

const (
	StateNotStarted   SessionState = "not_started"
	StateStarting     SessionState = "starting"
	StateIdle         SessionState = "idle"
	StateBusy         SessionState = "busy"
	StateShuttingDown SessionState = "shutting_down"
	StateError        SessionState = "error"
	StateDead         SessionState = "dead"
	StateSuccess      SessionState = "success"
)

type SessionState string

// A session represents an interactive shell.
type Session struct {
	ID        int               `json:"id"`
	AppID     string            `json:"appId"`
	Owner     string            `json:"owner"`
	ProxyUser string            `json:"proxyUser"`
	Kind      CodeKind          `json:"kind"`
	Log       []string          `json:"log"`
	State     SessionState      `json:"state"`
	AppInfo   map[string]string `json:"appInfo"`
}

type StartSessionRequest struct {
	Kind                     CodeKind          `json:"kind,omitempty"`
	ProxyUser                string            `json:"proxyUser,omitempty"`
	Jars                     []string          `json:"jars,omitempty"`
	PyFiles                  []string          `json:"pyFiles,omitempty"`
	Files                    []string          `json:"files,omitempty"`
	DriverMemory             string            `json:"driverMemory,omitempty"`
	DriverCores              int               `json:"driverCores,omitempty"`
	ExecutorMemory           string            `json:"executorMemory,omitempty"`
	ExecutorCores            int               `json:"executorCores,omitempty"`
	NumExecutors             int               `json:"numExecutors,omitempty"`
	Archives                 []string          `json:"archives,omitempty"`
	Queue                    string            `json:"queue,omitempty"`
	Name                     string            `json:"name,omitempty"`
	Conf                     map[string]string `json:"conf,omitempty"`
	HeartbeatTimeoutInSecond int               `json:"heartbeatTimeoutInSecond,omitempty"`
}

type GetSessionsRequest struct {
	From int `json:"from,omitempty"`
	Size int `json:"size,omitempty"`
}

type GetSessionsResponse struct {
	From     int       `json:"from"`
	Total    int       `json:"total"`
	Sessions []Session `json:"sessions"`
}

type GetSessionStateResponse struct {
	ID    int          `json:"id,omitempty"`
	State SessionState `json:"state"`
}

type GetSessionLogsRequest struct {
	From int `json:"from,omitempty"`
	Size int `json:"size,omitempty"`
}

type GetSessionLogsResponse struct {
	ID   int      `json:"id"`
	From int      `json:"from"`
	Size int      `json:"size"`
	Log  []string `json:"log"`
}
