# Link: https://leetcode-cn.com/problems/product-sales-analysis-i

SELECT `product_name`, `year`, `price`
FROM `Sales`
         JOIN `Product` ON `Sales`.`product_id` = `Product`.`product_id`