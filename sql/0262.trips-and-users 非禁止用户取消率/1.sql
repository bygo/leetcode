# Link: https://leetcode-cn.com/problems/trips-and-users

SELECT `t`.`request_at` `day`,
       ROUND(
               SUM(IF(`t`.`status` = 'completed', 0, 1)) / COUNT(`t`.`status`),
               2)       `cancellation rate`
FROM `trips` `t`
         JOIN
     `users` `u1`
         JOIN
     `users` `u2`
     ON `t`.`client_id` = `u1`.`users_id`
         AND `u1`.`banned` = 'No'
         AND `t`.`driver_id` = `u2`.`users_id`
         AND `u2`.`banned` = 'No'
         AND `t`.`request_at` BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY `t`.`request_at`