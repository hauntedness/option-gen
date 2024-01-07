# optiongen
a option gen tool for convenience

use it in Promgram
```go
go test -v -run TestExecuteAny github.com/hauntedness/optiongen
=== RUN   TestExecuteAny
=== RUN   TestExecuteAny/github.com/hauntedness/optiongen/internal

type CallOption func(*callOptions)
func (op *callOptions) ApplyOptions(opts ...CallOption) {
        for i := range opts {
                opts[i](op)
        }
}
var WithIntField = func(intField int) CallOption {
        return func(op *callOptions) {
                op.intField = intField
        }
}
var WithStringField = func(stringField string) CallOption {
        return func(op *callOptions) {
                op.stringField = stringField
        }
}
var WithInterfaceField = func(interfaceField interface{}) CallOption {
        return func(op *callOptions) {
                op.interfaceField = interfaceField
        }
}
var WithWriter = func(writer io.Writer) CallOption {
        return func(op *callOptions) {
                op.writer = writer
        }
}
--- PASS: TestExecuteAny (0.76s)
    --- PASS: TestExecuteAny/github.com/hauntedness/optiongen/internal (0.76s)
PASS
ok      github.com/hauntedness/optiongen       1.376s
```


Command line
```go
go run github.com/hauntedness/optiongen/cmd/optiongen -o callOptions -p github.com/hauntedness/optiongen/internal

type CallOption func(*callOptions)
func (op *callOptions) ApplyOptions(opts ...CallOption) {
        for i := range opts {
                opts[i](op)
        }
}
var WithIntField = func(intField int) CallOption {
        return func(op *callOptions) {
                op.intField = intField
        }
}
var WithStringField = func(stringField string) CallOption {
        return func(op *callOptions) {
                op.stringField = stringField
        }
}
var WithInterfaceField = func(interfaceField interface{}) CallOption {
        return func(op *callOptions) {
                op.interfaceField = interfaceField
        }
}
var WithWriter = func(writer io.Writer) CallOption {
        return func(op *callOptions) {
                op.writer = writer
        }
}
```