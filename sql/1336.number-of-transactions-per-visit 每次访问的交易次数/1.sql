# Link: https://leetcode-cn.com/problems/number-of-transactions-per-visit

WITH `cte` AS (SELECT COUNT(`amount`) `transactions_count`
               FROM `visits`                     `v`
                        LEFT JOIN `transactions` `t2`
                                  ON `v`.`user_id` = `t2`.`user_id`
                                      AND `v`.`visit_date` = `t2`.`transaction_date`
               GROUP BY `v`.`user_id`, `v`.`visit_date`)

SELECT `t1`.`transactions_count`, COUNT(`cte`.`transactions_count`) `visits_count`
FROM (SELECT 0 `transactions_count`
      UNION
      SELECT ROW_NUMBER() OVER () `transactions_count`
      FROM `transactions`) `t1`
         LEFT JOIN `cte`
                   ON `t1`.`transactions_count` = `cte`.`transactions_count`
GROUP BY 1
HAVING `t1`.`transactions_count` <= (SELECT MAX(`transactions_count`)
                                     FROM `cte`);
