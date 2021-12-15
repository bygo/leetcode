# Link: https://leetcode-cn.com/problems/maximum-transaction-each-day


SELECT `transaction_id`
FROM (
         SELECT `transaction_id`,
                RANK() OVER (PARTITION BY LEFT(`day`, 10) ORDER BY `amount` DESC) `r`
         FROM `transactions`
         ORDER BY `transaction_id`
     ) `t`
WHERE `r` = 1