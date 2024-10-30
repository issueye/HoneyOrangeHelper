package plugin

import (
	"context"
	"fmt"

	"github.com/issueye/ipc_grpc/grpc/pb"
	"github.com/issueye/ipc_grpc/server"
)

type Clients map[string]*Event

func (c Clients) Add(key string, client *Event) {
	c[key] = client
}

func (c Clients) Del(key string) {
	delete(c, key)
}

func (c Clients) Get(key string) (*Event, bool) {
	client, ok := c[key]
	return client, ok
}

func (c Clients) Len() int {
	return len(c)
}

func (c Clients) PushEvent(key string, t pb.EventType, info *pb.ServerInfo) {
	event, ok := c.Get(key)
	if ok {
		event.EventStream.Send(&pb.EventResponse{
			Type:   t,
			Server: info,
		})
	}
}

type Commands map[string]*Command

func (c Commands) Add(key string, client *Command) {
	c[key] = client
}

func (c Commands) Del(key string) {
	delete(c, key)
}

func (c Commands) Get(key string) (*Command, bool) {
	client, ok := c[key]
	return client, ok
}

func (c Commands) Len() int {
	return len(c)
}

func (c Commands) Show(key string) {
	fmt.Printf("key: %s\n", key)
	command, ok := c.Get(key)
	if ok {
		command.CommandStream.Send(&pb.CommandResponse{Command: "show"})
	}
}

type Event struct {
	EventStream pb.HostHelper_EventServer
	Client      *pb.ClientRequest
}

type Command struct {
	CommandStream pb.HostHelper_CommandServer
	Client        *pb.ClientRequest
}

type Service struct {
	Events   Clients
	Commands Commands
	srv      *server.Server
}

func NewService() (*Service, error) {
	// 初始化服务
	s := &Service{
		Events:   make(Clients),
		Commands: make(Commands),
	}

	srv, err := server.NewServer(true, s, func(pi server.PluginInfo) error { return nil })
	if err != nil {
		return nil, err
	}

	s.srv = srv
	return s, nil
}

// 服务状态
func (srv *Service) Status(ctx context.Context, req *pb.ServerRequest) (*pb.StatusResponse, error) {

	return nil, nil
}

// 服务停止
func (srv *Service) Stop(ctx context.Context, req *pb.ServerRequest) (*pb.Empty, error) {
	return nil, nil
}

// 服务启动
func (srv *Service) Start(ctx context.Context, req *pb.ServerRequest) (*pb.Empty, error) {
	return nil, nil
}

// 服务重启
func (srv *Service) Restart(ctx context.Context, req *pb.ServerRequest) (*pb.Empty, error) {
	return nil, nil
}

// 服务列表
func (srv *Service) List(ctx context.Context, req *pb.Empty) (*pb.ListResponse, error) {

	return nil, nil
}

// 事件推送，服务端持续推送事件
func (srv *Service) Event(client *pb.ClientRequest, stream pb.HostHelper_EventServer) error {
	srv.Events.Add(client.CookieKey, &Event{
		EventStream: stream,
		Client:      client,
	})

	select {}
	return nil
}

// 事件推送，服务端持续推送事件
func (srv *Service) Command(client *pb.ClientRequest, tream pb.HostHelper_CommandServer) error {
	fmt.Println("command", client.CookieKey)
	srv.Commands.Add(client.CookieKey, &Command{
		CommandStream: tream,
		Client:        client,
	})

	select {}
	return nil
}
