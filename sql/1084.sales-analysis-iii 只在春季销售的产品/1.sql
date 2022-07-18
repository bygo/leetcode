# Link: https://leetcode.cn/problems/sales-analysis-iii

SELECT `p`.`product_id`, `p`.`product_name`
FROM `product`        `p`
         JOIN `sales` `s` ON `p`.`product_id` = `s`.`product_id`
GROUP BY `product_id`
HAVING (SUM(`sale_date` BETWEEN '2019-01-01' AND '2019-03-31')) = COUNT(*);

#

SELECT `product_id`,
       `product_name`
FROM `product`
WHERE `product_id` NOT IN
      (SELECT `product_id`
       FROM `sales`
       WHERE `sale_date` NOT BETWEEN '2019-01-01' AND '2019-03-31')