from time import sleep, ctime
import thread

def loop0():
    print("start loop 0", ctime())
    sleep(4)
    print("loop 0 done", ctime())

def loop1():
    print("start loop 1", ctime())
    sleep(4)
    print("loop 1 done", ctime())

def main():
    print("start loop....", ctime())
    thread.start_new_thread(loop0, ())
    thread.start_new_thread(loop1, ())
    #loop0()
    #loop1()
    sleep(9)
    print("all loop done", ctime())

if __name__ == '__main__':
    main()
