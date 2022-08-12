# Link: https://leetcode.cn/problems/customers-who-bought-products-a-and-b-but-not-c

SELECT `c`.`customer_id`,
       `c`.`customer_name`
FROM `orders`                  `o`
         LEFT JOIN `customers` `c` ON `o`.`customer_id` = `c`.`customer_id`
GROUP BY `c`.`customer_id`
HAVING SUM(`product_name` = 'A') * SUM(`product_name` = 'B') > 0
   AND SUM(`product_name` = 'C') = 0
ORDER BY `c`.`customer_id`
