# Link: https://leetcode-cn.com/problems/capital-gainloss

SELECT `stock_name`,
       sum(IF(`operation` = 'Buy', -`price`, `price`)) `capital_gain_loss`
FROM `stocks`
GROUP BY `stock_name`;