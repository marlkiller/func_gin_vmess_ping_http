# func_gin_vmess_ping

#### go mod init {project_name}
#### go mod tidy
#### go mod download

#### go build

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build

#### linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
#### mac
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
#### windows
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build 

#####
go.fc.voidm.com?vmess=eyJhZGQiOiI1NC4xNjkuMTg4LjIzNSIsImFpZCI6IjAiLCJob3N0IjoiIiwiaWQiOiI1ZDQ4OTNhMC0xOGQ1LTExZWItYTUwMS0wMjk0MDViYjkyMGUiLCJuZXQiOiJ0Y3AiLCJwYXRoIjoiIiwicG9ydCI6IjMzMDYiLCJwcyI6ImF3cy10Y3AiLCJ0bHMiOiIiLCJ0eXBlIjoibm9uZSIsInYiOiIyIn0=

go.fc.voidm.com?instance


##### zip

tar -zcvf go_dev-code.zip ./*


