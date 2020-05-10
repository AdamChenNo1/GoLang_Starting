#include <iostream>

using namespace std;

extern "C" {    //向下兼容C的命名规范
    void test() {
        cout << "this is a test" << endl;
    }
}