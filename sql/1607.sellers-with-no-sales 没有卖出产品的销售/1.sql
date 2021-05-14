# Link: https://leetcode-cn.com/problems/sellers-with-no-sales


SELECT `seller_name`
FROM `seller`
WHERE `seller_id` NOT IN (SELECT DISTINCT `seller_id`
                          FROM `orders`
                          WHERE `sale_date` BETWEEN '2020-01-01' AND '2020-12-31')
ORDER BY `seller_name`;

SELECT `seller_name`
FROM `seller` `s`
WHERE NOT exists(SELECT DISTINCT `seller_id`
                 FROM `orders` `o`
                 WHERE `sale_date` BETWEEN '2020-01-01' AND '2020-12-31'
                   AND `o`.`seller_id` = `s`.`seller_id`)
ORDER BY `seller_name`;