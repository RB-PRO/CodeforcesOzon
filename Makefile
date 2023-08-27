all:
	go test -v -timeout 30s -run ^TestMain$ quickstart

run:
	go run main.go

	
push:
	git push git@github.com:RB-PRO/CodeforcesOzon.git
