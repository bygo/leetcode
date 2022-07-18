# Link: https://leetcode.cn/problems/total-sales-amount-by-year


SELECT `sales`.`product_id`,
       `product_name`,
       '2018' `report_year`,
       IF(`period_start` < '2019-01-01',
          (DATEDIFF(IF(`period_end` < '2019-01-01', `period_end`, DATE('2018-12-31')),
                    IF(`period_start` >= '2018-01-01', `period_start`, DATE('2018-01-01'))) + 1) *
          `average_daily_sales`, 0) `total_amount`
FROM `sales`
         JOIN `product` ON `sales`.`product_id` = `product`.`product_id`
HAVING `total_amount` > 0

UNION

SELECT `sales`.`product_id`,
       `product_name`,
       '2019' `report_year`,
       IF(`period_start` < '2020-01-01',
          (DATEDIFF(IF(`period_end` < '2020-01-01', `period_end`, DATE('2019-12-31')),
                    IF(`period_start` >= '2019-01-01', `period_start`, DATE('2019-01-01'))) + 1) *
          `average_daily_sales`, 0) `total_amount`
FROM `sales`
         JOIN `product` ON (`sales`.`product_id` = `product`.`product_id`)
HAVING `total_amount` > 0

UNION

SELECT `sales`.`product_id`,
       `product_name`,
       '2020' `report_year`,
       (DATEDIFF(IF(`period_end` < '2021-01-01', `period_end`, DATE('2020-12-31')),
                 IF(`period_start` >= '2020-01-01', `period_start`, DATE('2020-01-01'))) + 1) *
       `average_daily_sales` `total_amount`
FROM `sales`
         JOIN `product` ON (`sales`.`product_id` = `product`.`product_id`)
HAVING `total_amount` > 0

ORDER BY `product_id`, `report_year`


