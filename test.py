import os, sys, argparse

global_args =argparse.Namespace()
argv=None
parser = argparse.ArgumentParser(formatter_class=argparse.RawTextHelpFormatter)
subparsers = parser.add_subparsers(title="command", dest="command")
subparser = subparsers.add_parser("help", help="show help")
global_args = parser.parse_args(argv)

print(global_args)
print(sys.path)

list_=[1, 2, 4, 6]
print(list_,"list")

tuple_=(1,2,3,"adam")
print(tuple_,"tuples are lists that are immutable")

dict_={"ewa":1221,"gregor":1212} 
print(dict_,"hashmap")

set_={1,2,3,1} 
print(set_,"non repeatable variables")

print(os.environ.values())

print(os.environ.get("COMPOSE_PROJECT_DIR", None))

podman_path = "/usr/bin/podman"

print("Is file:",os.path.isfile(podman_path))
print("Is executable:",os.access("podman_path",os.X_OK))

print("Real path",os.path.realpath(podman_path))
print("HelloWorld")
print(sys._getframe().f_lineno,)