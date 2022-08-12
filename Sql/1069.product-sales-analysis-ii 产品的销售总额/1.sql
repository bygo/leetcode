# Link: https://leetcode.cn/problems/product-sales-analysis-ii

SELECT `product_id`, SUM(`quantity`) `total_quantity`
FROM `sales`
GROUP BY `product_id`