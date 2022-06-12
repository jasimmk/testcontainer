service_name:=testc
run-test:
	docker build -f Dockerfile.test . -t $(service_name)-tests
	docker run --rm -it \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-e "TC_HOST=host.docker.internal" \
		--add-host host.docker.internal:host-gateway \
		 $(service_name)-tests
	echo "Test completed"
# 172.17.0.1