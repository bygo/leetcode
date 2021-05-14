# Link: https://leetcode-cn.com/problems/product-sales-analysis-ii

SELECT `product_id`, SUM(`quantity`) AS `total_quantity`
FROM `sales`
GROUP BY `product_id`