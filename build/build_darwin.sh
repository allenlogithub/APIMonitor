rm -r /usr/local/Cellar/go/1.16.2/libexec/src/modules/*
cp -r modules /usr/local/Cellar/go/1.16.2/libexec/src/
env GOOS=darwin GOARCH=amd64 go build -v main.go
