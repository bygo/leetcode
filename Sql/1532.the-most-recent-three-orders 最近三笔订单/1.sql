# Link: https://leetcode.cn/problems/the-most-recent-three-orders


SELECT `name` `customer_name`,
       `customer_id`,
       `order_id`,
       `order_date`
FROM (
         SELECT `c`.`name`,
                `c`.`customer_id`,
                `o`.`order_id`,
                `o`.`order_date`,
                ROW_NUMBER() OVER (PARTITION BY `c`.`customer_id` ORDER BY `o`.`order_date` DESC) `r`
         FROM `customers`       `c`
                  JOIN `orders` `o` ON `c`.`customer_id` = `o`.`customer_id`
     ) AS `a`
WHERE `r` <= 3
ORDER BY `customer_name`, `customer_id`, `order_date` DESC