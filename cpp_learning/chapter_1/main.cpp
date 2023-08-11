#include <iostream>
#include "sales.hpp"

using namespace std;

int main()
{
    Sales_item total;
    cout << "书店售卖系统Demo"<<endl;
    cout << "请输入售卖的书编号，数量，单价；如：1-2-X 1 12："<<endl;
    if (cin >> total)
    {
        Sales_item trans;
        while (cin >> trans)
        {
            if (total.isbn() == trans.isbn())
            {
                total += trans;
            }
            else
            {
                cout << "出售统计：" << total << endl;
                total = trans;
            }
        }
        cout << "出售统计：" <<total << endl;
    }
    else
    {
        cerr << "No data?" << endl;
        return -1;
    }
    return 0;
}