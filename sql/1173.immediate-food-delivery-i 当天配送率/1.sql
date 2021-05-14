# Link: https://leetcode-cn.com/problems/immediate-food-delivery-i

SELECT round(
                       sum(if(`order_date` = `customer_pref_delivery_date`, 1, 0)) /
                       count(*) * 100, 2) `immediate_percentage`
FROM `delivery`