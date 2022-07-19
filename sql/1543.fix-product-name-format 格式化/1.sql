# Link: https://leetcode.cn/problems/fix-product-name-format


SELECT LOWER(TRIM(`product_name`)) `product_name`,
       LEFT(`sale_date`, 7)        `sale_date`,
       COUNT(*)                    `total`
FROM `sales`
GROUP BY LOWER(TRIM(`product_name`)), LEFT(`sale_date`, 7)
ORDER BY LOWER(TRIM(`product_name`)), LEFT(`sale_date`, 7)
