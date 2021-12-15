# Link: https://leetcode-cn.com/problems/find-customers-with-positive-revenue-this-year

SELECT `customer_id`
FROM `customers`
WHERE `year` = 2021
GROUP BY `customer_id`
HAVING SUM(`revenue`) > 0