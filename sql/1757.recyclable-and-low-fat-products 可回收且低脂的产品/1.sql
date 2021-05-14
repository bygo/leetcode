# Link: https://leetcode-cn.com/problems/recyclable-and-low-fat-products

SELECT `product_id`
FROM `products`
WHERE `low_fats` = 'Y'
  AND `recyclable` = 'Y'