# Link: https://leetcode-cn.com/problems/products-price-for-each-store

SELECT `product_id`,
       sum(IF(`store` = 'store1', `price`, NULL)) `store1`,
       sum(IF(`store` = 'store2', `price`, NULL)) `store2`,
       sum(IF(`store` = 'store3', `price`, NULL)) `store3`
FROM `products`
GROUP BY `product_id`;