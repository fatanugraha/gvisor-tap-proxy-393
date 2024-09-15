# create cidata
docker run -v $(pwd):/mnt cloud-image-utils cloud-localds -v /mnt/cidata.img /mnt/user-data.yaml
