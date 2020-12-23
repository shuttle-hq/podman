package generate

//go:generate go run ../generator/generator.go KubeOptions
// KubeOptions are optional options for generating kube YAML files
type KubeOptions struct {
	// Service - generate YAML for a Kubernetes _service_ object.
	Service *bool
}

//go:generate go run ../generator/generator.go SystemdOptions
// SystemdOptions are optional options for generating ssytemd files
type SystemdOptions struct {
	// Name - use container/pod name instead of its ID.
	UseName *bool
	// New - create a new container instead of starting a new one.
	New *bool
	// RestartPolicy - systemd restart policy.
	RestartPolicy *string
	// StopTimeout - time when stopping the container.
	StopTimeout *uint
	// ContainerPrefix - systemd unit name prefix for containers
	ContainerPrefix *string
	// PodPrefix - systemd unit name prefix for pods
	PodPrefix *string
	// Separator - systemd unit name separator between name/id and prefix
	Separator *string
}