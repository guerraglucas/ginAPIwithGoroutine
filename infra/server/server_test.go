package server

import (
	"fmt"
	"testing"

	entity "github.com/guerraglucas/ginAPIwithGoroutine/domain/entities"
)

func TestServer(t *testing.T) {
	s := NewServer()
	for i := 0; i < 100; i++ {
		player := &entity.Player{Name: fmt.Sprintf("Player %d", i), Age: i}
		go s.handleNewPlayer(player)
	}
}
