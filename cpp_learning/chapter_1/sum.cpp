#include <iostream>

using namespace std;

int main()
{
    int value = 0, sum = 0;
    while (cin >> value)
    {
        sum += value;
    }
    cout << "Sum:" << sum << endl;
    return 0;
}