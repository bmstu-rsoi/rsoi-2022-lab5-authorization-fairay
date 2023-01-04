all: gateway flights tickets privileges

# Creating images
IMAGES?=gateway flights tickets privileges
$(IMAGES):
	cd ./src/$@ && docker build --no-cache -t fairay/rsoi-lab4-$@ . && docker push fairay/rsoi-lab4-$@:latest
