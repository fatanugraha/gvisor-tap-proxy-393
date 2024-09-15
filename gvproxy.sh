rm /tmp/vfkit.sock
gvproxy -debug -listen unix:///tmp/network.sock --listen-vfkit unixgram:///tmp/vfkit.sock
