# Link: https://leetcode.cn/problems/monthly-transactions-i


SELECT DATE_FORMAT(`trans_date`, '%Y-%m')         `month`,
       `country`,
       COUNT(*)                                   `trans_count`,
       COUNT(IF(`state` = 'approved', 1, NULL))   `approved_count`,
       SUM(`amount`)                              `trans_total_amount`,
       SUM(IF(`state` = 'approved', `amount`, 0)) `approved_total_amount`
FROM `transactions`
GROUP BY `month`, `country`