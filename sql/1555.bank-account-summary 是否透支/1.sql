# Link: https://leetcode-cn.com/problems/bank-account-summary


SELECT `u`.`user_id`,
       `u`.`user_name`,
       `credit` + IFNULL(`t`.`amount`, 0)                       `credit`,
       IF(`credit` + IFNULL(`t`.`amount`, 0) >= 0, 'No', 'Yes') `credit_limit_breached`
FROM `users`                            `u`
         LEFT JOIN (SELECT `user_id`, sum(`amount`) `amount`
                    FROM (
                             (
                                 SELECT `paid_by` `user_id`, -sum(`amount`) `amount`
                                 FROM `transactions`
                                 GROUP BY `paid_by`
                                 UNION
                                 SELECT `paid_to` `user_id`, sum(`amount`) `amount`
                                 FROM `transactions`
                                 GROUP BY `paid_to`)
                         ) `t1`
                    GROUP BY `user_id`) `t`
                   ON `u`.`user_id` = `t`.`user_id`
GROUP BY `user_id`