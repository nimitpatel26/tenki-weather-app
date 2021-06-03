build:
	npm install
	npm run build
	mkdir -p functions
	go get ./go/api/weather/...
	go build -o functions/weather ./go/api/weather/...
	go get ./go/api/news/...
	go build -o functions/news ./go/api/news/...
	go get ./go/api/about/...
	go build -o functions/about ./go/api/about/...