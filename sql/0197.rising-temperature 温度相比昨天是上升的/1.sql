# Link: https://leetcode-cn.com/problems/rising-temperature

SELECT `w2`.`Id` `Id`
FROM `Weather` `w2`
         JOIN `Weather` `w1`
              ON DATEDIFF(`w2`.`recordDate`, `w1`.`recordDate`) = 1
                  AND `w1`.`Temperature` < `w2`.`Temperature`