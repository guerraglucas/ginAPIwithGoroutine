package server

import (
	"fmt"

	entity "github.com/guerraglucas/ginAPIwithGoroutine/domain/entities"
)

type Server struct {
	gameState *GameState
}

func NewServer() *Server {
	return &Server{
		gameState: NewGameState(),
	}
}

type GameState struct {
	players []*entity.Player
	msgChan chan any
}

func (gs *GameState) receive(msg any) {
	gs.msgChan <- msg
}

func (gs *GameState) loop() {
	for msg := range gs.msgChan {
		gs.handleMessage(msg)
	}
}

func (gs *GameState) handleMessage(message any) {
	switch msg := message.(type) {
	case *entity.Player:
		gs.addPlayer(msg)
	default:
		panic("Unknown message")
	}
}

func (gs *GameState) addPlayer(player *entity.Player) {
	gs.players = append(gs.players, player)
	fmt.Printf("Added player %s\n", player.Name)
}

func NewGameState() *GameState {
	g := &GameState{
		players: []*entity.Player{},
		msgChan: make(chan any, 100),
	}
	go g.loop()
	return g
}

func (s *Server) handleNewPlayer(player *entity.Player) error {
	s.gameState.receive(player)
	return nil
}
