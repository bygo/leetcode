# Link: https://leetcode.cn/problems/product-sales-analysis-iii


SELECT `product_id`, `year` `first_year`, `quantity`, `price`
FROM `sales`
WHERE (`product_id`, `year`) IN (SELECT `product_id`, MIN(`year`)
                                 FROM `sales`
                                 GROUP BY `product_id`);