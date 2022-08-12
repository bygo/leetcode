# Link: https://leetcode.cn/problems/rising-temperature

SELECT `w2`.`id` `id`
FROM `weather`          `w2`
         JOIN `weather` `w1`
              ON DATEDIFF(`w2`.`recorddate`, `w1`.`recorddate`) = 1
                  AND `w1`.`temperature` < `w2`.`temperature`