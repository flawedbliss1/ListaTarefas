Projeto pra praticar structs, métodos, tratamento de erro e I/O de arquivo.

## O que faz

- Adiciona tarefas com título
- Marca tarefa como feita pelo ID
- Remove tarefa pelo ID
- Salva as tarefas em `tarefas.txt`
- Carrega as tarefas do arquivo ao iniciar

## Como rodar

```
go run main.go
```

O arquivo `tarefas.txt` é criado automaticamente na primeira execução.
Rodar de novo carrega as tarefas que já estavam salvas.

## Aprendizado

- **Struct e método**: Métodos com receiver de ponteiro pra
  modificar e de valor pra ler.
- **Slice**: `append` pra adicionar, `append(s[:i], s[i+1:]...)` pra remover.
- **Erro**: Gestão de erros.
- **I/O de arquivo**: `os.ReadFile`, `os.WriteFile`, `errors.Is(err, os.ErrNotExist)`.
