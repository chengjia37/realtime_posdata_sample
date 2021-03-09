#!/bin/sh

for i in gen_csv read_csv
do
	cd $i
	GOOS=windows GOARCH=386 go build -o ${i}.exe ${i}.go
	cd ..
done
