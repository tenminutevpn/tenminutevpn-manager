package resource

type Metadata struct {
	Name string `yaml:"name"`
}

type Resource struct {
	Kind     string   `yaml:"kind"`
	Metadata Metadata `yaml:"metadata"`
}
