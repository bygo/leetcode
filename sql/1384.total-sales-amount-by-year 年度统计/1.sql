# Link: https://leetcode-cn.com/problems/total-sales-amount-by-year


SELECT `sales`.`product_id`,
       `product_name`,
       '2018' `report_year`,
       if(`period_start` < '2019-01-01',
          (datediff(if(`period_end` < '2019-01-01', `period_end`, date('2018-12-31')),
                    if(`period_start` >= '2018-01-01', `period_start`, date('2018-01-01'))) + 1) *
          `average_daily_sales`, 0) `total_amount`
FROM `sales`
         JOIN `product` ON `sales`.`product_id` = `product`.`product_id`
HAVING `total_amount` > 0

UNION

SELECT `sales`.`product_id`,
       `product_name`,
       '2019' `report_year`,
       if(`period_start` < '2020-01-01',
          (datediff(if(`period_end` < '2020-01-01', `period_end`, date('2019-12-31')),
                    if(`period_start` >= '2019-01-01', `period_start`, date('2019-01-01'))) + 1) *
          `average_daily_sales`, 0) `total_amount`
FROM `sales`
         JOIN `product` ON (`sales`.`product_id` = `product`.`product_id`)
HAVING `total_amount` > 0

UNION

SELECT `sales`.`product_id`,
       `product_name`,
       '2020' `report_year`,
       (datediff(if(`period_end` < '2021-01-01', `period_end`, date('2020-12-31')),
                 if(`period_start` >= '2020-01-01', `period_start`, date('2020-01-01'))) + 1) *
       `average_daily_sales` `total_amount`
FROM `sales`
         JOIN `product` ON (`sales`.`product_id` = `product`.`product_id`)
HAVING `total_amount` > 0

ORDER BY `product_id`, `report_year`


