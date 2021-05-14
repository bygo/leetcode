# Link: https://leetcode-cn.com/problems/market-analysis-i


SELECT `u`.`user_id`     `buyer_id`,
       `join_date`,
       count(`order_id`) `orders_in_2019`
FROM `users` `u`
         LEFT JOIN `orders` `o` ON `u`.`user_id` = `o`.`buyer_id` AND `order_date` BETWEEN '2019-01-01' AND '2019-12-31'
GROUP BY `user_id`