# Link: https://leetcode-cn.com/problems/product-sales-analysis-i

SELECT `product_name`, `year`, `price`
FROM `sales`
         JOIN `product` ON `sales`.`product_id` = `product`.`product_id`