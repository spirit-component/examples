# TODO

## build todo

#### create imports file

imports_todo.go

```go
package main

import (
	_ "github.com/go-spirit/spirit/component/mns"
	_ "github.com/spirit-component/examples/todo"
)
```

#### copy imports file to go-spirit dir

```bash
> cp imports_todo.go $GOPATH/src/github.com/go-spirit/go-spirit 
```

#### build

```bash
> cd $GOPATH/src/github.com/go-spirit/go-spirit
> go build -o todo ./main.go ./imports_todo.go
```

#### make sure the components are already registered

```bash
> ./todo list components
- function: github.com/go-spirit/spirit/component/function.newComponentFunc
- mns: github.com/go-spirit/spirit/component/mns.NewMNSComponent
- examples-todo: github.com/spirit-component/examples/todo.NewTodo
```


#### run components

```
> cd $GOPATH/src/github.com/spirit-component/examples/todo
> mv $GOPATH/src/github.com/go-spirit/go-spirit/todo .
> ./todo run --config configs/todo.conf
```

### work with post api

```
./postapi run --config configs/postapi.conf
```

#### todo.task.new

```bash
curl -X POST \
  http://127.0.0.1:8080/ \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'x-api: todo.task.new' \
  -d '{
	"name":"hello task"
}'
```


```json
{
    "code": 0,
    "result": {
        "id": "68be59dc-a0a2-4186-a476-337c4fd593bd"
    }
}
```

#### todo.task.get

```bash
curl -X POST \
  http://127.0.0.1:8080/ \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'x-api: todo.task.get' \
  -d '{
	"id":"68be59dc-a0a2-4186-a476-337c4fd593bd"
}'
```

```json
{
    "code": 0,
    "result": {
        "id": "68be59dc-a0a2-4186-a476-337c4fd593bd",
        "name": "hello task"
    }
}
```


## work flow

![](https://raw.githubusercontent.com/spirit-component/draft-assets/master/images/examples/todo/flow.png)
