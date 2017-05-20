package cmd

import (
	"net"
	"testing"

	"github.com/bpicode/fritzctl/config"
	"github.com/bpicode/fritzctl/mock"
	"github.com/stretchr/testify/assert"
)

// TestListInetStats  tests the command.
func TestListInetStats(t *testing.T) {
	config.ConfigDir = "../testdata"
	config.ConfigFilename = "config_localhost_https_test.json"
	srv := mock.New().UnstartedServer()
	l, err := net.Listen("tcp", ":61666")
	assert.NoError(t, err)
	defer l.Close()
	srv.Listener = l
	srv.Start()
	defer srv.Close()
	command, err := ListInetstats()
	assert.NoError(t, err)
	exitCode := command.Run([]string{})
	assert.Equal(t, 0, exitCode)
}

// TestListInetStatsHasHelp ensures that the tested command provides a help text.
func TestListInetStatsHasHelp(t *testing.T) {
	command, err := ListInetstats()
	assert.NoError(t, err)
	help := command.Help()
	assert.NotEmpty(t, help)
}

// TestListInetStatsHasSynopsis ensures the tested command provides short a synopsis text.
func TestListInetStatsHasSynopsis(t *testing.T) {
	command, err := ListInetstats()
	assert.NoError(t, err)
	synopsis := command.Synopsis()
	assert.NotEmpty(t, synopsis)
}