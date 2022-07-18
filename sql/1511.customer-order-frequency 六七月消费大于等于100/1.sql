# Link: https://leetcode.cn/problems/customer-order-frequency

SELECT `c`.`customer_id`, `c`.`name`
FROM `customers`        `c`
         JOIN `orders`  `o` ON `o`.`customer_id` = `c`.`customer_id`
         JOIN `product` `p` ON `p`.`product_id` = `o`.`product_id`
GROUP BY `c`.`customer_id`, `c`.`name`
HAVING SUM(IF(LEFT(`o`.`order_date`, 7) = '2020-06', `p`.`price` * `o`.`quantity`, 0)) >= 100
   AND SUM(IF(LEFT(`o`.`order_date`, 7) = '2020-07', `p`.`price` * `o`.`quantity`, 0)) >= 100;