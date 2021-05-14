# Link: https://leetcode-cn.com/problems/product-price-at-a-given-date

SELECT DISTINCT `product_id`, `price`
FROM (
         (SELECT `product_id`, @`price` := 10 `price`
          FROM `products` `p`
          GROUP BY `product_id`)
         UNION
         SELECT `product_id`, `price`
         FROM (SELECT `product_id`,
                      rank() OVER (PARTITION BY `product_id` ORDER BY `change_date` DESC) `r`,
                      `new_price`                                                         `price`
               FROM `products`
               WHERE `change_date` <= '2019-08-16'
               GROUP BY `product_id`) `t2`
         WHERE `r` = 1
     ) `t`;

SELECT DISTINCT `p`.`product_id`, IFNULL(`t`.`new_price`, 10) `price`
FROM `products` `p`
         LEFT JOIN (SELECT `product_id`,
                           `new_price`,
                           RANK() OVER (PARTITION BY `product_id` ORDER BY `change_date` DESC) `r`
                    FROM `products`
                    WHERE `change_date` <= '2019-08-16') `t`
                   ON `P`.`product_id` = `t`.`product_id`
                       AND `r` = 1;
