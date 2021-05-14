# Link: https://leetcode-cn.com/problems/list-the-products-ordered-in-a-period

SELECT `product_name`, SUM(`unit`) `unit`
FROM `products`        `p`
         JOIN `orders` `o`
              ON `p`.`product_id` = `o`.`product_id`
WHERE LEFT(`order_date`, 7) = '2020-02'
GROUP BY `product_name`
HAVING SUM(`unit`) >= 100;