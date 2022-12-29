# pfm



## debugger

go get -t github.com/google/gops/


## doctrine migrations

* install [doctrine migrate cli](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
* usage [doctrine migrate](https://github.com/golang-migrate/migrate)



dividing the layers I could develop faster because I only care about the domain and make dirty code on:
 * config -> injecting use cases on each controller
 * adapter out database -> deleting in cascade and mapping

and take care later



note that in goland when you are handling errors if you want to return an struct or and error you must return a new instance, so you are forced to instance a not valid in the meaning of domain, structure. In the class-way with the interface you can return nil value that it's more congruent