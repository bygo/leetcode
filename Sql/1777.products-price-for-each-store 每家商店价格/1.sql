# Link: https://leetcode.cn/problems/products-price-for-each-store

SELECT `product_id`,
       SUM(IF(`store` = 'store1', `price`, NULL)) `store1`,
       SUM(IF(`store` = 'store2', `price`, NULL)) `store2`,
       SUM(IF(`store` = 'store3', `price`, NULL)) `store3`
FROM `products`
GROUP BY `product_id`;