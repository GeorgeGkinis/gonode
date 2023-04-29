package node

import "github.com/georgegkinis/pubsub"

type NodeInterface interface {
	Name() string
	Type() string
	Init() error
	ParseConfig(cfg interface{}) error
	Inputs() pubsub.Subscribers
	Outputs() pubsub.Publishers
	InitInputs() error
	InitOutputs() error
}

type ports struct {
	in  pubsub.Subscribers
	out pubsub.Publishers
}

type Node struct {
	name     string
	nodeType string
	ports
	cfg interface{}
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) Type() string {
	return n.nodeType
}

func (n *Node) Inputs() pubsub.Subscribers {
	return n.in
}

func (n *Node) Outputs() pubsub.Publishers {
	return n.out
}
