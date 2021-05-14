# Link: https://leetcode-cn.com/problems/sales-by-day-of-the-week


SELECT DISTINCT `i`.`item_category`                                         `category`,
                sum(IF(dayofweek(`o`.`order_date`) = 2, `o`.`quantity`, 0)) `monday`,
                sum(IF(dayofweek(`o`.`order_date`) = 3, `o`.`quantity`, 0)) `tuesday`,
                sum(IF(dayofweek(`o`.`order_date`) = 4, `o`.`quantity`, 0)) `wednesday`,
                sum(IF(dayofweek(`o`.`order_date`) = 5, `o`.`quantity`, 0)) `thursday`,
                sum(IF(dayofweek(`o`.`order_date`) = 6, `o`.`quantity`, 0)) `friday`,
                sum(IF(dayofweek(`o`.`order_date`) = 7, `o`.`quantity`, 0)) `saturday`,
                sum(IF(dayofweek(`o`.`order_date`) = 1, `o`.`quantity`, 0)) `sunday`
FROM `items`                `i`
         LEFT JOIN `orders` `o` ON `i`.`item_id` = `o`.`item_id`
GROUP BY `category`
ORDER BY `category`
