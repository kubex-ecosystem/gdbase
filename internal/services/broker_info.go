package services

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	gl "github.com/kubex-ecosystem/logz/logger"
)

type BrokerInfo struct {
	Name string `json:"name"`
	Port string `json:"port"`
	PID  int    `json:"pid"`
	Time string `json:"time"`
	path string
}
type BrokerInfoLock struct {
	Name  string
	Port  string
	PID   int
	Time  string
	path  string
	flock sync.Mutex
}

func NewBrokerInfo(name, port string) *BrokerInfoLock {
	path, pathErr := GetBrokersPath()
	if pathErr != nil {
		gl.Log("error", "Error getting brokers path")
		return nil
	}

	if name == "" {
		name = RndomName()
	}

	path = filepath.Clean(filepath.Join(path, fmt.Sprintf("%s.json", name)))

	return &BrokerInfoLock{
		Name: name,
		Port: port,
		PID:  os.Getpid(),
		Time: time.Now().Format(time.RFC3339),
		path: path,
	}
}

func (bi *BrokerInfoLock) GetBrokerInfo() BrokerInfo {
	return BrokerInfo{
		Name: bi.Name,
		Port: bi.Port,
		PID:  bi.PID,
		Time: bi.Time,
		path: bi.path,
	}
}
func (bi *BrokerInfoLock) GetPath() string { return bi.path }
func (bi *BrokerInfoLock) GetPort() string { return bi.Port }
func (bi *BrokerInfoLock) GetName() string { return bi.Name }
func (bi *BrokerInfoLock) GetPID() int     { return bi.PID }
func (bi *BrokerInfoLock) GetTime() string { return bi.Time }
func (bi *BrokerInfoLock) Lock()           { bi.flock.Lock() }
func (bi *BrokerInfoLock) Unlock()         { bi.flock.Unlock() }
func (bi *BrokerInfoLock) String() string {
	return fmt.Sprintf("BrokerInfo{Name: %s, Port: %s, PID: %d, Time: %s}", bi.Name, bi.Port, bi.PID, bi.Time)
}
func (bi *BrokerInfoLock) trap() {
	bi.Lock()
	defer func() {
		bi.Unlock()
		if bi.path != "" {
			if rmErr := os.Remove(bi.path); rmErr != nil {
				gl.Log("error", "Error removing broker file")
			}
		}
	}()
}
