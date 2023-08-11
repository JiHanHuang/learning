#include <stdio.h>
#include <sys/wait.h>
#include <unistd.h>

#define RUN_CMD "/usr/bin/date"

int main(int argc, char *argv[]){
    int fd[2];
    int pid;
    char output[128] = {0};

    if (pipe(fd) < 0 ){
        printf("pipe error");
        return -1;
    }
    
    pid = fork();
    if (pid <0){
        printf("fork error");
        return -1;
    }else if (pid > 0){ //parent
        close(fd[1]);
        if (read(fd[0], output, sizeof(output)-1) < 0){
            printf("read error");
            return -1;
        }
        printf("Parent:\n%s", output);
    }else{ //child
        close(fd[0]);
        dup2(fd[1], STDOUT_FILENO);
        execv(RUN_CMD, argv);
    }
    return 0;
}