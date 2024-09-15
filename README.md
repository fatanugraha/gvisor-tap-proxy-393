Steps to Reproduce:

```bash
# install vfkit
brew tap cfergeau/crc
brew install vfkit

# clone the current repo
git clone git@github.com:fatanugraha/gvisor-tap-proxy-393.git
cd gvisor-tap-proxy-393

# download ubuntu cloudimg
curl https://cloud-images.ubuntu.com/minimal/releases/noble/release/ubuntu-24.04-minimal-cloudimg-arm64.img -o ubuntu.img

# convert ubuntu.img from qcow2 to raw
qemu-img convert ubuntu.img ubuntu.raw

# build gvisor-tap-vsock from source
git clone git@github.com:containers/gvisor-tap-vsock.git
cd gvisor-tap-vsock
go install ./cmd/gvproxy # or use go build
cd ..

# cross compile the reproduction code
cd repro
CGO_ENABLED=0 GOOS=linux go build .
```

Open 2 terminals:

- 1st terminal:

```bash
bash gvproxy.sh # run gvproxy to listen on /tmp/vfkit.sock
```

- 2nd terminal:

```bash
# you will see the VM GUI
bash run.sh
```

Inside the vm (interact via GUI):

```bash
# login info (see ubuntu/user-data.yaml)
# ubuntu:password

sudo mount -t virtiofs vfkit-share /mnt
/mnt/repro/repro
```
