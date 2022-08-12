# Link: https://leetcode.cn/problems/products-worth-over-invoices

SELECT `p`.`name`                     `name`,
       IFNULL(SUM(`i`.`rest`), 0)     `rest`,
       IFNULL(SUM(`i`.`paid`), 0)     `paid`,
       IFNULL(SUM(`i`.`canceled`), 0) `canceled`,
       IFNULL(SUM(`i`.`refunded`), 0) `refunded`
FROM `product`               `p`
         LEFT JOIN `invoice` `i`
                   ON `p`.`product_id` = `i`.`product_id`
GROUP BY `p`.`product_id`
ORDER BY `p`.`name`