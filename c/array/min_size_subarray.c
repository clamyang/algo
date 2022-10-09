// # 209 最小长度子数组

#include <stdio.h>
#include <stdlib.h>

int min_size_subarr(int nums[], int length, int target);
int min_size_sliding_window(int nums[], int length, int target);

int main(int argc, int *argv[])
{
  int nums[] = {1, 2, 5, 2, 6};
  int target = 10;

  int length = sizeof(nums) / sizeof(nums[0]);

  // printf("%d\n", min_size_subarr(nums, length, target));
  printf("%d\n", min_size_sliding_window(nums, length, target));
  return 0;
}

// easy way
int min_size_subarr(int nums[], int length, int target)
{
  int min_val = 0;

  for (int i = 0; i < length; i++)
  {
    int sum = nums[i];

    for (int j = i + 1; j < length; j++)
    {
      sum += nums[j];
      if (sum >= target)
      {
        int tmp = j - i + 1;
        if (min_val == 0 || tmp < min_val)
        {
          min_val = tmp;
        }
        break;
      }
    }
  }
  return min_val;
}

int min_size_sliding_window(int nums[], int length, int target)
{
  int sum = 0;
  int i = 0, j = 0;
  int min_val = length;

  for (j = 0; j < length; j++)
  {
    sum += nums[j];
    while (sum >= target)
    {
      if (min_val > (j - i + 1))
      {
        min_val = (j - i + 1);
      }

      sum -= nums[i++];
    }
  }

  return min_val;
}