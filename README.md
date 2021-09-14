# XMock

Simply generate mock structs for golang interfaces.

## Install

    go get github.com/MrHuxu/xmock

## Usage

### Mock a file

    $ cat person.go
    package main

    type Person interface {
            Say(msg string)
    }

    $ xmock --file=person.go
    [2021-09-14T15:03:48+08:00]     [INFO]  [xmock/main.go:22]      _args||file=person.go||directory=.||outpkg=mock
    [2021-09-14T15:03:48+08:00]     [INFO]  [xmock/main.go:91]      _generate_done||src=person.go||dist=mock/person.go

    $ cat mock/person.go 
    // Code generated by xmock. DO NOT EDIT.

    package mock

    type Person struct {
            MockSay func(msg string)
    }

    func (p *Person) Say(msg string) {
            p.MockSay(msg)
    }

### Mock a directory