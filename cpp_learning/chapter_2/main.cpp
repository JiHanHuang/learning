#include <iostream>
#include <typeinfo>

using namespace std;

int main()
{
    int kk = 1;
    double dval = 10.2;
    double dval2 = 20.2;
    double &rval = dval;
    dval = 11;
    rval = dval2;
    dval2 = 22;
    rval = kk;
    kk = 98;
    cout << typeid(rval).name() << ":" << rval << endl;

    int gg = 22;

    // ival 指向的值不能被修改
    const int *ival = &kk;
    //*ival = 33; //这种编译不过
    ival = &gg;

    // ival2 指针本身是个常量，不能被修改
    int *const ival2 = &kk;
    // ival2 = &cint; //这种编译不过
    *ival2 = 33;
    cout << *ival << ":" << *ival2 << "-" << gg << ":" << kk << endl;
}