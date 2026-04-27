package main

import "testing"

type mockScanner struct {
	inputs []string
	index  int
}

func newMockScanner(inputs []string) *mockScanner {
	return &mockScanner{inputs: inputs}
}

func (m *mockScanner) Scan() bool {
	return m.index < len(m.inputs)
}

func (m *mockScanner) Text() string {
	text := m.inputs[m.index]
	m.index++
	return text
}

func TestAddMovie(t *testing.T) {
	movies := []movie{
		{id: 1, name: "Test", quantity: 1},
	}

	// simulando input
	scanner = newMockScanner([]string{"New Movie", "5"})

	result := add_movies(movies)

	if len(result) != 2 {
		t.Errorf("esperado 2 filmes, recebeu %d", len(result))
	}

	if result[1].name != "New Movie" {
		t.Errorf("nome incorreto")
	}
}

func TestAddMovieInvalidQuantity(t *testing.T) {
	movies := []movie{
		{id: 1, name: "Test", quantity: 1},
	}

	scanner = newMockScanner([]string{"Movie", "abc"})

	result := add_movies(movies)

	if len(result) != 1 {
		t.Errorf("não deveria adicionar filme com quantidade inválida")
	}
}

func TestDeleteMovie(t *testing.T) {
	movies := []movie{
		{id: 1, name: "A", quantity: 1},
		{id: 2, name: "B", quantity: 2},
	}

	scanner = newMockScanner([]string{"1"})

	result := delete_movies(movies)

	if len(result) != 1 {
		t.Errorf("esperado 1 filme após deleção")
	}

	if result[0].id != 2 {
		t.Errorf("filme errado removido")
	}
}

func TestUpdateMovieName(t *testing.T) {
	movies := []movie{
		{id: 1, name: "Old", quantity: 1},
	}

	scanner = newMockScanner([]string{"1", "New Name", ""})

	result := update_movie(movies)

	if result[0].name != "New Name" {
		t.Errorf("nome não atualizado")
	}
}

func TestUpdateMovieQuantity(t *testing.T) {
	movies := []movie{
		{id: 1, name: "Movie", quantity: 1},
	}

	scanner = newMockScanner([]string{"1", "", "10"})

	result := update_movie(movies)

	if result[0].quantity != 10 {
		t.Errorf("quantidade não atualizada")
	}
}

