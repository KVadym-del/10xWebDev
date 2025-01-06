#include "calc.h"

int32_t fibonacci(int32_t n)
{
    if (n <= 1)
        return n;
    int32_t prev = 0, curr = 1;
    for (int32_t i = 2; i <= n; i++)
    {
        int32_t next = prev + curr;
        prev = curr;
        curr = next;
    }
    return curr;
}

double complex_calculation(double x, double y)
{
    return (x * x + y * y) / (x + y);
}