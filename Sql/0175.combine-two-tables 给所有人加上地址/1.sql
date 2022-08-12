# Link: https://leetcode.cn/problems/combine-two-tables

SELECT `firstname`, `lastname`, `city`, `state`
FROM `person`
         LEFT JOIN `address` ON `person`.`personid` = `address`.`personid`