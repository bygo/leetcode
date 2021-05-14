# Link: https://leetcode-cn.com/problems/monthly-transactions-ii

SELECT LEFT(`trans_date`, 7)                        `month`,
       `country`,
       SUM(IF(`state` = 'approved', 1, 0))          `approved_count`,
       SUM(IF(`state` = 'approved', `amount`, 0))   `approved_amount`,
       SUM(IF(`state` = 'chargeback', 1, 0))        `chargeback_count`,
       SUM(IF(`state` = 'chargeback', `amount`, 0)) `chargeback_amount`
FROM (SELECT `id`, `country`, 'chargeback' `state`, `amount`, `c`.`trans_date`
      FROM `transactions`         `t1`
               JOIN `chargebacks` `c`
                    ON `t1`.`id` = `c`.`trans_id`
      UNION
      SELECT *
      FROM `transactions`) `t2`
GROUP BY `month`, `country`
HAVING SUM(IF(`state` = 'approved' OR `state` = 'chargeback', 1, 0)) > 0
ORDER BY `month`;