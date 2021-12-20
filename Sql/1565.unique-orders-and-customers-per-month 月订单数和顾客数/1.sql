# Link: https://leetcode-cn.com/problems/unique-orders-and-customers-per-month

SELECT LEFT(`o`.`order_date`, 7)     `month`,
       COUNT(`order_id`)             `order_count`,
       COUNT(DISTINCT `customer_id`) `customer_count`
FROM `orders` `o`
WHERE `o`.`invoice` > 20
GROUP BY LEFT(`o`.`order_date`, 7);