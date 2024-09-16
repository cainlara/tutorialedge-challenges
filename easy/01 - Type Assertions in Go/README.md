# Challenge 01 - Type Assertions in Go

In this challenge, we are going to become familiar with the concept of `Type Assertions` in Go!

## Preface
If you are new to the language, then type assertions are a concept that can sometimes trip you up and appear a little tricky at first, but after overcoming the syntax it becomes far easier to understand.

Through using type assertions, we can retrieve the dynamic value of an interface. For example:

```
var myName interface{} = "Elliot"

name := myName.(string)
fmt.Println(name)
```


In this example, we have an interface which has a dynamic value of “Elliot”. We can then use a _type assertion_ to retrieve this dynamic value and use the value just like we would any other `string` value in Go.

## Challenge
In this challenge, we are going to define a function that is called `GetDeveloper` which will take in 2 `interface{}` arguments.

Within this function, you will have to declare a new `Developer` instance and use type assertion to populate the values correctly before then returning this new `Developer` instance.