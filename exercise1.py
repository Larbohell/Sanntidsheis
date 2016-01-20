# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread
from threading import Lock

i = 0

def thread_1_func(lock):
    	global i
	
	for j in range(0,1000000):
		lock.acquire()
		i+=1 
		lock.release()

def thread_2_func(lock):
    	global i
	
	for j in range(0,1000000):
		lock.acquire()
		i-=1 
		lock.release()
	

def main():
	lock = Lock()
	thread_1 = Thread(target = thread_1_func, args = ([lock]))
	thread_2 = Thread(target = thread_2_func, args = ([lock]))
	thread_1.start()
	thread_2.start()

	thread_1.join()
	thread_2.join()

	print i


main()
