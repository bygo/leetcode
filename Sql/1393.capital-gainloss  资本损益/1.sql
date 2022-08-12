# Link: https://leetcode.cn/problems/capital-gainloss

SELECT `stock_name`,
       SUM(IF(`operation` = 'Buy', -`price`, `price`)) `capital_gain_loss`
FROM `stocks`
GROUP BY `stock_name`;