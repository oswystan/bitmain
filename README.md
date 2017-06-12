## install (for unbuntu only)

```
#install dependencies
sudo apt-get install postgresql
sudo -u postgres createuser $USER
sudo -u postgres createdb $USER

#download source code
go get github.com/oswystan/bitmain

# build
cd bitmain
make

# run
./bitmain
```

## test

```
cd bitmain
make test

```
