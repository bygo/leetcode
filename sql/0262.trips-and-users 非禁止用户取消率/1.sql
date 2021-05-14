# Link: https://leetcode-cn.com/problems/trips-and-users

SELECT `t`.`request_at` `Day`,
       ROUND(
               SUM(IF(`t`.`STATUS` = 'completed', 0, 1)) / COUNT(`t`.`STATUS`),
               2)       `Cancellation Rate`
FROM `Trips` `t`
         JOIN
     `Users` `u1`
         JOIN
     `Users` `u2`
     ON `t`.`client_id` = `u1`.`users_id`
         AND `u1`.`banned` = 'No'
         AND `t`.`driver_id` = `u2`.`users_id`
         AND `u2`.`banned` = 'No'
         AND `t`.`request_at` BETWEEN '2013-10-01' AND '2013-10-03'
GROUP BY `t`.`request_at`