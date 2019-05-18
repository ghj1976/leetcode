## [17. Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/)

###### 2019/3/2

###### 真彻底被打败了,身份证掉了,期间想去建行银行保险柜拿个东西,提供了护照+身份证补办证明+户口本+指纹,我却无法证明我是我自己

### 题目💗
```go
var mapList map[string][]string = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}
```

给定`23`,请根据上面的地图映射关系,统计出总共有多少种映射可能

> ### 解题思路
> [ "" ]
> XXXXXXXXXXXXXXX
> ["a", "b", "c"]
> 🔽🔽🔽🔽🔽🔽🔽
> ["a", "b", "c"]
> ===============
> =====next======
> ===============
> ["a", "b", "c"]
> XXXXXXXXXXXXXXX
> ["d", "e", "f"]
> 🔽🔽🔽🔽🔽🔽🔽
> ["ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"]