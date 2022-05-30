DOCKER:

to build it:
docker build -t dev-walletgo .

to run it:
docker run -it --name debug-image-go -p 3000:3000 -p 40001:40001 -v $PWD:/go/src/github.com/nmarsollier/imagego debug-image-go
