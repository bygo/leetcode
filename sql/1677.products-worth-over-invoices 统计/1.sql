# Link: https://leetcode-cn.com/problems/products-worth-over-invoices

SELECT `p`.`name`                     `name`,
       ifnull(sum(`i`.`rest`), 0)     `rest`,
       ifnull(sum(`i`.`paid`), 0)     `paid`,
       ifnull(sum(`i`.`canceled`), 0) `canceled`,
       ifnull(sum(`i`.`refunded`), 0) `refunded`
FROM `product`               `p`
         LEFT JOIN `invoice` `i`
                   ON `p`.`product_id` = `i`.`product_id`
GROUP BY `p`.`product_id`
ORDER BY `p`.`name`