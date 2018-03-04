# TODO

## build todo

> Install go-spirit command before you build todo

#### Install go-spirit command

```bash
go get github.com/go-spirit/go-spirit
go install github.com/go-spirit/go-spirit
```

#### Build

```bash
> go-spirit build --config configs/build-todo.conf
```

#### Make sure the components are already registered

```bash
> ./todo list components
- function: github.com/go-spirit/spirit/component/function.newComponentFunc
- mns: github.com/go-spirit/spirit/component/mns.NewMNSComponent
- examples-todo: github.com/spirit-component/examples/todo.NewTodo
```


#### Run components

```
> cd $GOPATH/src/github.com/spirit-component/examples/todo
> mv $GOPATH/src/github.com/go-spirit/go-spirit/todo .
> ./todo run --config configs/todo.conf
```

### Work with post api

```
> go-spirit build --config configs/build-postapi.conf
> ./postapi run --config configs/build-postapi.conf
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


#### work flow

![](https://raw.githubusercontent.com/spirit-component/draft-assets/master/images/examples/todo/flow.png)



## combine the postapi and todo


#### Build

```bash
> go-spirit build --config configs/build-standalone.conf
```

#### Make sure the components are already registered

```bash
> ./todo list components
- function: github.com/go-spirit/spirit/component/function.newComponentFunc
- mns: github.com/go-spirit/spirit/component/mns.NewMNSComponent
- examples-todo: github.com/spirit-component/examples/todo.NewTodo
- post-api: github.com/spirit-component/postapi.NewPostAPI
```


#### Run components

```
> cd $GOPATH/src/github.com/spirit-component/examples/todo
> mv $GOPATH/src/github.com/go-spirit/go-spirit/todo .
> ./todo run --config configs/todo-standalone.conf
```


#### work flow

![](https://raw.githubusercontent.com/spirit-component/draft-assets/master/images/examples/todo/flow-standalone.png)

