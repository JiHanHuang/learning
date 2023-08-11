#include <sys/stat.h>
#include <sys/wait.h>
#include <fcntl.h>
#include <stdio.h>
#include <string.h> 

const char fifo_file[] = "/tmp/fifo_file";

int main(int argc, char *argv[]){
    int fd = 0;
    int pid1=0, pid2=0;
    char buf[128] = {0};
    mkfifo(fifo_file,O_CREAT |O_RDWR);
    pid1 = fork();
    if (pid1  > 0 ) {
        pid2 = fork();
    }
    if (pid1 > 0 && pid2 > 0 ){ //parent
        // 读写方式打开，read函数则会阻塞，一直等待数据来临
        // 只读方式打开，则子进程写完(也可能没写完)，read函数返回0值，退出
        //fd = open(fifo_file, O_RDWR);
        fd = open(fifo_file, O_RDONLY);
        //确保2个进程都写完
        int status;
        waitpid(pid1,&status,0);
        waitpid(pid2,&status,0);
        while (read(fd, buf, sizeof(buf) - 1) > 0)
        {
            printf("read:%s", buf);
            memset(buf, 0, sizeof(buf));
        }
    }else if(pid1 < 0 || pid2 < 0) { //error
        printf("fork error");
        return -1;
    }else{ //children
        fd = open(fifo_file, O_WRONLY);
        sprintf(buf, "hello: %d, %d\n",pid1, pid2);
        write(fd, buf, strlen(buf));
    }
    return 0;
}