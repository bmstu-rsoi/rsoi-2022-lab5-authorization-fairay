all: gateway flights tickets privileges

# Creating images
IMAGES?=gateway flights tickets privileges
$(IMAGES):
	cd ./src/$@ && docker build --no-cache -t fairay/rsoi-lab5-$@ . && docker push fairay/rsoi-lab5-$@:latest
