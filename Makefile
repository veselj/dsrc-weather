build:
	cd weather-chart && npm run build && npm run ship
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