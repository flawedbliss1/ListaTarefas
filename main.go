package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ErrTarefaNaoEncontrada error = errors.New("tarefa não encontrada")

type Tarefa struct {
	Id     int
	Titulo string
	Feita  bool
}

type Lista struct {
	Tarefas   []Tarefa
	ProximoID int
}

func (l *Lista) Adicionar(titulo string) {
	tarefa := Tarefa{Id: l.ProximoID, Titulo: titulo, Feita: false}
	l.Tarefas = append(l.Tarefas, tarefa)
	l.ProximoID++
}

func (l *Lista) marcarFeita(id int) error {
	for i := range l.Tarefas {
		if l.Tarefas[i].Id == id {
			l.Tarefas[i].Feita = true
			return nil
		}
	}
	return ErrTarefaNaoEncontrada
}

func (l *Lista) Remover(id int) error {
	for i := range l.Tarefas {
		if l.Tarefas[i].Id == id {
			l.Tarefas = append(l.Tarefas[:i], l.Tarefas[i+1:]...)
			fmt.Println("Removido: ", id)
			return nil
		}
	}
	return ErrTarefaNaoEncontrada
}

func (l Lista) Imprimir() {
	for _, t := range l.Tarefas {
		if t.Feita {
			fmt.Printf("%d [x] %s\n", t.Id, t.Titulo)
		} else {
			fmt.Printf("%d [ ] %s\n", t.Id, t.Titulo)
		}
	}
}

func (l *Lista) Salvar() error {
	var b strings.Builder
	for _, t := range l.Tarefas {
		linha := fmt.Sprintf("%d|%s|%t \n", t.Id, t.Titulo, t.Feita)
		b.WriteString(linha)
	}
	conteudo := b.String()
	err := os.WriteFile("tarefas.txt", []byte(conteudo), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (l *Lista) Carregar() error {
	data, err := os.ReadFile("tarefas.txt")
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}

	linhas := strings.Split(string(data), "\n")
	for _, linha := range linhas {
		if linha == "" {
			continue
		}
		partes := strings.Split(linha, "|")
		id, _ := strconv.Atoi(partes[0])
		feita, _ := strconv.ParseBool(partes[2])
		titulo := partes[1]
		t := Tarefa{Id: id, Titulo: titulo, Feita: feita}
		l.Tarefas = append(l.Tarefas, t)
	}
	l.ProximoID = len(l.Tarefas) + 1
	return nil
}

func main() {
	li := Lista{ProximoID: 1}
	li.Carregar()

	li.Adicionar("comprar leite")
	li.Adicionar("pagar conta")
	li.Salvar()

	li.Imprimir()
}
