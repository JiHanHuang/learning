#include <stdio.h>
#include <sys/msg.h>
#include <sys/wait.h>
#include <unistd.h>
#include <string.h>
#include <errno.h>

#define msg_bytes 512

typedef struct
{
    long type;
    char text[msg_bytes];
} msg;


int main(int argc, char *argv[]){
    int pid, pid1=0, pid2=0;
    int ret,msgid;

    msgid = msgget(IPC_PRIVATE,0666);
    if (msgid < 0){
        printf("msgget error\n");
        return -1;
    }
    pid1 = fork();
    if (pid1  > 0 ) {
        pid2 = fork();
    }
    if (pid1 > 0 && pid2 > 0 ){ //parent
        pid = getpid();
        msg tmp;
        //阻塞方式
        //while ((ret = msgrcv(msgid, (void *)&tmp, msg_bytes, 1, 0)) > 0){
        //非阻塞方式
        while ((ret = msgrcv(msgid, (void *)&tmp, msg_bytes, 1, IPC_NOWAIT))){
            if ((ret < 0) && (errno != ENOMSG)) {
                break;
            }
            printf("%d, %d read:%s\n", pid, ret, tmp.text);
            memset(&tmp.text, 0, sizeof(tmp.text));
            sleep(2);
        }
        printf("msg receive over. %s\n", strerror(errno));
        int status;
        waitpid(pid1,&status,0);
        waitpid(pid2,&status,0);
        printf("parent process[%d] exit\n", pid);
    }else if(pid1 < 0 || pid2 < 0) { //error
        printf("fork error\n");
        return -1;
    }else{ //children
        int i;
        pid = getpid();
        msg tmp = {.type = 1};
        for (i = 0; i < 5;i++){
            tmp.type = 1;
            sprintf(tmp.text, "hello: %d, N:%d\n", pid, i);
            if ((ret = msgsnd(msgid, &tmp, msg_bytes, 0)) < 0){
                printf("msg send error:%s\n",strerror(errno));
                return -1;
            }
            memset(&tmp.text, 0, sizeof(tmp.text));
            sleep(1);
        }
        printf("child process[%d] exit\n", pid);
    }
    return 0;
}