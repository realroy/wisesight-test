start:
	docker-compose up --build -d
	if [[ "$OSTYPE" == "darwin"* ]]; then
		open http://localhost:3000
	fi

stop:
	docker-compose down