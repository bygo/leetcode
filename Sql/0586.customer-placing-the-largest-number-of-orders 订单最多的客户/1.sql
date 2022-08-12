# Link: https://leetcode.cn/problems/customer-placing-the-largest-number-of-orders

SELECT `customer_number`
FROM `orders`
GROUP BY `customer_number`
ORDER BY COUNT(*) DESC
LIMIT 1