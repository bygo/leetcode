# Link: https://leetcode.cn/problems/immediate-food-delivery-i

SELECT ROUND(
                       SUM(IF(`order_date` = `customer_pref_delivery_date`, 1, 0)) /
                       COUNT(*) * 100, 2) `immediate_percentage`
FROM `delivery`