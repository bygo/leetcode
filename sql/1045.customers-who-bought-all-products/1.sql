# Title: Customers Who Bought All Products
# Link: https://leetcode-cn.com/problems/customers-who-bought-all-products

SELECT `customer_id`
FROM `Customer`
GROUP BY `customer_id`
HAVING COUNT(DISTINCT `product_key`) = (SELECT COUNT(*) AS `cc` FROM `product`)