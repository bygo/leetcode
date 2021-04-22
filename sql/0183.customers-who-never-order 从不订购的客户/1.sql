# Link: https://leetcode-cn.com/problems/customers-who-never-order

SELECT `Customers`.`Name` `Customers`
FROM `Customers`
         LEFT JOIN `Orders` ON `Customers`.`Id` = `Orders`.`CustomerId`
WHERE `Orders`.`CustomerId` IS NULL;

#

SELECT `name` `Customers`
FROM `Customers`
WHERE `Id` NOT IN (SELECT `CustomerId` FROM `Orders`)