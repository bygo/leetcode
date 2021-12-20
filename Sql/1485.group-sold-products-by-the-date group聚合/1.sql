# Link: https://leetcode-cn.com/problems/group-sold-products-by-the-date

SELECT `sell_date`,
       COUNT(DISTINCT `product`)        `num_sold`,
       GROUP_CONCAT(DISTINCT `product`) `products`
FROM `activities`
GROUP BY `sell_date`
ORDER BY `sell_date`