# Link: https://leetcode.cn/problems/the-most-frequently-ordered-products-for-each-customer

SELECT `o`.`customer_id`, `o`.`product_id`, `p`.`product_name`
FROM (
         SELECT `customer_id`, `product_id`, RANK() OVER (PARTITION BY `customer_id` ORDER BY `c` DESC) `r`
         FROM (
                  SELECT `customer_id`, `product_id`, COUNT(*) `c`
                  FROM `orders`
                  GROUP BY `customer_id`, `product_id`
              ) `a`
     )                   `o`
         JOIN `products` `p`
              ON `o`.`product_id` = `p`.`product_id`
WHERE `o`.`r` = 1
ORDER BY `customer_id`, `product_id`