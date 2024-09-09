package api

// About is some general information about the API
type About struct {
	App       string    `json:"app"`
	Auths     []string  `json:"auths"`
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	CreatedAt string    `json:"created_at"`
	Uptime    uint64    `json:"uptime_seconds"`
	Version   Version   `json:"version"`
	Resources Resources `json:"resources"`
}

// Version is some information about the binary
type Version struct {
	Number   string `json:"number"`
	Commit   string `json:"repository_commit"`
	Branch   string `json:"repository_branch"`
	Build    string `json:"build_date"`
	Arch     string `json:"arch"`
	Compiler string `json:"compiler"`
}

type Resources struct {
	IsThrottling bool    `json:"is_throttling"`      // Whether this core is currently throttling
	NCPU         float64 `json:"ncpu"`               // Number of CPU on this node
	CPU          float64 `json:"cpu_used"`           // Current CPU load, 0-100*ncpu
	CPULimit     float64 `json:"cpu_limit"`          // Defined CPU load limit, 0-100*ncpu
	CPUCore      float64 `json:"cpu_core"`           // Current CPU load of the core itself, 0-100*ncpu
	Mem          uint64  `json:"memory_used_bytes"`  // Currently used memory in bytes
	MemLimit     uint64  `json:"memory_limit_bytes"` // Defined memory limit in bytes
	MemTotal     uint64  `json:"memory_total_bytes"` // Total available memory in bytes
	MemCore      uint64  `json:"memory_core_bytes"`  // Current used memory of the core itself in bytes
}

// MinimalAbout is the minimal information about the API
type MinimalAbout struct {
	App     string         `json:"app"`
	Auths   []string       `json:"auths"`
	Version VersionMinimal `json:"version"`
}

type VersionMinimal struct {
	Number string `json:"number"`
}
