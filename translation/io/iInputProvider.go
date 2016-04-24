package io



type IInputProvider interface{
    GetNext() []string
}