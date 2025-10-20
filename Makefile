build: build-lambda build-chart
	echo "Build complete."

build-chart:
	cd weather-chart && npm run build && npm run ship

build-lambda:
	rm -f weather-data/bin/bootstrap
	rm -f cloud/weather-data.zip
	cd weather-data && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o bin/bootstrap main.go

plan:
	cd cloud && terraform init
	cd cloud && terraform plan

deploy: build confirm
	cd cloud && terraform plan
	cd cloud && terraform apply -auto-approve

make ship:
	cd cloud && terraform plan
	cd cloud && terraform apply -auto-approve

confirm: plan
	@read -p "Do you want to apply these changes? (y/n): " ans; \
	if [ "$$ans" != "y" ]; then \
		echo "Aborting apply."; \
		exit 1; \
	fi