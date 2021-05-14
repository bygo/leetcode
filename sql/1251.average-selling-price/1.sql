# Link: https://leetcode-cn.com/problems/average-selling-price

SELECT `product_id`,
       ROUND(SUM(`sales`) / SUM(`units`), 2) `average_price`
FROM (
         SELECT `p`.`product_id`          `product_id`,
                `p`.`price` * `u`.`units` `sales`,
                `u`.`units`               `units`
         FROM `prices`             `p`
                  JOIN `unitssold` `u` ON `p`.`product_id` = `u`.`product_id`
         WHERE `u`.`purchase_date` BETWEEN `p`.`start_date` AND `p`.`end_date`
     ) `t`
GROUP BY `product_id`