// gitTest.cpp : 定义控制台应用程序的入口点。
//


#include "stdafx.h"
#include <vector>
#include <algorithm>
#include <numeric>
using namespace std;

int candy(vector<int> &ratings)
{
	int n = ratings.size();
	vector<int> increment(n);

	for (int i = 1, inc = 1; i < n; i++)
	{
		if (ratings[i]>ratings[i - 1])
			increment[i] = max(inc++, increment[i]);
		else
		{
			inc = 1;
		}
	}
	for (int i = n - 2, inc = 1; i >= 0; i--)
	{
		if (ratings[i]>ratings[i + 1])
			increment[i] = max(inc++, increment[i]);
		else
		{
			inc = 1;
		}
	}

	int sum = n;

	for (size_t i = 0; i < n; i++)
	{
		sum += increment[i];
	}
	return sum;
}

int _tmain(int argc, _TCHAR* argv[])
{

	int a[4] = { 7, 2, 11, 15 };
	vector<int> num(a, a + 4);
	int result = candy(num);

	int b = 10;
	int c = 20;
	return 0;
}



