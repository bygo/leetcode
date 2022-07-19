# Link: https://leetcode.cn/problems/customer-who-visited-but-did-not-make-any-transactions

SELECT `v`.`customer_id`,
       COUNT(DISTINCT `v`.`visit_id`) `count_no_trans`
FROM `visits`                     `v`
         LEFT JOIN `transactions` `t` ON `v`.`visit_id` = `t`.`visit_id`
WHERE `t`.`visit_id` IS NULL
GROUP BY `v`.`customer_id`

