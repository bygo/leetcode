# Link: https://leetcode.cn/problems/find-customer-referee


SELECT `name`
FROM `customer`
WHERE `referee_id` != 2
   OR `referee_id` IS NULL