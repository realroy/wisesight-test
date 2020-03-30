start:
	wget -P ./db_seed "https://s3-ap-southeast-1.amazonaws.com/wisesight-public/dev-test/rawdata.csv"
	python ./db_seed/clean.py
	docker-compose up --build -d
	open http://localhost:3000

stop:
	docker-compose down