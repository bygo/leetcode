# Link: https://leetcode-cn.com/problems/apples-oranges

SELECT `sale_date`,
       SUM(IF(`fruit` = `apples`, `sold_num`, -`sold_num`)) `diff`
FROM `sales`
GROUP BY `sale_date`
ORDER BY `sale_date`