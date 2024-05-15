// A generated module for Grafana functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
)

type Grafana struct{}

// Returns a container that echoes whatever string argument is provided
func (m *Grafana) ContainerEcho(stringArg string) *Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

func (g *Grafana) Hello() string {
  return "HELLO ALEX FROM CI"
}
// Returns lines that match a pattern in the files of the provided Directory
func (m *Grafana) GrepDir(ctx context.Context, directoryArg *Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

func (g *Grafana) Carre(input int) int {
  return input * input
}

func (m *Grafana) GetUser2(ctx context.Context) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithExec([]string{"apk", "add", "curl"}).
		WithExec([]string{"apk", "add", "jq"}).
		WithExec([]string{"sh", "-c", "curl https://randomuser.me/api/ | jq .results[0].name"}).
		Stdout(ctx)
}

func (g *Grafana) GetUser() string {
  resp, err := http.Get("https://randomuser.me/api/")
  if err != nil {
    log.Fatalln(err)
  }
  body, err := io.ReadAll(resp.Body)
  if err != nil {
    log.Fatalln(err)
  }
  log.Print(string(body))
  return string(body)
}

func (g *Grafana) DeployDashboards(dir string) string {
  files, err := os.ReadDir(dir)
  if err != nil {
    log.Fatal(err)
  }
  for _, f := range files {
		log.Println(f.Name())
    return f.Name()
	}
  return "fini"
}
