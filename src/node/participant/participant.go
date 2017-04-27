package participant

import (
  "net/rpc"
  "net"
  "log"
)

type Participant struct {
  Objects map[string]Object
  Transactions map[int32]Transaction
  Address string
  Id int
}

var self Participant

func Start(hostname string, id int) error {
  log.Println("Starting participant")
  self = New(hostname, id)
  go self.setupRPC()
  return nil
}

func (p Participant) setupRPC()  {
  log.Println("Setting up participant RPCs")
  rpc.Register(&self)
  l, e := net.Listen("tcp", ":3000")
  if e != nil {
    log.Println("Error in setup RPC:", e)
  }
  go rpc.Accept(l)
}

func New(addr string, id int) Participant {
  objs := make(map[string]Object, 0)
  trans := make(map[int32]Transaction, 0)
  return Participant{objs, trans, addr, id}
}
