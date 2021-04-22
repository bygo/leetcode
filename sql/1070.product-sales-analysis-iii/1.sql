# Title: Product Sales Analysis III
# Link: https://leetcode-cn.com/problems/product-sales-analysis-iii


SELECT `product_id`, `year` AS `first_year`, `quantity`, `price`
FROM `Sales`
WHERE (`product_id`, `year`) IN (SELECT `product_id`, min(`year`)
                                 FROM `Sales`
                                 GROUP BY `product_id`);