import os, sys, argparse

global_args =argparse.Namespace()


print(global_args)
print(sys.path)

podman_path = "/usr/bin/podman"

print("Is file:",os.path.isfile(podman_path))
print("Is executable:",os.access("podman_path",os.X_OK))

print("Real path",os.path.realpath(podman_path))
print("HelloWorld")