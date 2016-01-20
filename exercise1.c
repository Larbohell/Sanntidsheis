// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o exercise1 exercise1.c -lpthread

#include <pthread.h>
#include <stdio.h>
#include <semaphore.h>

int i = 0;
sem_t sem;

void* thread_1_func(){
	for (int j = 0; j<1000000; j++){
		sem_wait(&sem);
		i++;
		sem_post(&sem);
	}
	return NULL;
}

void* thread_2_func(){
	
	for (int j = 0; j<1000000; j++){
		sem_wait(&sem);
		i--;
		sem_post(&sem);
	}
	
	return NULL;
}


int main(){

	
	sem_init(&sem, 0, 1);

	pthread_t thread_1;
	pthread_create(&thread_1, NULL, thread_1_func, &i);

	pthread_t thread_2;
	pthread_create(&thread_2, NULL, thread_2_func, &i);

	pthread_join(thread_1, NULL);
	pthread_join(thread_2, NULL);

	printf("%d\n", i);
	
	return 0;
}

