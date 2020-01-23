PROJECT_NAME = rajnikanthjokes

build:
	GOOS=linux go build -o main

test:
	@go test -v main_test.go

clean:
	@echo "cleaning"
	@rm -f main main.zip

zip: clean build
	@echo "zipping"
	@zip -r -9 --quiet main.zip main

publish: zip
	@echo "publishing ${PROJECT_NAME}"
	@aws lambda update-function-code --function-name ${PROJECT_NAME} --zip-file fileb://main.zip
	@echo finished
