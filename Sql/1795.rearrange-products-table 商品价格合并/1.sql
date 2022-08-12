# Link: https://leetcode.cn/problems/rearrange-products-table

SELECT `product_id`, 'store1' `store`, `store1` `price`
FROM `products`
WHERE `store1` IS NOT NULL
UNION ALL
SELECT `product_id`, 'store2' `store`, `store2` `price`
FROM `products`
WHERE `store2` IS NOT NULL
UNION ALL
SELECT `product_id`, 'store3' `store`, `store3` `price`
FROM `products`
WHERE `store3` IS NOT NULL
