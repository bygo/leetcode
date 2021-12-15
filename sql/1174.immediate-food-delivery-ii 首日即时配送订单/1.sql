# Link: https://leetcode-cn.com/problems/immediate-food-delivery-ii

SELECT ROUND(
                   SUM(`order_date` = `customer_pref_delivery_date`) * 100 /
                   COUNT(*),
                   2
           ) `immediate_percentage`
FROM `delivery`
WHERE (`customer_id`, `order_date`) IN (
    SELECT `customer_id`, MIN(`order_date`)
    FROM `delivery`
    GROUP BY `customer_id`
)