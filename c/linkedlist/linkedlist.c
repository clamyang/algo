#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

struct Node
{
  struct Node *next;
  int val;
  int len;
};

// 打印节点值
int print_val(struct Node *);
// 遍历链表
void traverse(struct Node *);
// 初始化链表
void linkedlist_init(struct Node *head);
// 在特定位置插入
void insert_at(struct Node *head, int val, int index);
// 在特定位置删除
void delete_at(struct Node *head, int index);
// 判断是否有环
bool is_cycle(struct Node *head);
// 获取最后一个节点
struct Node *get_last(struct Node *head);
// 反转链表
struct Node *reverse(struct Node *head);
// 合并两个有序链表
struct Node *merge_2_sorted(struct Node *h1, struct Node *h2);

int main(void)
{
  struct Node *head = (struct Node *)malloc(sizeof(struct Node));
  head->val = 1000;
  head->len = 0;
  head->next = NULL;

  linkedlist_init(head);
  // traverse(head);

  insert_at(head, 256, 3);
  // traverse(head);

  // true == 1
  // false == 0
  // printf("%d\n", is_cycle(head));

  // struct Node * last = get_last(head);
  //  make cycle
  // last->next = head;

  printf("%d\n", is_cycle(head));

  struct Node *h1;
  h1 = reverse(head);

  // traverse(h1);

  printf("*********\n");

  struct Node *head1 = (struct Node *)malloc(sizeof(struct Node));
  head1->val = 100;
  head1->len = 0;
  head1->next = NULL;
  linkedlist_init(head1);

  traverse(merge_2_sorted(head, head1));
  return 0;
}

void linkedlist_init(struct Node *head)
{
  struct Node *cur = head;

  for (int i = 1; i <= 5; i++)
  {
    struct Node *node = (struct Node *)malloc(sizeof(struct Node));
    node->val = i + 1000 + head->val;
    node->next = NULL;

    cur->next = node;
    cur = cur->next;

    head->len++;
  }

  return;
}

void traverse(struct Node *head)
{
  struct Node *cur = head;
  for (;;)
  {
    if (cur == NULL)
    {
      printf("over\n");
      break;
    }

    printf("%d\n", print_val(cur));
    cur = cur->next;
  }
}

int print_val(struct Node *node)
{
  return (node->val);
}

void insert_at(struct Node *head, int val, int index)
{
  struct Node *node = (struct Node *)malloc(sizeof(struct Node));
  node->val = val;

  struct Node *cur = head;
  for (int i = 0; i < index; i++)
  {
    cur = cur->next;
  }

  node->next = cur->next;
  cur->next = node;

  head->len++;
  return;
}

void delete_at(struct Node *head, int index)
{
  if (index > head->len)
  {
    return;
  }

  struct Node *cur = head;
  for (int i = 0; i < index; i++)
  {
    cur = cur->next;
  }

  cur = cur->next;

  head->len--;
  return;
}

struct Node *reverse(struct Node *head)
{
  // answer 1
  // 时间复杂度 O(n)
  // 空间复杂度 O(n)
  struct Node *h1 = (struct Node *)malloc(sizeof(struct Node));
  struct Node *tmp = head;
  for (int i = 0; i <= head->len; i++)
  {
    struct Node *node = (struct Node *)malloc(sizeof(struct Node));

    node->val = tmp->val;
    tmp = tmp->next;

    node->next = h1->next;
    h1->next = node;
    h1->len++;
  }

  // TOOD answer 2
  // 时间复杂度 O(n)
  // 空间复杂度 O(1)
  //

  return h1;
}

bool is_cycle(struct Node *head)
{
  // answer 1 这种写法只能判断环发生在头上的。
  // 末尾有环就没法判断了
  //
  // struct Node * cur = head;
  // for (int i = 0; i<head->len; i++) {
  //  cur = cur->next;

  //  if (cur->next != NULL && head == cur->next) {
  //      return true;
  //  }
  //}

  // answer 2 快慢指针的方式
  // 快的一次走俩个
  // 慢的一次走一个
  struct Node *slow, *fast;
  slow = fast = head;

  for (;;)
  {
    slow = slow->next;
    fast = fast->next;

    if (fast->next != NULL)
    {
      fast = fast->next;
    }

    if (fast->next == NULL)
    {
      return false;
    }

    if (fast == slow)
    {
      return true;
    }
  }
}

struct Node *get_last(struct Node *head)
{
  struct Node *cur = head;
  for (int i = 0; i < head->len; i++)
  {
    cur = cur->next;
  }

  return cur;
}

struct Node *merge_2_sorted(struct Node *h1, struct Node *h2)
{
  struct Node *head = (struct Node *)malloc(sizeof(struct Node));

  struct Node *cur = head;
  for (; h1 != NULL && h2 != NULL;)
  {
    struct Node *node = (struct Node *)malloc(sizeof(struct Node));
    if (h1->val > h2->val)
    {
      node->val = h1->val;
      h1 = h1->next;
    }
    else
    {
      node->val = h2->val;
      h2 = h2->next;
    }

    cur->next = node;
    cur = cur->next;
  }

  if (h1 != NULL)
  {
    cur->next = h1;
  }

  if (h2 != NULL)
  {
    cur->next = h2;
  }

  return head;
}
