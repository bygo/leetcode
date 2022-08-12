# Link: https://leetcode.cn/problems/the-most-recent-orders-for-each-product

SELECT `product_name`, `p`.`product_id`, `order_id`, `order_date`
FROM `products`          `p`
         JOIN
     (SELECT `o`.`product_id`,
             `o`.`order_id`,
             `o`.`order_date`,
             RANK() OVER (PARTITION BY `o`.`product_id` ORDER BY `o`.`order_date` DESC) `r`
      FROM `orders` `o`) `t` ON `p`.`product_id` = `t`.`product_id`
WHERE `r` = 1
ORDER BY `p`.`product_name`, `p`.`product_id`, `t`.`order_id`;
