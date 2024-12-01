%: ## Build all commands 
	CGO_ENABLED=0 go build -ldflags "-w"  -o build/$@ cmd/$@/main.go

build-lambda: ## Build strava auth script for AWS Lambda
	GOOS=linux GOARCH=arm64 go build -ldflags "-w" -tags lambda.norpc -o build/bootstrap cmd/lambda-strava-auth/main.go

package-lambda: ## Zip strava auth script
	zip -j build/function.zip build/bootstrap

build-and-package-lambda: build-lambda package-lambda