package socketgo

type socket_type uint8

const (
	UDP        socket_type = 0
	TCP        socket_type = 1
	WEB_SOCKET socket_type = 2
)

type IO interface {
	Write(b []byte) error
}

type Worker interface {
	Run()
}

type Conn struct {
}

type Pr struct {
	workers []Worker
}

func (pr *Pr) RunAll() {
	pr.workers[0].Run()
	l := len(pr.workers)
	for i := 0; i < l; i++ {
		go pr.workers[i].Run()
	}
	for {

	}
}

func getConn(id int64) IO {
	for _, c := range connectors {
		if c.id == id {
			return c.conn
		}
	}
	return nil
}

func sendToClient(id int64, b []byte) bool {
	io := getConn(id)
	if io != nil {
		_ = io.Write(b)
		return true
	}
	return false
}

func (pr *Pr) AddServer(server *Server) {
	switch server.sType {
	case TCP:
		worker := &TCPWorker{server}
		pr.workers = append(pr.workers, worker)
		break
	case WEB_SOCKET:
		worker := &WSWorker{server}
		pr.workers = append(pr.workers, worker)
		break
	default:

	}
}

var (
	counter    int64 = 0
	connectors []*Connector
)

type Connector struct {
	id   int64
	conn IO
}

type Server struct {
	id        int64
	sType     socket_type
	addr      string
	OnMessage func(conn *Connector, message []byte)
	OnError   func(err error)
	OnClose   func(conn *Connector)
	OnConnect func(conn *Connector)
	OnStart   func()
	OnOpen    func()
}
