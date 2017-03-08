def print_list(slice):
    print(slice)

def insert_list(slice, args):
    slice.insert(args[0], args[1])

def remove_list(slice, ints):
    slice.remove(ints[0])

def append_list(slice, ints):
    slice.append(ints[0])

def sort_list(slice):
    slice.sort()

def pop_list(slice):
    slice.pop()
def reverse_list(slice):
    slice.reverse()

if __name__ == "__main__":
    n = int(input())
    slice = []
    for _ in range(n):
        args = input().split(' ')
        command = args[0]
        if len(args) > 1:
            ints = list(map(int, args[1:]))

        if command == "insert":
            insert_list(slice, ints)
        elif command == "print":
            print_list(slice)
        elif command == "remove":
            remove_list(slice, ints)
        elif command == "append":
            append_list(slice, ints)
        elif command == "sort":
            sort_list(slice)
        elif command == "pop":
            pop_list(slice)
        elif command == "reverse":
            reverse_list(slice)
        else:
            print("nothing")
