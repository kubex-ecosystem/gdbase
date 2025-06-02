package services

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/goccy/go-json"
	"github.com/pebbe/zmq4"
	gl "github.com/rafa-mori/gdbase/logger"
)

const (
	HeartbeatInterval = 2500 * time.Millisecond // Interval between heartbeats
)

type BrokerImpl struct {
	context     *zmq4.Context
	frontend    *zmq4.Socket // FRONTEND (ROUTER) with clients
	backend     *zmq4.Socket // BACKEND (DEALER) with workers
	services    map[string]*Service
	workers     map[string]*Worker
	waiting     []*Worker
	mu          sync.Mutex
	heartbeatAt time.Time
	brokerInfo  *BrokerInfoLock
	verbose     bool
}
type Service struct {
	name     string
	requests [][]string
	waiting  []*Worker
}
type Worker struct {
	identity string
	service  *Service
	expiry   time.Time
	broker   *BrokerImpl
}

func NewBrokerConn(port string) (*zmq4.Socket, error) {
	ctx, err := zmq4.NewContext()
	if err != nil {
		return nil, fmt.Errorf("error creating ZMQ context: %v", err)
	}

	frontend, err := ctx.NewSocket(zmq4.ROUTER)
	if err != nil {
		return nil, fmt.Errorf("error creating FRONTEND (ROUTER): %v", err)
	}
	frontendSetRouterMandatoryErr := frontend.SetRouterMandatory(1)
	if frontendSetRouterMandatoryErr != nil {
		return nil, frontendSetRouterMandatoryErr
	}
	frontendSetRouterHandoverErr := frontend.SetRouterHandover(true)
	if frontendSetRouterHandoverErr != nil {
		return nil, frontendSetRouterHandoverErr
	}

	if hostBindErr := frontend.Bind(`tcp://0.0.0.0:` + port); hostBindErr != nil {
		return nil, fmt.Errorf("error binding FRONTEND (ROUTER): %v", hostBindErr)
	}

	return frontend, nil
}
func NewBroker(verbose bool) (*BrokerImpl, error) {
	ctx, err := zmq4.NewContext()
	if err != nil {
		return nil, fmt.Errorf("error creating ZMQ context: %v", err)
	}

	frontend, err := ctx.NewSocket(zmq4.ROUTER)
	if err != nil {
		return nil, fmt.Errorf("error creating FRONTEND (ROUTER): %v", err)
	}
	frontendSetRouterMandatoryErr := frontend.SetRouterMandatory(1)
	if frontendSetRouterMandatoryErr != nil {
		return nil, frontendSetRouterMandatoryErr
	}
	frontendSetRouterHandoverErr := frontend.SetRouterHandover(true)
	if frontendSetRouterHandoverErr != nil {
		return nil, frontendSetRouterHandoverErr
	}

	if hostBindErr := frontend.Bind(`tcp://0.0.0.0:5555`); hostBindErr != nil {
		return nil, fmt.Errorf("error binding FRONTEND (ROUTER): %v", hostBindErr)
	}

	backend, err := ctx.NewSocket(zmq4.DEALER)
	if err != nil {
		return nil, fmt.Errorf("error creating BACKEND (DEALER): %v", err)
	}
	if bindErr := backend.Bind("inproc://backend"); bindErr != nil {
		return nil, fmt.Errorf("error binding BACKEND (DEALER): %v", bindErr)
	}

	broker := &BrokerImpl{
		brokerInfo:  NewBrokerInfo(RndomName(), "5555"),
		context:     ctx,
		frontend:    frontend,
		backend:     backend,
		services:    make(map[string]*Service),
		workers:     make(map[string]*Worker),
		waiting:     []*Worker{},
		heartbeatAt: time.Now().Add(HeartbeatInterval),
		verbose:     verbose,
	}

	if broker.brokerInfo == nil {
		gl.Log("error", "Error creating broker")
		return nil, fmt.Errorf("error creating broker: Empty broker info")
	}
	data, marshalErr := json.Marshal(broker.brokerInfo.GetBrokerInfo())
	if marshalErr != nil {
		gl.Log("error", "Error marshalling broker info")
		return nil, marshalErr
	}
	if writeErr := os.WriteFile(broker.brokerInfo.GetPath(), data, 0644); writeErr != nil {
		gl.Log("error", "Error writing broker file")
		return nil, writeErr
	}

	// Launch workers
	for i := 0; i < 5; i++ {
		go broker.workerTask()
	}

	// Start the proxy
	go broker.startProxy()

	// Start heartbeat management
	//go broker.handleHeartbeats()

	return broker, nil
}

func (b *BrokerImpl) startProxy() {
	gl.Log("info", "Starting proxy between FRONTEND and BACKEND...")
	err := zmq4.Proxy(b.frontend, b.backend, nil)
	if err != nil {
		gl.Log("error", "Error in proxy between FRONTEND and BACKEND")
	}
}
func (b *BrokerImpl) workerTask() {
	worker, err := b.context.NewSocket(zmq4.DEALER)
	if err != nil {
		gl.Log("error", "Error creating socket for worker")
		return
	}

	if connErr := worker.Connect("inproc://backend"); connErr != nil {
		gl.Log("error", "Error connecting worker to BACKEND")
		return
	}

	for {
		// msg, _ := worker.RecvMessage(0)
		// if len(msg) < 2 {
		// 	gl.Log("debug", "Malformed message received in WORKER")
		// 	continue
		// }

		// //id, msg := splitMessage(msg)

		// //payload := msg[len(msg)-1]
		// //pld := []byte(payload)

		// // deserializedModel, deserializedModelErr := m.NewModelRegistryFromSerialized(pld).(Model)
		// // if deserializedModelErr != nil {
		// // 	gl.Log("error", "Error deserializing payload in WORKER")
		// // 	continue
		// // }

		// gl.Log("debug", "Payload deserialized in WORKER")

		// // tp, tpErr := deserializedModel.GetType()
		// // if tpErr != nil {
		// // 	gl.Log("error", "Error getting payload type in WORKER")
		// // 	continue
		// // }

		// gl.Log("debug", "Payload type in WORKER")

		// if tp.Name() == "PingImpl" {
		// 	response := fmt.Sprintf(`{"type":"ping","data":{"ping":"%v"}}`, "pong")
		// 	if _, workerSendMessageErr := worker.SendMessage(id, response); workerSendMessageErr != nil {
		// 		gl.Log("error", "Error sending response to BACKEND in WORKER")
		// 	} else {
		// 		gl.Log("debug", "Response sent to BACKEND in WORKER")
		// 	}
		// } else {
		// 	gl.Log("debug", "Unknown command in WORKER")
		// }
	}
}
func (b *BrokerImpl) handleHeartbeats() {
	ticker := time.NewTicker(HeartbeatInterval)
	//defer ticker.Stop()
	//defer b.mu.Unlock()

	for range ticker.C {
		b.mu.Lock()
		now := time.Now()
		for id, worker := range b.workers {
			if now.After(worker.expiry) {
				gl.Log("warn", fmt.Sprintf("Expired worker: %s", id))
				delete(b.workers, id)
			}
		}
		b.mu.Unlock()
	}

	b.mu.Unlock()
}
func (b *BrokerImpl) Stop() {
	_ = b.frontend.Close()
	_ = b.backend.Close()
	_ = b.context.Term()
	gl.Log("info", "Broker stopped")
}
