service_name:=testc
run-test:
	docker build . -t $(service_name)-tests
	docker run --rm -p 40000:40000 -e "DOCKER_URL=172.17.0.1" -v /var/run/docker.sock:/var/run/docker.sock $(service_name)-tests
	echo "Test completed"
