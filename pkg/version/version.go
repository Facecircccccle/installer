package version

type ComponentVersion struct {
	DockerVersion  []string
	EtcdVersion    []string
	CoreDNSVersion []string
	PauseVersion   []string
}

var version = map[string]ComponentVersion{
	"1.20.1": {
		DockerVersion:  []string{"docker-ce-19.03.15"},
		EtcdVersion:    []string{"3.4.13-0"},
		CoreDNSVersion: []string{"1.7.0"},
		PauseVersion:   []string{"3.2"},
	},
	"1.20.2": {
		DockerVersion:  []string{"docker-ce-19.03.15"},
		EtcdVersion:    []string{"3.4.13-0"},
		CoreDNSVersion: []string{"1.7.0"},
		PauseVersion:   []string{"3.2"},
	},
	"1.20.3": {
		DockerVersion:  []string{"docker-ce-19.03.15"},
		EtcdVersion:    []string{"3.4.13-0"},
		CoreDNSVersion: []string{"1.7.0"},
		PauseVersion:   []string{"3.2"},
	},
	"1.20.4": {
		DockerVersion:  []string{"docker-ce-19.03.15"},
		EtcdVersion:    []string{"3.4.13-0"},
		CoreDNSVersion: []string{"1.7.0"},
		PauseVersion:   []string{"3.2"},
	},
	"1.20.5": {
		DockerVersion:  []string{"docker-ce-19.03.15"},
		EtcdVersion:    []string{"3.4.13-0"},
		CoreDNSVersion: []string{"1.7.0"},
		PauseVersion:   []string{"3.2"},
	},
	"1.20.6": {
		DockerVersion:  []string{"docker-ce-19.03.15"},
		EtcdVersion:    []string{"3.4.13-0"},
		CoreDNSVersion: []string{"1.7.0"},
		PauseVersion:   []string{"3.2"},
	},
}

func GetComponentVersion() map[string]ComponentVersion {
	return version
}
