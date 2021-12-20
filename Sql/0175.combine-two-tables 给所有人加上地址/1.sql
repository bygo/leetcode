# Link: https://leetcode-cn.com/problems/combine-two-tables

SELECT `firstname`, `lastname`, `city`, `state`
FROM `person`
         LEFT JOIN `address` ON `person`.`personid` = `address`.`personid`