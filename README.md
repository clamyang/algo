# algo

|  类型  |                             题目                             | 思路                                                         | 备注                                                         |
| :----: | :----------------------------------------------------------: | :----------------------------------------------------------- | ------------------------------------------------------------ |
|  数组  | [二分查找](https://www.programmercarl.com/0704.%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE.html) | 没啥思路，凭感觉写？有两种形式左闭右开和左闭右闭             |                                                              |
|        | [移除元素](https://www.programmercarl.com/0027.%E7%A7%BB%E9%99%A4%E5%85%83%E7%B4%A0.html) | 1.暴力解决 <br />2.更优解为双指针                            | 需要巩固                                                     |
|        | [有序数组的平方](https://www.programmercarl.com/0977.%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E7%9A%84%E5%B9%B3%E6%96%B9.html) | 1.先全部进行平方操作，在进行排序<br />2.需要额外的存储空间，双指针 | 因为负数的存在                                               |
|        | [长度最小子数组](https://www.programmercarl.com/0209.%E9%95%BF%E5%BA%A6%E6%9C%80%E5%B0%8F%E7%9A%84%E5%AD%90%E6%95%B0%E7%BB%84.html#%E6%9A%B4%E5%8A%9B%E8%A7%A3%E6%B3%95) | 1.暴力解决<br />2.滑动窗口，降低时间复杂度                   | 在看过滑动窗口思路后，可以完整的实现出来。                   |
|        | [螺旋矩阵II](https://www.programmercarl.com/0059.%E8%9E%BA%E6%97%8B%E7%9F%A9%E9%98%B5II.html) | 1.没有任何思路，抄答案了解了一下                             | 虽然没涉及算法，但是感觉比较难！                             |
|        |                           三数之和                           | 1.双指针方式，但是需要先将数组进行排序                       | 在处理重复元素时，需要注意如何过滤。                         |
| 字符串 | [反转字符串](https://www.programmercarl.com/0344.%E5%8F%8D%E8%BD%AC%E5%AD%97%E7%AC%A6%E4%B8%B2.html) | 1. 双指针，两两交换                                          | 想到了这种方式值得鼓励                                       |
|        | [反转字符串](https://www.programmercarl.com/0541.%E5%8F%8D%E8%BD%AC%E5%AD%97%E7%AC%A6%E4%B8%B2II.html)II | 1.最开始也是没读懂题目，多读几遍就能解出来了                 | 大脑过载的时候确实需要休息！                                 |
|        | [替换空格](https://www.programmercarl.com/%E5%89%91%E6%8C%87Offer05.%E6%9B%BF%E6%8D%A2%E7%A9%BA%E6%A0%BC.html) | 1.暴力解决，空间复杂度较高<br />2.看了双指针思路自己实现     | 在于如何将一个空格扩大                                       |
|        | [反转字符串里的单词](https://www.programmercarl.com/0151.%E7%BF%BB%E8%BD%AC%E5%AD%97%E7%AC%A6%E4%B8%B2%E9%87%8C%E7%9A%84%E5%8D%95%E8%AF%8D.html) | 1.想到了一种简单的方式，直接 split 然后根据索引交换，但是失去了本题的意义<br />2.全部反转后，在进行部分的反转，一开始以空格作为分割条件，但是处理结束字符时出现了问题。最后的字符需要使用 \0 处理。 | 全部反转再部分反转                                           |
|        | [左旋转字符串](https://www.programmercarl.com/%E5%89%91%E6%8C%87Offer58-II.%E5%B7%A6%E6%97%8B%E8%BD%AC%E5%AD%97%E7%AC%A6%E4%B8%B2.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC) | 1.暴力解决<br />2.通过反转的方式实现，这种方法第一时间真的想不出来.. | 部分反转，部分反转，全部反转                                 |
|        | [strstr](https://www.programmercarl.com/0028.%E5%AE%9E%E7%8E%B0strStr.html) | 1.复习了 KMP 算法是怎么运行的（主要是理解最长公共前后缀的含义）以及 next 数组的构造和使用。<br />2.有两种形式，next数组中索引统一减一的方式和常规的方式。这两种方式对于 KMP 原理没有影响，主要影响在于 KMP 算法的实现。 | KMP 算法思想：当遇到不匹配的字符时，充分利用之前已经匹配过的字符，避免重新匹配。 |
|        | [重复子字符串](https://www.programmercarl.com/0459.%E9%87%8D%E5%A4%8D%E7%9A%84%E5%AD%90%E5%AD%97%E7%AC%A6%E4%B8%B2.html) | 结合了 KMP 使用，需要充分理解最长相等前后缀的含义。          |                                                              |
|  链表  |                           翻转链表                           | 1.新建一个链表，然后遍历原来的，采用头插法，得到的就是反转链表。时间、空间复杂度都是 O(n)<br />2.采用双指针的方式，pre, cur, next 一边遍历一遍翻转，最后时限的时候忘记了把 cur 指向 next 需要注意。 |                                                              |
|        |                       判断链表是否有环                       | 1.因为之前做过这个，所以直接采用了双指针的方式实现。慢指针一次走一个，快指针一次走两个，当快指针追上慢指针时，就说明链表中存在环。 | 如果说快指针一次走三个不是两个？还能证明链表中有环吗？我认为是可以的，他们都是 1 的倍数，只不过遍历的节点的数量会比2多。 |
|        |                         链表是否相交                         | 1.自己想的是，通过节点计数的方式实现。<br />2.双指针方式实现，不过需要一些额外的处理，因为两条链表长度不同，在开始遍历前需要将两条链表的起点对齐。 |                                                              |
|        |                          环形链表II                          | 寻找环的入口<br />自己这边没思路。                           | 既然没思路，就背下来！                                       |
|        |                   移除链表倒数第 n 个节点                    | 关键是如何找到被删除节点的前一个节点。<br />双指针方式实现，快指针先向前移动 n 个节点，然后快慢指针同时移动，当快指针指向末尾时，慢指针指向的就是被删除节点的前一个节点。 | 很奇妙。                                                     |
