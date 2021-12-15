# Title: Sales Analysis I
# Link: https://leetcode-cn.com/problems/sales-analysis-i

SELECT `seller_id`
FROM (
         SELECT `seller_id`, DENSE_RANK() OVER (ORDER BY `total` DESC) `n`
         FROM (
                  SELECT `seller_id`, SUM(`price`) `total` FROM `sales` GROUP BY `seller_id`) `t1`) `t2`
WHERE `n` = 1