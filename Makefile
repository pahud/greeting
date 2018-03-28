build:
	docker build -t greeting:latest .
tag:
	docker tag  greeting:latest pahud/greeting:latest
push:
	docker push pahud/greeting:latest
