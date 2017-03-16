import time

timeinput = "07:05:45PM"

# note the %I here, it should only be used when there is am pm notation
struct_time = time.strptime(timeinput, "%I:%M:%S%p")
print(time.strftime("%H:%M:%S", struct_time))
