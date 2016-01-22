
# Steal This

Use the go tool to get this repo

```
go get github.com/anxiousmodernman/stats
```

Open the project in Atom.

```
cd $GOPATH/src/github.com/anxiousmodernman/stats
atom .
```

My style for buildable binaries is to keep them under the `/cmd` folder. Right now there's
only one. So do this to build it.

```
cd cmd/stats
go build
```

A binary will appear called **stats**, because it takes the name of the directory that
contains it. Run it like this:

```
./stats
```

It will fail, of course. You will have to edit the code to make it work.

