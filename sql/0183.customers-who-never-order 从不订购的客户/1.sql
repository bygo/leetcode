# Link: https://leetcode-cn.com/problems/customers-who-never-order

SELECT `customers`.`name` `customers`
FROM `customers`
         LEFT JOIN `orders` ON `customers`.`id` = `orders`.`customerid`
WHERE `orders`.`customerid` IS NULL;

#

SELECT `name` `customers`
FROM `customers`
WHERE `id` NOT IN (SELECT `customerid` FROM `orders`)