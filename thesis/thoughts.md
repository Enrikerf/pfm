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