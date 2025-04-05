exec:
	podman run -it --network=host docker.restate.dev/restatedev/restate-cli:1.2 $(cmd)
