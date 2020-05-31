
go build -o htree.so -buildmode=c-shared ../main.go

g++ main.cpp -L ./ -l:htree.so

./a.exe -htree=true