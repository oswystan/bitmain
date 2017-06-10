## install (for unbuntu only)

```

#install dependencies
sudo apt-get install postgresql

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
./run_st.sh

```