## Go installation

### apt 

    $ sudo apt install golang-go

### repo

    $ wget "https://golang.org/dl/go${VERSION}.linux-${ARCH}.tar.gz"
    $ tar -xf "go${VERSION}.linux-${ARCH}.tar.gz"
    $ sudo mv -v go /usr/local
    $ vim ~/.bashrc


    # set up Go lang path #
    export GOPATH=$HOME/go
    export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

    $ source ~/.bashrc

### Goland

you could download go from goland

![asedfasdf](golandSdk.png)


And then edit gopath

    $ vim /.bashrc

    PATH="$PATH:/home/kerf/sdk/go1.17.6/bin"
