// Package deploy handles running deployment scripts and
// persisting metadata about each deployment
package deploy

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"
)

// Deployer is interface for managing deployments
type Deployer interface {
	// Deploy invokes the deployment script and stores a new deployment record
	Deploy() (err error)

	// List retrieves all deployments
	List() (deployments []Deployment, err error)

	// Get retrieves a deployment by index
	Get(idx int) (deployment *Deployment, err error)
}

// Status represents the variou deployment states
type Status uint8

const (
	// StatusSuccess indicates a succcessful deployment
	StatusSuccess Status = iota

	// StatusInProgress indicates a deployment is underway
	StatusInProgress

	// StatusFail indicates a deployment failed
	StatusFail
)

func (s Status) String() string {
	switch s {
	case StatusSuccess:
		return "success"
	case StatusInProgress:
		return "inprogress"
	case StatusFail:
		return "fail"
	default:
		return "unknown"
	}
}

// Deployment encapsulates the log file and the status of the
// script execution
type Deployment struct {
	File   *os.File
	Status Status
}

type deployer struct {
	logPath    string
	scriptPath string

	mux         sync.Mutex
	deployments []Deployment
}

func (d *deployer) Deploy() (err error) {
	d.mux.Lock()
	defer d.mux.Unlock()

	idx := len(d.deployments)
	logPath := fmt.Sprintf("%s/%d.log", d.logPath, idx)
	logFd, err := os.Create(logPath)
	if err != nil {
		return err
	}

	dep := Deployment{
		Status: StatusInProgress,
		File:   logFd,
	}
	d.deployments = append(d.deployments, dep)

	go d.run(idx)

	return nil
}

func (d *deployer) run(depIdx int) {
	dep := d.deployments[depIdx]

	var errs = make([]error, 0)
	cmd := exec.Command("/bin/bash", d.scriptPath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		errs = append(errs, err)
	}

	go func() {
		io.Copy(dep.File, stdout)
	}()

	err = cmd.Run()
	if err != nil {
		errs = append(errs, err)
	}

	d.mux.Lock()
	if len(errs) > 0 {
		dep.Status = StatusFail
		fmt.Println(errs)
		// prob should append to log file instead of printing
	} else {
		dep.Status = StatusSuccess
	}
	d.deployments[depIdx] = dep
	d.mux.Unlock()
}

func (d *deployer) List() ([]Deployment, error) {
	return d.deployments, nil
}

func (d *deployer) Get(idx int) (*Deployment, error) {
	if idx < 0 || idx >= len(d.deployments) {
		return nil, fmt.Errorf("deployment #%d not found", idx)
	}
	return &d.deployments[idx], nil
}

// New returns a simple Deployer using in-memory storage.
// So can't scale or persist between process restarts. This
// is just a simple solution. Can reimplement a datastore later.
func New(scriptPath, logPath string) (Deployer, error) {
	if err := os.RemoveAll(logPath); err != nil {
		return nil, err
	}

	if err := os.MkdirAll(logPath, 0755); err != nil {
		return nil, err
	}

	return &deployer{
		logPath:    logPath,
		scriptPath: scriptPath,

		mux:         sync.Mutex{},
		deployments: make([]Deployment, 0),
	}, nil
}
