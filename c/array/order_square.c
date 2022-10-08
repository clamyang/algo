#include <stdio.h>
#include <stdlib.h>

void order_square(int nums[], int length);
void order_square_hard(int nums[], int length);
void order_square_fault(int nums[], int length);

int main(int argc, char *argv[])
{
  int nums[] = {-12, -4, -3, 0, 1, 2, 11};

  int length = sizeof(nums) / sizeof(nums[0]);

  // order_square(nums, length);
  // order_square_hard(nums, length);
  order_square_fault(nums, length);

  for (int i = 0; i < length; i++)
  {
    printf("%d\n", nums[i]);
  }

  return 0;
}

int cmp_int(const void *p1, const void *p2) /*参数格式为 无符号型*/
{
  int *a = (int *)p1; /*强制类型转换*/
  int *b = (int *)p2;
  return *a - *b; /* a > b时升序排列  a < b 时降序排列*/
}

void order_square(int nums[], int length)
{
  // square
  for (int i = 0; i < length; i++)
  {
    nums[i] *= nums[i];
  }

  // sort
  qsort(nums, length, sizeof(nums[0]), cmp_int);
}

void order_square_hard(int nums[], int length)
{
  // i point to left
  // j point to right

  int i, j;

  i = 0;
  j = length - 1;

  int *ans = (int *)malloc(sizeof(int) * length);
  for (int index = j; index >= 0; index--)
  {
    int left = nums[i] * nums[i];
    int right = nums[j] * nums[j];

    if (left > right)
    {
      ans[index] = left;
      i++;
    }
    else
    {
      ans[index] = right;
      j--;
    }
  }
}

// when j point to neg value, algo is wrong
// in this case we must have a external array
void order_square_fault(int nums[], int length)
{
  int i, j;

  i = 0;
  j = length - 1;

  // int nums[] = {-12, -4, -3, 0, 1, 2, 11};
  for (int index = j; index >= 0; index--)
  {
    int left = nums[i] * nums[i];
    int right = nums[j] * nums[j];

    if (left > right)
    {
      nums[i] = nums[index];
      nums[index] = left;
      j--;
    }
    else
    {
      nums[index] = right;
      j--;
    }
  }
}
