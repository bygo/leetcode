# Link: https://leetcode-cn.com/problems/triangle-judgement

SELECT *, IF(`x` + `y` > `z` AND `x` + `z` > `y` AND `y` + `z` > `x`, "yes", "no") `triangle`
FROM `triangle`