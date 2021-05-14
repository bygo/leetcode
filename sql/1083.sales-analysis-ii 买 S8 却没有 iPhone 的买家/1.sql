# Link: https://leetcode-cn.com/problems/sales-analysis-ii

SELECT `buyer_id`
FROM (SELECT `buyer_id`,
             IF(`p`.`product_name` = 'S8', 1, 0)     `s8`,
             IF(`p`.`product_name` = 'iPhone', 1, 0) `ip`
      FROM `sales` `s`
               JOIN `product` `p` ON `s`.`product_id` = `p`.`product_id`) `t`
GROUP BY `buyer_id`
HAVING 0 < sum(`ip`)
   AND sum(`s8`) = 0