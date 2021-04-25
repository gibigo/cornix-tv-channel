package config

import (
	"testing"
)

func TestGetBeforeConfigIsLoaded(t *testing.T) {
	defer func() { recover() }()
	Get()
	t.Fatal("Should've panicked because the configuration hasn't been loaded yet")
}

/* func TestLoadFileThatDoesNotExistAndWithoutEnvVars(t *testing.T) {
	defer func() { recover() }()
	_ = Load("file-that-does-not-exist.yaml")
	t.Error("Should've panicked, because the file specified doesn't exist and env is not set")
}

func TestLoadDefaultConfigurationFile(t *testing.T) {
	defer func() { recover() }()
	_ = Load(DefaultConfigurationFilePath)
	t.Error("Should've panicked, because there's no configuration files at the default path nor the default fallback path nor is the env set")

} */
