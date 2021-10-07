# restSanity
A very bad way of testing things are working


### Kube Deploy
 - `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o restsanity`
 - Build dockerfile `docker build . -t cloud.canister.io:5000/brudtech/restsanity:latest`
 - `docker push cloud.canister.io:5000/brudtech/restsanity:latest` 
 - `kubectl rollout restart deploy restsanity-dpy` 