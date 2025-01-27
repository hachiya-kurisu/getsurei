all: getsurei

again: clean all

getsurei: getsurei.go cmd/getsurei/main.go
	go build -o getsurei cmd/getsurei/main.go

clean:
	rm -f getsurei

test:
	go test -cover

cover:
	go test -coverprofile=cover.out
	go tool cover -html cover.out

push:
	got send
	git push github

fmt:
	gofmt -s -w *.go cmd/*/main.go

install:
	install getsurei /usr/local/bin

README.md: README.gmi
	cat README.gmi | sisyphus -a "." -f markdown > README.md

doc: README.md

release: push
	got send -T
	git push github --tags

