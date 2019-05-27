# [606. Construct String from Binary Tree](https://leetcode.com/problems/construct-string-from-binary-tree/)

## 2019/05/28

### 推荐程度

👍524 👎724

### 题目 💗[easy]

You need to construct a string consists of parenthesis and integers from a binary tree with the preorder traversing way.

你需要一个字符串结构,他与二叉树相关

---

The null node needs to be represented by empty parenthesis pair "()". And you need to omit all the empty parenthesis pairs that don't affect the one-to-one mapping relationship between the string and the original binary tree.

这个Null节点代表空节点, 用`()`表示,并且你需要发射

---

> Example 1:
>
> ```bash
> Input: Binary tree: [1,2,3,4]
>        1
>      /   \
>     2     3
>    /
>   4
> Output: "1(2(4))(3)"
> ```
>
> Explanation: Originallay it needs to be "1(2(4)())(3()())",
> but you need to omit all the unnecessary empty parenthesis pairs.
> And it will be "1(2(4))(3)".

> Example 2:
>
> ```bash
> Input: Binary tree: [1,2,3,null,4]
>        1
>      /   \
>     2     3
>      \
>       4
>
> Output: "1(2()(4))(3)"
> ```
>
> Explanation: Almost the same as the first example,
> except we can't omit the first parenthesis pair to break the one-to-one mapping relationship between the input and the output.
