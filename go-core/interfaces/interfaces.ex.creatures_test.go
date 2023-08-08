package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Creature interface {
	intro() string
	attack(*int) int
}

type Player struct {
	name   string
	health int
}

type Mob struct {
	name     string
	health   int
	category bool
}

func (p Player) intro() string {
	return p.name
}

func (p Player) attack(m_health *int) int {
	*m_health = *m_health - 50
	return *m_health
}

func (m Mob) intro() string {
	return m.name
}
func (m Mob) attack(p_health *int) int {
	*p_health = *p_health - 30
	return *p_health
}

func TestCreatureInterface(t *testing.T) {
	t.Run("verify interface implementations", func(t *testing.T) {
		player := Player{name: "Steve", health: 100}
		mob := Mob{name: "Zombie", health: 140}

		assert.Equal(t, "Steve", player.intro(), "interface function impl")
		assert.Equal(t, "Zombie", mob.intro(), "interface function impl")

		assert.Equal(t, 100, player.health, "interface function impl")
		assert.Equal(t, 140, mob.health, "interface function impl")

		fmt.Println(player.attack(&mob.health))
		fmt.Println(mob.attack(&player.health))

		assert.Equal(t, 70, player.health, "interface function impl")
		assert.Equal(t, 90, mob.health, "interface function impl")
	})
}
