import threading
from time import sleep, ctime

loops = [2, 4]

def loop(n, nsec):
    print('start', n, 'loop at', ctime())
    sleep(nsec)
    print('loop', n, 'done at', ctime())

def main():
    threads = []
    for i in range(len(loops)):
        t = threading.Thread(target=loop, args=(i, loops[i]))
        threads.append(t)
    for i in range(len(loops)):
        threads[i].start()
    for i in range(len(loops)):
        threads[i].join()
    print("All thread done!!!")

if __name__ == "__main__":
    main()
