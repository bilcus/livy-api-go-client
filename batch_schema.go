package livy

type Batch struct {
	ID      int               `json:"id"`
	AppID   string            `json:"appId"`
	AppInfo map[string]string `json:"appInfo"`
	Log     []string          `json:"log"`
	State   string            `json:"state"`
}

type GetBatchesRequest struct {
	From int `json:"from,omitempty"`
	Size int `json:"size,omitempty"`
}

type GetBatchesResponse struct {
	From     int     `json:"from"`
	Total    int     `json:"total"`
	Sessions []Batch `json:"sessions"`
}

type PostBatchRequest struct {
	File           string            `json:"file,omitempty"`
	ProxyUser      string            `json:"proxyUser,omitempty"`
	ClassName      string            `json:"className,omitempty"`
	Args           []string          `json:"args,omitempty"`
	Jars           []string          `json:"jars,omitempty"`
	PyFiles        []string          `json:"pyFiles,omitempty"`
	Files          []string          `json:"files,omitempty"`
	DriverMemory   string            `json:"driverMemory,omitempty"`
	DriverCores    int               `json:"driverCores,omitempty"`
	ExecutorMemory string            `json:"executorMemory,omitempty"`
	ExecutorCores  int               `json:"executorCores,omitempty"`
	NumExecutors   int               `json:"numExecutors,omitempty"`
	Archives       []string          `json:"archives,omitempty"`
	Queue          string            `json:"queue,omitempty"`
	Name           string            `json:"name,omitempty"`
	Conf           map[string]string `json:"conf,omitempty"`
}

type GetBatchStateResponse struct {
	ID    int    `json:"id"`
	State string `json:"state"`
}

type GetBatchLogsRequest struct {
	From int `json:"from,omitempty"`
	Size int `json:"size,omitempty"`
}

type GetBatchLogsResponse struct {
	ID   int      `json:"id"`
	From int      `json:"from"`
	Size int      `json:"size"`
	Log  []string `json:"log"`
}
