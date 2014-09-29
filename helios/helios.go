package helios

type Port struct {
	ExternalPort float64 `json:"externalPort"`
	InternalPort float64 `json:"internalPort"`
	Protocol     string  `json:"protocol"`
}

type Job struct {
	Command     []interface{}   `json:"command"`
	Env         struct{}        `json:"env"`
	Expires     interface{}     `json:"expires"`
	GracePeriod interface{}     `json:"gracePeriod"`
	ID          string          `json:"id"`
	Image       string          `json:"image"`
	Ports       map[string]Port `json:"ports"`

	// Registration parameters for service discovery, not currently implemented by this client.
	Registration       struct{}          `json:"registration"`
	RegistrationDomain string            `json:"registrationDomain"`
	Volumes            map[string]string `json:"volumes"`
}

type TaskStatus struct {
	ContainerId string            `json:"containerId"`
	Env         map[string]string `json:"env"`
	Goal        string            `json:"goal"`

	// Job that this task is an instance of.
	Job       Job             `json:"job"`
	Ports     map[string]Port `json:"ports"`
	State     string          `json:"state"`
	Throttled string          `json:"throttled"`
}

type HostStatus struct {
	AgentInfo struct {
		InputArguments []string `json:"inputArguments"`
		Name           string   `json:"name"`
		SpecName       string   `json:"specName"`
		SpecVendor     string   `json:"specVendor"`
		SpecVersion    string   `json:"specVersion"`
		StartTime      float64  `json:"startTime"`
		Uptime         float64  `json:"uptime"`
		Version        string   `json:"version"`
		VmName         string   `json:"vmName"`
		VmVendor       string   `json:"vmVendor"`
		VmVersion      string   `json:"vmVersion"`
	} `json:"agentInfo"`

	Environment map[string]string `json:"environment"`

	HostInfo struct {
		Architecture  string  `json:"architecture"`
		Cpus          float64 `json:"cpus"`
		DockerVersion struct {
			ApiVersion    string `json:"apiVersion"`
			Arch          string `json:"arch"`
			GitCommit     string `json:"gitCommit"`
			GoVersion     string `json:"goVersion"`
			KernelVersion string `json:"kernelVersion"`
			Os            string `json:"os"`
			Version       string `json:"version"`
		} `json:"dockerVersion"`
		Hostname         string  `json:"hostname"`
		LoadAvg          float64 `json:"loadAvg"`
		MemoryFreeBytes  float64 `json:"memoryFreeBytes"`
		MemoryTotalBytes float64 `json:"memoryTotalBytes"`
		OsName           string  `json:"osName"`
		OsVersion        string  `json:"osVersion"`
		SwapFreeBytes    float64 `json:"swapFreeBytes"`
		SwapTotalBytes   float64 `json:"swapTotalBytes"`
		Uname            string  `json:"uname"`
	} `json:"hostInfo"`

	// Status of the host ("UP", "DOWN", etc.).
	Status string `json:"status"`

	// Tasks is a map of job ID to task status. A task is a running instance of a job.
	Tasks map[string]TaskStatus `json:"statuses"`
}
