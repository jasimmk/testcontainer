# Test container debugging using delv and IntelliJ


## Base documentation
 Documentation is based on this [article on golang](https://blog.jetbrains.com/go/2018/04/30/debugging-containerized-go-applications/)
 
### Setup Dockerfile

Create a dockerfile as is

![](./docs/delv-dockerfile.png)

### Setup Dockerfile 

At intellij go to "Edit Runconfiguration" > "Add New configuration" > "Docker" > Dockerfile
![](./docs/docker-file.png)

### Setup Go Remote 

"Edit Runconfiguration" > "Add New configuration" > "Goremote"

![](./docs/go-remote-file.png)


### Running application
- Run `sample` app which will create docker containers ready
- Create break point in main application. And run go-remote, it breaks at needed breakpoints

### Running tests from docker
```
make run-test
```
### Configuring testcontainer-go to use docker host
Getting docker host from containers are tricky
Check this [thread at stack overflow](https://stackoverflow.com/questions/24319662/from-inside-of-a-docker-container-how-do-i-connect-to-the-localhost-of-the-mach)

Make sure you start docker container with following options so that reaper will work
```
docker run --rm -it \
    -v /var/run/docker.sock:/var/run/docker.sock \
    -e "TC_HOST=host.docker.internal" \
    --add-host host.docker.internal:host-gateway \
     $(service_name)-tests
```
