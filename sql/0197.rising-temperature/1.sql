# Title: Rising Temperature
# Link: https://leetcode-cn.com/problems/rising-temperature

SELECT `w1`.`Id` AS `Id`
FROM `Weather` `w1`
         JOIN `Weather` `w2`
WHERE DATEDIFF(`w1`.`recordDate`, `w2`.`recordDate`) = 1
  AND `w1`.`Temperature` > `w2`.`Temperature`