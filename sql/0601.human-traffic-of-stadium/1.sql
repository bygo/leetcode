# Title: Human Traffic of Stadium
# Link: https://leetcode-cn.com/problems/human-traffic-of-stadium

WITH `countT` AS (SELECT `id`,
                         COUNT(*) OVER (PARTITION BY `rn` ORDER BY `rn` ) AS `counter`
                  FROM (SELECT `id`,
                               `id` - ROW_NUMBER() OVER (ORDER BY `id`) AS `rn`
                        FROM `stadium`
                        WHERE `people` >= 100) `rowT`)
SELECT `s`.*
FROM `stadium` `s`
         JOIN `countT` ON
    `s`.`id` = `countT`.`id`
WHERE `countT`.`counter` > 2
ORDER BY `s`.`visit_date`;