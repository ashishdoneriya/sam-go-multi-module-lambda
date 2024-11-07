all:
	cd service1 && go build
	cd service2 && go build

clean:
	rm -f service1/service1 service2/service2

deploy:
	sam build && sam deploy