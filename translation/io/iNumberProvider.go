package io



type INumberProvider interface{
    GetNext() []string
}