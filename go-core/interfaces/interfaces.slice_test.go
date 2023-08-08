package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatureInterfaceSlice(t *testing.T) {
	t.Run("interface slices", func(t *testing.T) {
		entity := []Creature{Player{name: "A"}, Mob{name: "X"}, Mob{name: "Y"}, Player{name: "B"}}

		assert.Equal(t, "A", entity[0].intro(), "verify interface slice item")
		assert.Equal(t, "X", entity[1].intro(), "verify interface slice item")
		assert.Equal(t, "Y", entity[2].intro(), "verify interface slice item")
		assert.Equal(t, "B", entity[3].intro(), "verify interface slice item")
	})
}
