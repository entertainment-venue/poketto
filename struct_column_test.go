package poketto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type User struct {
	Id   int
	Name string
}

func TestStructColumnForMap(t *testing.T) {
	u1 := User{
		Id:   1,
		Name: "uzi",
	}
	u2 := User{
		Id:   2,
		Name: "faker",
	}
	users := []User{u1, u2}

	var desk map[int]User
	if err := StructColumn(&desk, users, "", "Id"); err != nil {
		assert.Fail(t, err.Error())
		return
	}
	assert.Equal(t, len(desk), 2)
	assert.Equal(t, desk[1].Name, u1.Name)
	assert.Equal(t, desk[2].Name, u2.Name)

	var desk2 map[int]string
	if err := StructColumn(&desk2, users, "Name", "Id"); err != nil {
		assert.Fail(t, err.Error())
		return
	}
	assert.Equal(t, len(desk2), 2)
	assert.Equal(t, desk2[1], u1.Name)
	assert.Equal(t, desk2[2], u2.Name)
}

func TestStructColumnForSlice(t *testing.T) {
	u1 := User{
		Id:   1,
		Name: "uzi",
	}
	u2 := User{
		Id:   2,
		Name: "faker",
	}
	users := []User{u1, u2}

	var desk []int
	if err := StructColumn(&desk, users, "Id", ""); err != nil {
		assert.Fail(t, err.Error())
		return
	}
	assert.Equal(t, len(desk), 2)
	assert.Equal(t, desk[0], 1)
	assert.Equal(t, desk[1], 2)

	var desk2 []string
	if err := StructColumn(&desk2, users, "Name", ""); err != nil {
		assert.Fail(t, err.Error())
		return
	}
	assert.Equal(t, len(desk2), 2)
	assert.Equal(t, desk2[0], "uzi")
	assert.Equal(t, desk2[1], "faker")
}
