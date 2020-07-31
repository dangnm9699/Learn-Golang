go run ./server/main.go -port=8010 -tz=US/Eastern & 
go run ./server/main.go -port=8020 -tz=Asia/Tokyo & 
go run ./server/main.go -port=8030 -tz=Europe/London & 
go run ./client/main.go New_York=localhost:8010 Tokyo=localhost:8020 London=localhost:8030