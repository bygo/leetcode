# Link: https://leetcode-cn.com/problems/immediate-food-delivery-ii

SELECT round(
                   sum(`order_date` = `customer_pref_delivery_date`) * 100 /
                   count(*),
                   2
           ) `immediate_percentage`
FROM `Delivery`
WHERE (`customer_id`, `order_date`) IN (
    SELECT `customer_id`, min(`order_date`)
    FROM `delivery`
    GROUP BY `customer_id`
)