# Link: https://leetcode-cn.com/problems/sales-analysis-iii

SELECT `p`.`product_id`, `p`.`product_name`
FROM `product` `p`
         JOIN `sales` `s` ON `p`.`product_id` = `s`.`product_id`
GROUP BY `product_id`
HAVING (sum(`sale_date` BETWEEN '2019-01-01' AND '2019-03-31')) = count(*);

#

SELECT `product_id`,
       `product_name`
FROM `Product`
WHERE `product_id` NOT IN
      (SELECT `product_id`
       FROM `Sales`
       WHERE `sale_date` NOT BETWEEN '2019-01-01' AND '2019-03-31')