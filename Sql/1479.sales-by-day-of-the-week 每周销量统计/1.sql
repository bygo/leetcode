# Link: https://leetcode.cn/problems/sales-by-day-of-the-week


SELECT DISTINCT `i`.`item_category`                                         `category`,
                SUM(IF(DAYOFWEEK(`o`.`order_date`) = 2, `o`.`quantity`, 0)) `monday`,
                SUM(IF(DAYOFWEEK(`o`.`order_date`) = 3, `o`.`quantity`, 0)) `tuesday`,
                SUM(IF(DAYOFWEEK(`o`.`order_date`) = 4, `o`.`quantity`, 0)) `wednesday`,
                SUM(IF(DAYOFWEEK(`o`.`order_date`) = 5, `o`.`quantity`, 0)) `thursday`,
                SUM(IF(DAYOFWEEK(`o`.`order_date`) = 6, `o`.`quantity`, 0)) `friday`,
                SUM(IF(DAYOFWEEK(`o`.`order_date`) = 7, `o`.`quantity`, 0)) `saturday`,
                SUM(IF(DAYOFWEEK(`o`.`order_date`) = 1, `o`.`quantity`, 0)) `sunday`
FROM `items`                `i`
         LEFT JOIN `orders` `o` ON `i`.`item_id` = `o`.`item_id`
GROUP BY `category`
ORDER BY `category`
