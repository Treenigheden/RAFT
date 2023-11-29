package raft

import (
	"math/rand"
	"time"
)

type Server struct {
	port         int
	isLeader     bool
	leaderPort   int
	gotHeartbeat bool
	term         int
}

func (s *Server) leaderCandidate(prot int, term int) (vote bool) {
	if s.term > term {
		//call this on all conected node
		s.leaderCandidate(s.port, s.term)
		return false
	}
	s.gotHeartbeat = true
	return true
}

func (s *Server) checkForElection() {
	for !s.isLeader {
		//this needs to be something other than sleep, because
		//it needs a reset if input turns true
		time.Sleep(time.Duration(rand.Intn(300-150) + 150))
		if !s.gotHeartbeat {
			//call this on all conected node
			s.leaderCandidate(s.port, s.term)
		}
	}
}

func findAvailablePort(startPort int) int {
	return startPort
}

func main() {
	var server Server
	startPort := 8000
	server.port = findAvailablePort(startPort)

	go server.checkForElection()

}
