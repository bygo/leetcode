# Title: Sales Person
# Link: https://leetcode-cn.com/problems/sales-person

SELECT DISTINCT `name`
FROM `salesperson`
WHERE `sales_id` NOT IN (SELECT `sales_id`
                         FROM `orders`
                         WHERE `com_id` = (SELECT `com_id` FROM `company` WHERE `name` = 'RED'))