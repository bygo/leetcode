# Link: https://leetcode-cn.com/problems/fix-product-name-format


SELECT lower(trim(`product_name`)) `product_name`,
       left(`sale_date`, 7)        `sale_date`,
       count(*)                    `total`
FROM `sales`
GROUP BY lower(trim(`product_name`)), left(`sale_date`, 7)
ORDER BY lower(trim(`product_name`)), left(`sale_date`, 7)
