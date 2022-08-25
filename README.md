# go-template-client

### Steps to run it
* first build & deploy [go-template](https://github.com/juanpablopizarro/go-template)
* clone this repo and go to the root folder
* be sure about to run `eval $(minikube docker-env)` first in the same terminal
* run `docker build -t demo-client -f build/Dockerfile .`
* run `kubectl apply -f deploy/deployment.yaml`
* run `kubectl get pods`
* run `kubectl logs` + the client pod retrieved in the previous command. i.e.: `kubectl logs client-deployment-58b84d6f8d-clh6j`
* you must see something like:
```bash
http://go-service:8080/hash
{"hash":"9fceb8c1411452ea1857eb0891d9d8f5"}
```

As you can see, pod to pod communication is made using the service name (check the `service.yaml` file)
