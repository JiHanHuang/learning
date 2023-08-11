
int __kk(int k, int *b){
    printf("%d,%d",k,*b);
    return 0;
}

weak_alias (__kk, kk);

int main(int argc, char *argv[]){
    int a;
	char *c;
	c = 0;
	printf("%c",*c);
    //__kk(a);
    return 0;
}
